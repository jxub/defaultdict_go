package main

import (
	"log"
	"reflect"
)

// defaultType is a type alias for clearer semantics IMHO
type defaultType interface{}

// DefaultDict is the container that replicates
// `defaultdict` functionality
type DefaultDict struct {
	// DefaultFactory can be any element of any type
	// although it's better to use the same elements as those
	// returned by default Python constructors, and by extension here
	DefaultFactory defaultType
	// actual `dict`
	Content map[interface{}]interface{}
}

// NewDefaultDict creates a map with the `def` argument for missing values
func NewDefaultDict(def interface{}) *DefaultDict {
	cont := make(map[interface{}]interface{})
	return &DefaultDict{
		DefaultFactory: def,
		Content:        cont}
}

// Get gets the value from the map, or returns the default key for this type
func (d *DefaultDict) Get(key interface{}) defaultType {
	val, ok := d.Content[key]
	if !ok {
		val := d.getDefault()
	}
	return val
}

// Set sets the value for the given key
// returns error if the type isn't the default one
func (d *DefaultDict) Set(key interface{}, value interface{}) error {
	err := checkDefault(value)
	if err != nil {
		return error
	}
	d.Content[key] = value

	return nil
}

// checkDefault checks it the type of the argument corresponds
// with the type of the DefaultFactory, and returns an error if not
func (d *DefaultDict) checkDefault(value interface{}) error {
	v := reflect.ValueOf(value)
	dv := reflect.ValueOf(d.DefaultFactory)
	if v.Kind() != dv.Kind() {
		return error
	}

	return nil
}

// getDefault returns the default value for a type, following
// the Python defaultdict API: https://docs.python.org/3/library/collections.html#collections.defaultdict
func (d *DefaultDict) getDefault() defaultType {
	t := reflect.ValueOf(d.DefaultFactory)

	switch t.Kind() {
	// bool
	case reflect.Bool:
		return false
	// int types
	case reflect.Int:
		return int(0)
	case reflect.Int8:
		return int8(0)
	case reflect.Int32:
		return int32(0)
	case reflect.Int64:
		return int64(0)
	// uint types
	case reflect.Uint:
		return uint(0)
	case reflect.Uint8:
		return uint8(0)
	case reflect.Uint32:
		return uint32(0)
	case reflect.Uint64:
		return uint64(0)
	// float types
	case reflect.Float32:
		return float32(0)
	case reflect.Float64:
		return float64(0)
	// string
	case reflect.String:
		return string("")
	// more complex types created with make
	case reflect.Slice:
		s := make([]interface{}, 0)
		return s
	case reflect.Map:
		m := make(map[interface{}]interface{})
		return m
	case reflect.Chan:
		c := make(chan interface{}, 0)
		return c
	// return nil if we can't get type by reflection
	default:
		return nil
	}
}

func main() {
	log.Fatal("this package is intended to be imported, not run!")
}
