package epub_file

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-common/utils"
)

const (
	EpubFileNameRoot           = "package.opf"
	EpubFileNameToc            = "toc.xhtml"
	EpubFileNameBibleReference = "biblereferences.xml"
	EpubFileNameMimetype       = "mimetype"
	EpubFileNameContainerXml   = "container.xml"
	EpubFileNameTocNcx         = "toc.ncx"
)

type EPubFile struct {
	Name                  string
	SourceDir             string
	DestDir               string
	Writer                *zip.Writer
	Key                   string
	IgnoreEncryptionFiles []string
}

func (ef *EPubFile) Create(name, key, sourceDir, destDir string) error {
	if !strings.HasSuffix(name, ".epub") {
		name += ".epub"
		logs.Debugf("Change file name to %s", name)
	}
	ef.Name = name
	destFile := filepath.Join(destDir, name)
	file, err := os.Create(destFile)
	if err != nil {
		logs.Errorf("Failed to create file %s, the error is %v", ef.Name, err)
		return err
	}
	ef.Writer = zip.NewWriter(file)
	ef.Key = key
	ef.SourceDir = sourceDir
	ef.DestDir = destDir
	ef.IgnoreEncryptionFiles = []string{
		EpubFileNameBibleReference,
		EpubFileNameToc,
		EpubFileNameRoot,
		EpubFileNameContainerXml,
		EpubFileNameMimetype,
		EpubFileNameTocNcx}
	return nil
}

func (ef *EPubFile) Generate() error {
	err := ef.AddAll(ef.SourceDir, false)
	if err != nil {
		logs.Errorf("Failed to create epub file, the error is %v", err)
		return err
	}
	err = ef.Close()
	if err != nil {
		logs.Errorf("Failed to close epub file, the error is %v", err)
		return err
	}
	return nil
}

func (ef *EPubFile) Add(name string, content []byte) error {
	iow, err := ef.Writer.Create(name)
	if err != nil {
		return err
	}

	_, err = iow.Write(content)

	return err
}

func (ef *EPubFile) AddFile(name, headerName string) error {
	if headerName == "" {
		headerName = name
	}
	stat, err := os.Stat(name)
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(stat)
	if err != nil {
		return err
	}
	zippedFile, err := ef.Writer.CreateHeader(header)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(filepath.Join(name))
	if err != nil {
		return err
	}
	encryptedContent, err := ef.encryptContent(name, content)
	if err != nil {
		return err
	}
	zippedFile.Write(encryptedContent)
	return nil
}

func (ef *EPubFile) AddAll(dir string, includeCurrentFolder bool) error {
	dir = path.Clean(dir)
	mimetypeFilePath := path.Join(dir, EpubFileNameMimetype)
	if !utils.FileExists(mimetypeFilePath) {
		logs.Errorf("Failed to get mimetype file path")
	}
	if err := ef.AddFile(mimetypeFilePath, EpubFileNameMimetype); err != nil {
		return err
	}
	return addAll(dir, dir, includeCurrentFolder, func(info os.FileInfo, file io.Reader, entryName string) (err error) {
		if entryName == EpubFileNameMimetype {
			return nil
		}

		if strings.HasPrefix(entryName, ".") {
			logs.Debugf("Skip hidden file %s", entryName)
			return nil
		}

		if filepath.Ext(entryName) == ".epub" {
			logs.Debugf("Skip epub file %s", entryName)
			return nil
		}

		// Create a header based off of the fileinfo
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// If it's a file, set the compression method to deflate (leave directories uncompressed)
		if !info.IsDir() {
			header.Method = zip.Deflate
		}

		// Set the header's name to what we want--it may not include the top folder
		header.Name = entryName

		// Add a trailing slash if the entry is a directory
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		}

		// Get a writer in the archive based on our header
		writer, err := ef.Writer.CreateHeader(header)
		if err != nil {
			return err
		}

		// If we have a file to write (i.e., not a directory) then pipe the file into the archive writer
		if file != nil {
			content, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			encryptedContent, err := ef.encryptContent(entryName, content)
			if err != nil {
				return err
			}
			writer.Write(encryptedContent)
			//if _, err := io.Copy(writer, file); err != nil {
			//	return err
			//}
		}

		return nil
	})
}

func (ef *EPubFile) Close() error {
	err := ef.Writer.Close()
	return err
}

func (ef *EPubFile) encryptContent(name string, content []byte) ([]byte, error) {
	logs.Debugf("The name is %s", name)
	if ef.Key == "" {
		return content, nil
	}
	_, fileName := filepath.Split(name)

	for _, ignoreEncryptioinFile := range ef.IgnoreEncryptionFiles {
		if fileName == ignoreEncryptioinFile {
			logs.Debugf("Return plain text file %s", name)
			return content, nil
		}
	}

	return content, nil
}

type ArchiveWriteFunc func(info os.FileInfo, file io.Reader, entryName string) (err error)

func addAll(dir string, rootDir string, includeCurrentFolder bool, writerFunc ArchiveWriteFunc) error {
	// Get a list of all entries in the directory, as []os.FileInfo
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Loop through all entries
	for _, info := range fileInfos {

		full := path.Join(dir, info.Name())

		// If the entry is a file, get an io.Reader for it
		var file *os.File
		var reader io.Reader
		if !info.IsDir() {
			file, err = os.Open(full)
			if err != nil {
				return err
			}
			reader = file
		}

		// Write the entry into the archive
		subDir := getSubDir(dir, rootDir, includeCurrentFolder)
		entryName := path.Join(subDir, info.Name())
		if err := writerFunc(info, reader, entryName); err != nil {
			if file != nil {
				file.Close()
			}
			return err
		}

		if file != nil {
			if err := file.Close(); err != nil {
				return err
			}

		}

		// If the entry is a directory, recurse into it
		if info.IsDir() {
			addAll(full, rootDir, includeCurrentFolder, writerFunc)
		}
	}

	return nil
}

func getSubDir(dir string, rootDir string, includeCurrentFolder bool) (subDir string) {
	subDir = strings.Replace(dir, rootDir, "", 1)
	// Remove leading slashes, since this is intentionally a subdirectory.
	if len(subDir) > 0 && subDir[0] == os.PathSeparator {
		subDir = subDir[1:]
	}

	if includeCurrentFolder {
		parts := strings.Split(rootDir, string(os.PathSeparator))
		subDir = path.Join(parts[len(parts)-1], subDir)
	}

	return
}
