package model

import "encoding/json"

type Position struct {
	Time       string
	Salery     string
	City       string
	Experience string
	Education  string
	Ismarket   string
	Personal   string
	Position   string
	Require    string
}

func FromJsonObj(o interface{}) (Position, error) {
	var profile Position
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
