package utils

import "reflect"

func StructTags(stru interface{}, tag string) (tags []string) {
	t := reflect.TypeOf(stru).Elem()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(tag) == "-" || t.Field(i).Tag.Get(tag) == "" {
			continue
		}
		tags = append(tags, t.Field(i).Tag.Get(tag))
	}
	return tags
}
