package utils

import (
	"flag"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

var resp []string

func OrgDicFunc(IntputSlice []string) []string {

	args := pinyin.NewArgs()

	for _, v := range IntputSlice {
		var OrgSlice1 []string
		var OrgSlice2 []string
		var OrgStr1 string
		var OrgStr2 string
		var OrgStr3 string
		var OrgStr4 string
		var OrgStr5 string

		hans := flag.Args()
		hans = append(hans, v)

		args.Fallback = func(r rune, a pinyin.Args) []string {
			return []string{string(r)}
		}

		args.Style = pinyin.Normal
		for _, s := range pinyin.Pinyin(strings.Join(hans, ""), args) {
			OrgSlice1 = append(OrgSlice1, strings.Join(s, ","))
			OrgStr1 = OrgStr1 + strings.Join(s, ",")
			//fmt.Print(strings.Join(s, ","), " ")
		}

		args.Style = pinyin.FirstLetter
		for _, s := range pinyin.Pinyin(strings.Join(hans, ""), args) {
			OrgSlice2 = append(OrgSlice2, strings.Join(s, ","))
			OrgStr2 = OrgStr2 + strings.Join(s, ",")
			//fmt.Print(strings.Join(s, ","), " ")
		}

		OrgStr3 = strings.ToUpper(string(OrgStr1[0])) + OrgStr1[1:]
		OrgStr4 = strings.ToUpper(string(OrgStr2[0])) + OrgStr2[1:]
		OrgStr5 = strings.ToUpper(OrgStr2)

		// fmt.Println(v)
		// baidu
		// fmt.Println("OrgStr1:", OrgStr1)
		// bd
		// fmt.Println("OrgStr2:", OrgStr2)
		// Baidu
		// fmt.Println("OrgStr3:", OrgStr3)
		// Bd
		// fmt.Println("OrgStr4:", OrgStr4)
		// BD
		// fmt.Println("OrgStr5:", OrgStr5)

		resp = append(resp, OrgStr1)
		resp = append(resp, OrgStr2)
		resp = append(resp, OrgStr3)
		resp = append(resp, OrgStr4)
		resp = append(resp, OrgStr5)

	}

	return resp
}
