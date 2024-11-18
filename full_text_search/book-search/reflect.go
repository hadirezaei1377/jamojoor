package booksearch

import (
	"fmt"
	"reflect"
)

func setField(obj interface{}, fieldName string, newValue interface{}) {
	v := reflect.ValueOf(obj).Elem()
	f := v.FieldByName(fieldName)

	if f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(newValue))
	}
}

type Person struct {
	Name string
	Age  int
}

func main3() {
	p := &Person{Name: "Ali", Age: 25}
	fmt.Println("Before:", p)

	setField(p, "Name", "Boby")
	setField(p, "Age", 30)

	fmt.Println("After:", p)
}
