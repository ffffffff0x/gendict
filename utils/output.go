package utils

import (
	"flag"
	"gendict/mod"
)

var InNum = 0
var OutputSlice []string

func OutputProcess() {
	if mod.Org != "" {
		OutputSlice = OrgDicFunc(FileToSlice(mod.Org))
		OutputTxt(mod.Orgout, OutputSlice)
		InNum += 1
	}
	if mod.Name != "" {
		OutputSlice = NameDicFunc(FileToSlice(mod.Name))
		OutputTxt(mod.Nameout, OutputSlice)
		InNum += 1
	}
	if mod.List != "" {
		OutputSlice = ListDicFunc(FileToSlice(mod.List))
		OutputTxt(mod.Listout, OutputSlice)
		InNum += 1
	}
	if InNum == 0 { // 无参数输出Help
		flag.Usage()
	}
}
