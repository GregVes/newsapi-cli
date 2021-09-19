package utilities

import (
	"reflect"
)

type GenericStruct interface{}

func GetStructKeysNames(obj GenericStruct) []string {
	var res = []string{}

	e := reflect.ValueOf(obj).Elem()

	for i := 0; i < e.NumField(); i++ {
		res = append(res, e.Type().Field(i).Name)
	}

	return res
}
