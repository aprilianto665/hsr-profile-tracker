package util

import "hsr-profile-tracker/internal/model"

const BaseIconURL = "https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/"

func NormalizeIconPath(path string) string {
	return BaseIconURL + path
}

func NormalizePlayerAvatar(p model.Player) model.Player {
	p.Avatar.Icon = NormalizeIconPath(p.Avatar.Icon)
	return p
}

func NormalizeRelicSetIcons(sets []model.RelicSet) []model.RelicSet {
	for i := range sets {
		sets[i].Icon = NormalizeIconPath(sets[i].Icon)
	}
	return sets
}
