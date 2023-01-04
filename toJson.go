package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var person []Person

func (p *Person) PrintFieldsByName() {
	val := reflect.ValueOf(*p)
	for i := 0; i < val.Type().NumField(); i++ {
		fmt.Println(val.Type().Field(i).Name)
	}
}

func (p *Person) PrintFieldsByTag() {
	val := reflect.ValueOf(*p)
	for i := 0; i < val.Type().NumField(); i++ {
		fmt.Println(val.Type().Field(i).Tag.Get("json"))
	}
}

func EncodeStructByTag(person []Person) string {
	var test string
	test += "["
	for index, p := range person {
		test += "{"
		val := reflect.ValueOf(p)
		for i := 0; i < val.Type().NumField(); i++ {
			data := reflect.Indirect(val).Field(i)
			tag := val.Type().Field(i).Tag.Get("json")
			data1 := data.Interface().(string)
			if i == (val.Type().NumField() - 1) {
				test += "\"" + tag + "\"" + ": " + "\"" + data1 + "\""
			} else {
				test += "\"" + tag + "\"" + ": " + "\"" + data1 + "\"" + ", "
			}

		}
		if index == (len(person) - 1) {
			test += "}"
		} else {
			test += "},"
		}
		//fmt.Println(val)
	}
	test += "]"
	return test
}

func EncodeStructByName(person []Person) string {
	var test string
	test += "["
	for index, p := range person {
		test += "{"
		val := reflect.ValueOf(p)
		for i := 0; i < val.Type().NumField(); i++ {
			data := reflect.Indirect(val).Field(i)
			tag := val.Type().Field(i).Name
			data1 := data.Interface().(string)
			if i == (val.Type().NumField() - 1) {
				test += "\"" + tag + "\"" + ": " + "\"" + data1 + "\""
			} else {
				test += "\"" + tag + "\"" + ": " + "\"" + data1 + "\"" + ", "
			}

		}
		if index == (len(person) - 1) {
			test += "}"
		} else {
			test += "},"
		}
		//fmt.Println(val)
	}
	test += "]"
	return test
}

func decode(st string) {
	//var p Person

	if strings.HasPrefix(st, "[") && strings.HasSuffix(st, "]") {
		st = strings.ReplaceAll(st, "[", "")
		st = strings.ReplaceAll(st, "]", "")

		objects := strings.Split(st, ",")

		for _, object := range objects {
			fmt.Println(object)
		}

	} else if strings.HasPrefix(st, "{") {

	} else {
		fmt.Println("Invalid string")
	}

}

func main() {

	person1 := &Person{"1", "James", "Gunn"}
	person2 := &Person{"2", "John", "Smith"}
	person = append(person, *person1)
	person = append(person, *person2)
	//fmt.Println(person)

	// p := reflect.TypeOf(*person1)

	// p1, _ := p.FieldByName("ID")

	// fmt.Println(p1)

	// v, ok := p1.Tag.Lookup("json")
	// fmt.Printf("%s, %t\n", v, ok)
	// fmt.Println(v, ok)

	// person1.PrintFieldsByName()
	// person1.PrintFieldsByTag()

	jsonTag := EncodeStructByTag(person)
	jsonName := EncodeStructByName(person)

	fmt.Println(jsonName)
	fmt.Println(jsonTag)

	//decode("[{\"ID\": \"1\", \"FirstName\": \"James\", \"LastName\": \"Gunn\"},{\"ID\": \"2\", \"FirstName\": \"John\", \"LastName\": \"Smith\"}]")

}
