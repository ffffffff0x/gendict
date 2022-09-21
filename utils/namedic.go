package utils

import (
	"flag"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

var resp3 []string

func NameDicFunc(IntputSlice []string) []string {

	args := pinyin.NewArgs()
	for _, v := range IntputSlice {
		var NameSlice1 []string
		var NameSlice2 []string
		var NameStr1 string
		var NameStr2 string
		var NameStr3 string
		var NameStr4 string
		var NameStr5 string
		var NameStr6 string
		var NameStr7 string
		var NameStr8 string
		var NameStr9 string
		var NameStr10 string

		if is_all_chinese(v) {
			hans := flag.Args()
			hans = append(hans, v)

			args.Fallback = func(r rune, a pinyin.Args) []string {
				return []string{string(r)}
			}

			args.Style = pinyin.Normal
			for _, s := range pinyin.Pinyin(strings.Join(hans, ""), args) {
				NameSlice1 = append(NameSlice1, strings.Join(s, ","))
				NameStr1 = NameStr1 + strings.Join(s, ",")
				//fmt.Print(strings.Join(s, ","), " ")
			}

			args.Style = pinyin.FirstLetter
			for _, s := range pinyin.Pinyin(strings.Join(hans, ""), args) {
				NameSlice2 = append(NameSlice2, strings.Join(s, ","))
				NameStr2 = NameStr2 + strings.Join(s, ",")
				//fmt.Print(strings.Join(s, ","), " ")
			}

			NameStr3 = NameSlice1[0] + NameStr2[1:]

			for _, s := range NameSlice2[:len(NameSlice2)-1] {
				NameStr4 = NameStr4 + s
			}
			NameStr4 = NameStr4 + NameSlice1[len(NameSlice2)-1]

			for _, s := range NameSlice1[1:] {
				NameStr5 = NameStr5 + s
			}
			NameStr6 = NameStr5 + NameSlice2[0]
			NameStr10 = NameStr5 + "." + NameSlice1[0]
			NameStr5 = NameStr5 + NameSlice1[0]

			for _, s := range NameSlice2[1:] {
				NameStr7 = NameStr7 + s
			}
			NameStr7 = NameStr7 + NameSlice2[0]

			NameStr8 = strings.ToUpper(string(NameStr2[0])) + NameStr2[1:]
			NameStr9 = strings.ToUpper(NameStr2)

			//fmt.Println(v)
			// zhangsanfeng
			//fmt.Println("NameStr1:", NameStr1)
			// zsf
			//fmt.Println("NameStr2:", NameStr2)
			// zhangsf
			//fmt.Println("NameStr3:", NameStr3)
			// zsfeng
			//fmt.Println("NameStr4:", NameStr4)
			// sanfengzhang
			//fmt.Println("NameStr5:", NameStr5)
			// sanfengz
			//fmt.Println("NameStr6:", NameStr6)
			// sfz
			//fmt.Println("NameStr7:", NameStr7)
			// ZSF
			//fmt.Println("NameStr8:", NameStr8)
			// Zsf
			//fmt.Println("NameStr9:", NameStr9)
			// sanfeng.zhang
			//fmt.Println("NameStr10:", NameStr10)

			resp3 = append(resp3, NameStr1)
			resp3 = append(resp3, NameStr2)
			resp3 = append(resp3, NameStr3)
			resp3 = append(resp3, NameStr4)
			resp3 = append(resp3, NameStr5)
			resp3 = append(resp3, NameStr6)
			resp3 = append(resp3, NameStr7)
			resp3 = append(resp3, NameStr8)
			resp3 = append(resp3, NameStr9)
			resp3 = append(resp3, NameStr10)

		} // 拼音暂时不支持
	}

	return resp3
}
