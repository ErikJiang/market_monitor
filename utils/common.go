package utils

import (
	"encoding/json"
	"io/ioutil"
)


type JsonStruct struct {}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (j *JsonStruct) Load(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}
