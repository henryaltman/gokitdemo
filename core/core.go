package core

import (
	"errors"
	"fmt"
	"gokitdemo/services"
	"reflect"
)

func CallReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		return nil
	} else {
		return v.Call(inputs)
	}
}

func CallReflectMethod(svc services.Service, method string, args ...interface{}) (res interface{}, err error) {
	if result := CallReflect(svc, method, args...); result != nil {
		return result[0], nil
	}
	err = errors.New(fmt.Sprintf("not found method %s", method))
	return nil, err
}
