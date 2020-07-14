package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	callReflectMethod("Chu", 1, 1)
}

type Base struct {
	Rs  interface{}
	Err error
}
type Test struct{}

func (t *Test) Chu(a int, b int) (br Base) {
	//fmt.Println("call method PrintInfo i", i, ",s :", s)
	if b == 0 {
		br.Err = errors.New("dddd")
		return br
	}
	br.Rs = a / b
	return br
}

func (t *Test) PrintInfo(i int, s string) string {
	fmt.Println("call method PrintInfo i", i, ",s :", s)
	return s + strconv.Itoa(i)
}

func (t *Test) ShowMsg() string {
	fmt.Println("\nshow msg input 'call reflect'")
	return "ShowMsg"
}
func (t *Test) ShowMsgArgs(a, b int) int {
	fmt.Println("\nshow msg input 'call reflect'")
	return a + b
}

func callReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
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

func callReflectMethod(method string, args ...interface{}) {
	// fmt.Printf("\n callReflectMethod PrintInfo :%s", callReflect(&Test{}, "PrintInfo", 10, "TestMethod")[0].String())
	// fmt.Printf("\n callReflectMethod ShowMsg  %s", callReflect(&Test{}, "ShowMsg")[0].String())

	// //<invalid Value> case
	// callReflect(&Test{}, "ShowMs")
	// if result := callReflect(&Test{}, "ShowMs"); result != nil {
	// 	fmt.Printf("\n callReflectMethod ShowMs %s", result[0].String())
	// } else {
	// 	fmt.Println("\n callReflectMethod ShowMs didn't run ")
	// }

	result1 := callReflect(&Test{}, method, args...)
	//b := result1[0].MapRange().Value().(Base)
	fmt.Println("result1[0].Bytes()", result1[0].Interface())
	//json.Unmarshal(, &b)

	//fmt.Println("result1", b.Rs, b.Err)

	fmt.Println("\n reflect all ")
}
