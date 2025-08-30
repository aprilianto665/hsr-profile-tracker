package util

const BaseIconURL = "https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/"

func NormalizeIconPath(path *string) {
	*path = BaseIconURL + *path
}
