package util

import (
	"hsr-profile-tracker/internal/model"
)

func MergeAttributes(attrs, adds []model.Attribute) []model.Attribute {
	final := make([]model.Attribute, len(attrs))
	copy(final, attrs)
	idx := make(map[string]int, len(final))
	for i, a := range final {
		idx[a.Name] = i
	}
	for _, add := range adds {
		if i, ok := idx[add.Name]; ok {
			final[i].Value += add.Value
		} else {
			idx[add.Name] = len(final)
			final = append(final, add)
		}
	}
	return final
}
