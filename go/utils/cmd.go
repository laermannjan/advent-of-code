package utils

import (
	"path"
)

func GetYearDay(pkg_path string) (int, int) {
	dir, file := path.Split(pkg_path)
	year := Atoi(path.Base(dir)[4:])
	day := Atoi(file[3:])
	return year, day
}
