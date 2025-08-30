package util

import (
	"fmt"
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

func FormatAttributeValue(attr model.Attribute) string {
	if attr.Percent {
		return fmt.Sprintf("%.1f%%", attr.Value*100)
	}
	return fmt.Sprintf("%d", int(attr.Value))
}

func BuildRelicSummaryOut(r model.Relic) model.RelicSummary {
	main := model.AttributeSummary{
		Name:  r.MainAffix.Name,
		Icon:  NormalizeIconPath(r.MainAffix.Icon),
		Value: FormatAttributeValue(*r.MainAffix),
	}

	subs := make([]model.AttributeSummary, 0, len(r.SubAffix))
	for _, s := range r.SubAffix {
		subs = append(subs, model.AttributeSummary{
			Name:  s.Name,
			Icon:  NormalizeIconPath(s.Icon),
			Value: FormatAttributeValue(s),
		})
	}

	return model.RelicSummary{
		Name:      r.Name,
		Type:      r.Type,
		Icon:      NormalizeIconPath(r.Icon),
		Rarity:    r.Rarity,
		Level:     r.Level,
		MainAffix: &main,
		SubAffix:  subs,
	}
}
