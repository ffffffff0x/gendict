package mod

import "flag"

var (
	Version = "v1.0.0"
	V       bool
	Org     string
	Name    string
	List    string
	Orgout  string
	Nameout string
	Listout string
	Suffix  string
)

func init() {

	flag.BoolVar(&V, "version", false, "显示版本号")
	flag.StringVar(&Org, "org", "", "组织名称列表")
	flag.StringVar(&Name, "name", "", "用户名称列表")
	flag.StringVar(&List, "list", "", "批量拼接后缀")
	flag.StringVar(&Orgout, "orgout", "orgresult.txt", "组织名称结果输出")
	flag.StringVar(&Nameout, "nameout", "nameresult.txt", "用户名称结果输出")
	flag.StringVar(&Listout, "listout", "listresult.txt", "批量拼接结果输出")
	flag.StringVar(&Suffix, "suffix", "suffix.txt", "自定义后缀")

}
