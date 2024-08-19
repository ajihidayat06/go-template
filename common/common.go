package common

import (
	"fmt"
	"reflect"
)

func IsStructPopulated(s interface{}) bool {
	// Reflect over the struct to iterate over its fields
	val := reflect.ValueOf(s)

	// Check if the value is a pointer and get the underlying element
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Ensure the provided value is a struct
	if val.Kind() != reflect.Struct {
		return false
	}

	// Iterate over the fields of the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// Check if the field has a non-zero value
		if !field.IsZero() {
			return true
		}
	}
	return false
}

func ArrInterfaceToArrStr(arrInterface ...interface{}) []string {
	var result []string
	for _, v := range arrInterface {
		result = append(result, fmt.Sprintf(`%v`, v))
	}

	return result
}
