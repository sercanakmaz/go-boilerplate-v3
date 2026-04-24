package mongo

import "reflect"

func collectionNameFromType[T any]() string {
	var model T

	if named, ok := any(model).(CollectionNamer); ok {
		return named.CollectionName()
	}

	if named, ok := any(&model).(CollectionNamer); ok {
		return named.CollectionName()
	}

	t := reflect.TypeOf((*T)(nil)).Elem()
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	typeName := t.Name()
	if typeName == "" {
		return ""
	}

	return typeName + "s"
}
