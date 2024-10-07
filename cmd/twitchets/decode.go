package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// DecodeHookFuncStructField is a function like mapstructure's DecodeHookFunc
// but it has complete information about the source type and the target field.
// As the target must be a struct field, this hook will only be invoked when the
// target is a struct field.
// This can be convert into a DecodeHookFunc compatible with mapstructure using ...
type DecodeHookFuncStructField func(reflect.StructField, reflect.Type, any) (any, error)

func ConvertToDecodeHookFunc(hooks ...DecodeHookFuncStructField) mapstructure.DecodeHookFunc {
	return func(
		source reflect.Type,
		target reflect.Type,
		data any,
	) (any, error) {
		if target.Kind() != reflect.Struct {
			return data, nil
		}

		for i := 0; i < target.NumField(); i++ {
			targetField := target.Field(i)
			for _, hook := range hooks{
				hook(source, targetField, data)
			}
		}

}

func DecodeHookFuncField() mapstructure.DecodeHookFuncType {
	return func(
		source reflect.Type,
		target reflect.Type,
		data any,
	) (any, error) {
		switch source.Kind() {
		case reflect.Map, reflect.Struct, reflect.Array:
			return data, nil
		}

		if target.Kind() != reflect.Struct {
			return data, nil
		}

		// Find the field with the mapstructure shorthand tag
		var shorthandFieldName string
		var shorthandFieldType reflect.Type
		for i := 0; i < target.NumField(); i++ {
			field := target.Field(i)

			// Get mapstructure tag value
			tagValue, ok := field.Tag.Lookup("mapstructure")
			if !ok {
				continue
			}

			// Get index of first comma - anything before first comma is the key name
			commaIdx := strings.Index(tagValue, ",")
			if commaIdx == -1 {
				continue
			}

			// Check if "shorthand" is in the mapstructure tag (after first comma)
			if !strings.Contains(tagValue[commaIdx+1:], "shorthand") {
				continue
			}

			if shorthandFieldName != "" {
				return nil, fmt.Errorf("multiple shorthand tags found in struct")
			}

			shorthandFieldName = field.Name
			shorthandFieldType = field.Type
		}
		if shorthandFieldName == "" {
			return data, nil
		}

		// Check data type matches the target field type
		dataValue := reflect.ValueOf(data)
		dataType := dataValue.Type()
		if dataType != shorthandFieldType {
			return data, nil
		}

		// Get pointer to new value of target type, and set the value
		t := reflect.New(target).Elem()
		t.FieldByName(shorthandFieldName).Set(dataValue)

		return t.Interface(), nil
	}
}

func ShorthandHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		source reflect.Type,
		target reflect.Type,
		data any,
	) (any, error) {
		switch source.Kind() {
		case reflect.Map, reflect.Struct, reflect.Array:
			return data, nil
		}

		if target.Kind() != reflect.Struct {
			return data, nil
		}

		// Find the field with the mapstructure shorthand tag
		var shorthandFieldName string
		var shorthandFieldType reflect.Type
		for i := 0; i < target.NumField(); i++ {
			field := target.Field(i)

			// Get mapstructure tag value
			tagValue, ok := field.Tag.Lookup("mapstructure")
			if !ok {
				continue
			}

			// Get index of first comma - anything before first comma is the key name
			commaIdx := strings.Index(tagValue, ",")
			if commaIdx == -1 {
				continue
			}

			// Check if "shorthand" is in the mapstructure tag (after first comma)
			if !strings.Contains(tagValue[commaIdx+1:], "shorthand") {
				continue
			}

			if shorthandFieldName != "" {
				return nil, fmt.Errorf("multiple shorthand tags found in struct")
			}

			shorthandFieldName = field.Name
			shorthandFieldType = field.Type
		}
		if shorthandFieldName == "" {
			return data, nil
		}

		// Check data type matches the target field type
		dataValue := reflect.ValueOf(data)
		dataType := dataValue.Type()
		if dataType != shorthandFieldType {
			return data, nil
		}

		// Get pointer to new value of target type, and set the value
		t := reflect.New(target).Elem()
		t.FieldByName(shorthandFieldName).Set(dataValue)

		return t.Interface(), nil
	}
}
