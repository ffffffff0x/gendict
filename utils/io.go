package utils

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
)

//go:embed suffix.txt
var local embed.FS

func FileToSlice(filePath string) []string {
	var resp2 []string

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		resp2 = append(resp2, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return resp2
}

func embedread(filePath string) []string {
	var resp5 []string

	f, err := local.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// do something with a line
		resp5 = append(resp5, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return resp5

}

func OutputTxt(filename string, dic []string) {

	filePath := filename
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	for _, v := range dic {
		write.WriteString(v + "\n")
	}

	write.Flush()

	fmt.Println("结果已输出到: ", filename)

}
