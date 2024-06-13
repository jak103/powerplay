package validators

import (
	"reflect"
	"strings"
)

// IsValidSortField checks if the given field is a valid sort field in the provided model type.
func IsValidSortField(field string, modelType reflect.Type) bool {
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if modelType.Kind() != reflect.Struct {
		panic("IsValidSortField: expected a struct or pointer to a struct")
	}

	lowerField := strings.ToLower(field) // Convert field to lowercase

	// Traverse through the fields of the struct
	for i := 0; i < modelType.NumField(); i++ {
		structField := modelType.Field(i)

		if structField.Anonymous {
			// Check fields in the embedded struct
			embeddedFieldType := structField.Type
			if embeddedFieldType.Kind() == reflect.Ptr {
				embeddedFieldType = embeddedFieldType.Elem()
			}
			if embeddedFieldType.Kind() == reflect.Struct {
				for j := 0; j < embeddedFieldType.NumField(); j++ {
					embeddedField := embeddedFieldType.Field(j)
					jsonTag := embeddedField.Tag.Get("json")
					jsonField := strings.Split(jsonTag, ",")[0]
					if strings.ToLower(jsonField) == lowerField {
						return true
					}
				}
			}
		} else {
			jsonTag := structField.Tag.Get("json")
			jsonField := strings.Split(jsonTag, ",")[0]
			if strings.ToLower(jsonField) == lowerField {
				return true
			}
		}
	}
	return false
}
