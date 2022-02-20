package informator

import (
	"github.com/gookit/color"
	"os"
	"reflect"
)

type Informator struct {
	EntityName string
	Strings    map[string]string
	Integers   map[string]int64
	Booleans   map[string]bool
	Slices     map[string][]string
}

func NewInformator(entity interface{}) *Informator {
	informator := Informator{
		Strings:  make(map[string]string),
		Integers: make(map[string]int64),
		Booleans: make(map[string]bool),
		Slices:   make(map[string][]string),
	}
	informator.scanEntity(entity)

	return &informator
}

func (in *Informator) GetAll() map[string]string {
	fields := make(map[string]string)
	for i := range in.Strings {
		fields[i] = "string"
	}
	for i := range in.Integers {
		fields[i] = "integer"
	}
	for i := range in.Booleans {
		fields[i] = "boolean"
	}
	for i := range in.Slices {
		fields[i] = "list"
	}

	return fields
}

func (in *Informator) scanEntity(entity interface{}) {
	val := reflect.ValueOf(entity)

	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	if val.Kind() != reflect.Struct {
		color.Red.Printf("unexpected type\n")
		os.Exit(1)
	}
	structType := val.Type()
	in.EntityName = structType.Name()

	for i := 0; i < val.NumField(); i++ {
		switch val.Field(i).Kind() {
		case reflect.String:
			name := reflect.TypeOf(entity).Field(i).Tag
			in.Strings[name.Get("yaml")] = val.Field(i).String()
		case reflect.Int64:
			name := reflect.TypeOf(entity).Field(i).Tag
			in.Integers[name.Get("yaml")] = val.Field(i).Int()
		case reflect.Bool:
			name := reflect.TypeOf(entity).Field(i).Tag
			in.Booleans[name.Get("yaml")] = val.Field(i).Bool()
		case reflect.Slice:
			data := val.Field(i).Slice(0, val.Field(i).Len())
			name := reflect.TypeOf(entity).Field(i).Tag
			in.Slices[name.Get("yaml")] = data.Interface().([]string)
		}
	}
}
