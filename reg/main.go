package main

import (
	"io/ioutil"
	"regexp"

	"github.com/swanwish/go-common/logs"
)

var reg *regexp.Regexp
var pattern string
var source string

func regexpMatch() {
	//  xy 匹配x y
	// x|y  匹配x或者y 优先x
	// source = "asdfdsxxxyyfergsfasfxyfa"
	// pattern = `x|y|a`

	//x* 匹配零个或者多个x,优先匹配多个
	//x+ 匹配一个或者多个x，优先匹配多个
	//x? 匹配零个或者一个x，优先匹配一个

	//source = "xxxxewexxxasdfdsxxxyyfergsfasfxyfa"
	//pattern = `x*`

	// x{n,m} 匹配n个到m个x，优先匹配m个
	// x{n,}  匹配n个到多个x，优先匹配更多
	// x{n} 或者x{n}?  只匹配n个x
	//source = "xxxxxxxewexxxasdfdsxxxyyfergsfasfxyfa"
	//pattern = `x{4,}`

	// x{n,m}? 匹配n个到m个x，优先匹配n个
	// x{n,}?  匹配n个到多个x，优先匹配n个
	// x*?   匹配零个或者多个x，优先匹配0个
	// x+?   匹配一个或者多个x，优先匹配1个
	// x??   匹配零个或者一个x，优先匹配0个
	//source = "xxxxxxxewexxxasdfdsxxxyyfergsfasfxyfa"
	//pattern = `x??`

	//[\d] 或者[^\D] 匹配数字
	//[^\d]或者 [\D] 匹配非数字
	//source = "xx435ff5237yy6346fergsfasfxyfa"
	//pattern = `[\d]{3,}` //匹配3个或者更多个数字

	//source = "xx435ffGUTEYgjk52RYPHFY37yy6346ferg6987sfasfxyfa"
	//pattern = `[a-z]{3,}` //三个或者多个小写字母

	//source = "xx435ffGUTEYgjk52RYPHFY37yy6346ferg6987sfasfxyfa"
	//pattern = `[[:alpha:]]{5,}` //5个或者多个字母，相当于A-Za-z

	//source = "xx435,./$%(*(_&jgshgs发个^$%ffG返回福hjh放假啊啥UTEYgjk52R创YPHFY37yy6346ferg6987sfasfxyfa"
	//pattern = `[\p{Han}]+` //匹配连续的汉字

	//source = "13244820821HG74892109977HJA15200806084S11233240697hdgsfhah假发发货"
	//pattern = `1[3|5|7|8|][\d]{9}` //匹配电话号码

	//source = "创1:2-3"
	//source = "创 1:2"
	//source = "创 1:2~2:3"
	////logs.Debugf("The first character is %s", source[:1])
	//pattern = `^[\p{Han}]+`
	//pattern = `[\d]+(:[\d]+(-[\d]+)*)*(~([\d]+(:[\d]+(-[\d]+)*)*))*`

	//source = "132@12.comGKGk15@163.cn200806084S11233240697hdgsfhah假发发货"
	//pattern = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱

	//匹配用户名或者密码 `^[a-zA-Z0-9_-]{4,16}$`  字母或者数字开头，区分大小写，最短4位最长16位
	//匹配IP地址1 `^$(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`

	//匹配IP地址2
	//pattern = `((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`

	//匹配日期 年-月-日 `(\d{4}|\d{2})-((1[0-2])|(0?[1-9]))-(([12][0-9])|(3[01])|(0?[1-9]))`
	//匹配日期 月-日-年  `((1[0-2])|(0?[1-9]))/(([12][0-9])|(3[01])|(0?[1-9]))/(\d{4}|\d{2})`
	//匹配时间 小时：分钟 24小时制 ` ((1|0?)[0-9]|2[0-3]):([0-5][0-9]) `
	//匹配邮编  `[1-9][\d]5`
	//匹配URL `[a-zA-z]+://[^\s]*`

	//reg = regexp.MustCompile(pattern)
	//strs := reg.FindAllString(source, -1)
	//for index, str := range strs {
	//	logs.Debugf("The str %d is %s", index, str)
	//}

	//result := reg.FindAllStringSubmatch(source, -1)
	//logs.Debugf("The result is: %v", result)

	//fmt.Printf("%s\n", reg.FindAllString(source, -1))

	pattern = `<img[^s]+src=("(.+)"|'(.+)'|(.+))[^/<]+(/>|</img>)`
	pattern = `<img[^s]+src="(.+)"/>`
	fileContent, err := ioutil.ReadFile("/Users/Stephen/Downloads/book_6/EPUB/120.xhtml")
	if err != nil {
		logs.Errorf("Failed to read file, the error is %v", err)
		return
	}
	reg = regexp.MustCompile(pattern)

	strContent := string(fileContent)
	logs.Debugf("The file content is %s", strContent)
	result := reg.FindAllSubmatchIndex(fileContent, -1) //.FindAllStringIndex(strContent, -1)
	logs.Debugf("The result is %v", result)
	for _, item := range result {
		imgLine := strContent[item[0]:item[1]]
		logs.Debugf("img line is: %s", imgLine)
		matchItem := strContent[item[2]:item[3]]
		logs.Debugf("match item: %s", matchItem)
		//srcIndex := strings.Index(imgLine, `src="`)
		//if srcIndex != -1 {
		//
		//}
	}

}
func main() {
	regexpMatch()
}
