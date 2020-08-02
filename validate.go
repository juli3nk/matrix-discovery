package main

import (
	"reflect"

	"github.com/thoas/go-funk"
	"gopkg.in/go-playground/validator.v10"
)

func validatePort(fl validator.FieldLevel) bool {
	portNum := fl.Field().Int()

	if portNum > 65535 || portNum < 1 {
		return false
	}

	return true
}

func validateCorsAllowMethods(fl validator.FieldLevel) bool {
	methods := []string{"GET", "HEAD", "OPTIONS"}

	field := fl.Field()
	kind := field.Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		return false
	}

	values := field.Interface().([]string)

	for _, v := range values {
		if !funk.Contains(methods, v) {
			return false
		}
	}

	return true
}
