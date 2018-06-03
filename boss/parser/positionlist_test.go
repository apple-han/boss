package parser

import (
	"io/ioutil"
	"testing"
)

func TestPrintPositionList(t *testing.T) {
	contents, err := ioutil.ReadFile("positionlist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := PrintPositionList(contents)
	const resultSize = 823
	expectedUrls := []string{
		"https://www.zhipin.com/c100010000-p100101/?ka=sel-city-100010000",
		"https://www.zhipin.com/c100010000-p100103/?ka=sel-city-100010000",
		"https://www.zhipin.com/c100010000-p100901/?ka=sel-city-100010000",
	}
	expectedPositiones := []string{
		"Java", "PHP", "web前端",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d",
			resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, position := range expectedPositiones {
		if result.Items[i].(string) != position {
			t.Errorf("expected position #%d: %s; but "+
				"was %s", i, position, result.Items[i].(string))
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d",
			resultSize, len(result.Items))
	}
}
