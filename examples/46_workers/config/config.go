package config

import "search-and-navigation/common/nsq"

type Configuration struct {
 	SomethingParent struct {
		Something `json:"something"`
	} `json:"something_parent"`
	SomeIds    []string `json:"some_ids"`
}

type Something struct {
	Name string `json:"name"`
	Age int `json:"age"`
}