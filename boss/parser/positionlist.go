package parser

import (
	"fmt"
	"learn/job/engine"
	"regexp"
)

const baseUrl = "https://www.zhipin.com/c100010000-%s/?ka=sel-city-100010000"

func PrintPositionList(contents []byte) engine.ParseResult {
	// [^>]*> 这个意思是如果你不是>我就会匹配进来
	re := regexp.MustCompile(`<a href="/c[0-9]{9}-(p[0-9]{6})/"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	// FindAllSubmatch 返回的是[][][]byte,所以我们需要循环出来，在取对应的索引
	for _, m := range matches {
		Rurl := fmt.Sprintf(baseUrl, string(m[1]))
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        Rurl,
			ParserFunc: parserPosition,
		})
	}
	return result
}
