package parser

import (
	"io/ioutil"
	"learn/job/model"
	"testing"
)

func TestParsePositionInfo(t *testing.T) {
	contents, err := ioutil.ReadFile("positioninfo_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParsePositionInfo(contents)
	if len(result.Items) != 1 {
		t.Error("Item should  contain 1 element;"+
			"but was %v", result.Items)
	}

	position := result.Items[0].(model.Position)
	expected := model.Position{
		Time:       "2018-05-30 17:32",
		Salery:     "5K-10K",
		City:       "兰州",
		Experience: "1年以内",
		Education:  "本科",
		Ismarket:   "已上市",
		Personal:   "1000-9999",
		Position:   "1.Java类项目开发；2.负责按时按质完成项目经理交给的项目工作任务；3.负责主动了解工作范围边界和需求规格定义，对不明确的地方，第一时间主动提出来细化、清晰化；4.负责主动积累、学习完成工作所需的业务知识和技术知识；5.完成上级领导交办的其他工作。",
		Require:    "1.1到3年Java类项目开发经验，了解Java常用的设计模式和开发框架；2.熟练掌握数据结构、常用算法；3.掌握至少1种主流联机数据库例如ORACLE，DB2等，熟练使用SQL语言；4.熟悉Linux或UNIX系统的操作；5.熟悉软件工程和软件开发流程，具备良好的软件开发相关文档的编写经验。",
	}
	if position != expected {
		t.Error("expected %v; but was %v", "", position)
	}
}
