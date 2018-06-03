package parser

import (
	"learn/job/engine"
	"learn/job/model"
	"regexp"
	"strings"
)

var timeRe = regexp.MustCompile(`<div class="job-author"><span class="time">发布于(.*)[^<]*</span></div>`)
var saleryRe = regexp.MustCompile(`<div class="name"><h1>java开发工程师</h1> <span class="badge">(.*)[^<]*</span></div>`)
var cityRe = regexp.MustCompile(`<p>城市：(.*)[^<]*<em class="vline"></em>经验：1年以内<em class="vline"></em>学历：本科</p>`)
var experienceRe = regexp.MustCompile(`<p>城市：兰州<em class="vline"></em>经验：(.*)[^<]*<em class="vline"></em>学历：本科</p>`)
var educationRe = regexp.MustCompile(`<p>城市：兰州<em class="vline"></em>经验：1年以内<em class="vline"></em>学历：(.*)[^<]*</p>`)
var ismarketRe = regexp.MustCompile(`<p>(.*)[^<]*<em class="vline"></em>1000-9999人<em class="vline"></em><a ka="job-detail-brandindustry" href="/i100021/">计算机软件</a></p>`)
var personalRe = regexp.MustCompile(`<p>已上市<em class="vline"></em>(.*)人<em class="vline"></em><a ka="job-detail-brandindustry" href="/i100021/">计算机软件</a></p>`)
var positionRe = regexp.MustCompile(`岗位职责：(.*)任职要求`)
var requireRe = regexp.MustCompile(`任职要求：(.*)[^<]*`)

func ParsePositionInfo(contents []byte) engine.ParseResult {
	position := model.Position{}
	position.Time = extractString(contents, timeRe)
	position.Education = extractString(contents, educationRe)
	position.City = extractString(contents, cityRe)
	position.Experience = extractString(contents, experienceRe)
	position.Personal = extractString(contents, personalRe)
	position.Position = strings.Replace(extractString(contents, positionRe), "<br/>", "", -1)
	position.Salery = extractString(contents, saleryRe)
	position.Ismarket = extractString(contents, ismarketRe)
	position.Require = strings.Replace(extractString(contents, requireRe), "<br/>", "", -1)
	result := engine.ParseResult{
		Items: []interface{}{position},
	}
	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
