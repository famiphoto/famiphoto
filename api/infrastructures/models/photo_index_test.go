package models

import (
	"reflect"
	"testing"
)

// TestPhotoElasticSearchMappingMatchesPhotoIndex tests that the mapping returned by PhotoElasticSearchMapping
// matches the fields in the PhotoIndex struct.
func TestPhotoElasticSearchMappingMatchesPhotoIndex(t *testing.T) {
	// Get the mapping
	mapping := PhotoElasticSearchMapping()

	// Check that all top-level fields in PhotoIndex have corresponding properties in the mapping
	photoIndexType := reflect.TypeOf(PhotoIndex{})
	for i := 0; i < photoIndexType.NumField(); i++ {
		field := photoIndexType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue // Skip fields without json tags
		}

		// Check if the field exists in the mapping
		_, exists := mapping.Properties[jsonTag]
		if !exists {
			t.Errorf("Field %s (json: %s) exists in PhotoIndex but not in mapping", field.Name, jsonTag)
		}
	}

	// Check that all properties in the mapping have corresponding fields in PhotoIndex
	for propName := range mapping.Properties {
		found := false
		for i := 0; i < photoIndexType.NumField(); i++ {
			field := photoIndexType.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == propName {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Property %s exists in mapping but not in PhotoIndex", propName)
		}
	}

	// The test passes if all fields in PhotoIndex have corresponding properties in the mapping
	// and all properties in the mapping have corresponding fields in PhotoIndex
}