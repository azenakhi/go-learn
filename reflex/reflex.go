package reflex

import "reflect"

import "fmt"

type Person struct {
	Name string `my_tag:"bob"`
	Age  string `my_tag:"20years"`
}

func (p *Person) Reflect() string {
	values := reflect.ValueOf(p).Elem()
	fields := reflect.TypeOf(p).Elem()
	str := ""
	for i := 0; i < values.NumField(); i++ {
		//value := values.Field(i)
		field := fields.Field(i)
		tag := field.Tag.Get("my_tag")
		str += fmt.Sprintf("%s = %s ", field.Name, tag)
	}
	return str
}
