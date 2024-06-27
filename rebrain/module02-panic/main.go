package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inFilePath := "C:\\CODE\\GO\\GO_stydy\\rebrain\\module02-panic\\in.txt"
	inFile, err := os.Open(inFilePath)
	if err != nil {
		fmt.Println("load file eror: ", err)
	}
	defer inFile.Close()

	outFilePath := "C:\\CODE\\GO\\GO_stydy\\rebrain\\module02-panic\\out.txt"
	outFile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println("error create file: ", err)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		p0 := parts[0]
		p1 := parts[1]
		p2 := parts[2]
		row++

		if len(p0) > 0 && len(p1) > 0 && len(p2) > 0 {
			outString := fmt.Sprintf("%-4d %-30s %-40s %-30s\n", row, p0, p1, p2)
			_, err := writer.WriteString(outString)
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		} else {
			_, err := writer.WriteString(">>>>>ERROR NO FIELD>>>>>>\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
	}
	writer.Flush()
}
