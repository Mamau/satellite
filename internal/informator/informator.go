package informator

import (
	"os"
	"reflect"

	"github.com/gookit/color"
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
			in.Strings[structType.Field(i).Name] = val.Field(i).String()
		case reflect.Int64:
			in.Integers[structType.Field(i).Name] = val.Field(i).Int()
		case reflect.Bool:
			in.Booleans[structType.Field(i).Name] = val.Field(i).Bool()
		case reflect.Slice:
			data := val.Field(i).Slice(0, val.Field(i).Len())
			in.Slices[structType.Field(i).Name] = data.Interface().([]string)
		}
	}
}
