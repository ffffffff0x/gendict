package utils

import (
	"gendict/mod"
)

var resp4 []string
var TempSlice []string

func ListDicFunc(IntputSlice []string) []string {

	if mod.Suffix == "suffix.txt" {
		TempSlice = embedread(mod.Suffix)
	} else {
		TempSlice = FileToSlice(mod.Suffix)
	}

	for _, v1 := range IntputSlice {
		for _, v2 := range TempSlice {
			resp4 = append(resp4, v1+v2)
		}
	}

	return resp4
}
