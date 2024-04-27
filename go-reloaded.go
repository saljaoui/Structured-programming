package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage ¯_(ツ)_/¯: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFileName := args[0]
	outputFileName := args[1]

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		contentField1 := strings.Fields(line)
		contentField1 = BinToDec(contentField1)
		contentField1 = HexToDec(contentField1)
		contentField1 = ToUp(contentField1)
		contentField1 = ToLow(contentField1)
		contentField1 = Capitalize(contentField1)
		contentField1 = Ponctuation(contentField1)
		contentField1 = QuotesCheck(contentField1)
		contentField1 = CheckA(contentField1)
		line = strings.Join(contentField1, " ")
		outputFile.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Check if the last line is only a newline character
	fileInfo, err := outputFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.Size() > 0 {
		_, err = outputFile.Seek(-1, io.SeekEnd)
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 1)
		_, err := outputFile.Read(buf)
		if err == nil && buf[0] == '\n' {
			// If the last line is only a newline character, delete it
			err := outputFile.Truncate(fileInfo.Size() - 1)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func BinToDec(array []string) []string {
	for i, val := range array {
		if (val == "(bin)" || val == "(BIN)") && i == 0 {
			array[i] = ""
			fmt.Println("plz enter something before (bin)")
		} else if (val == "(bin)" || val == "(BIN)") && i > 0 {
			if intValue, err := strconv.ParseInt(array[i-1], 2, 64); err == nil {
				array[i-1] = strconv.Itoa(int(intValue))
				array[i] = ""
			} else {
				fmt.Printf("Conversion failed (ㆆ_ㆆ): %s\n", err)
			}
		}
	}
	return TrimEmpties(array)
}

func HexToDec(array []string) []string {
	for i, val := range array {
		if (val == "(hex)" || val == "(HEX)") && i == 0 {
			array[i] = ""
			fmt.Println("plz enter something before (hex)")
		} else if (val == "(hex)" || val == "(HEX)") && i > 0 {
			if intValue, err := strconv.ParseInt(array[i-1], 16, 64); err == nil {
				array[i-1] = strconv.Itoa(int(intValue))
				array[i] = ""
			} else {
				fmt.Printf("Conversion failed (ㆆ_ㆆ): %s\n", err)
			}
		}
	}
	return TrimEmpties(array)
}

func ToUp(array []string) []string {
	for i, val := range array {
		if (val == "(up," || val == "(UP,") && i == 0 {
			array[i] = ""
			fmt.Println("plz enter something before (up,")
		} else if (val == "(up)" || val == "(UP)" ) && i == 0 {
			array[i] = ""
			fmt.Println("plz enter something before (up)")
		} else if (val == "(up)" || val == "(up," || val == "(UP)" || val == "(UP,") && i > 0 {
			if val == "(up)" || val == "(UP)" {
				array[i-1] = strings.ToUpper(array[i-1])
			} else {
				parameter := toInt(array[i+1][:len(array[i+1])-1])
				if parameter > len(array) {
					fmt.Println("Error with length (╯°□°)╯!")
					os.Exit(1)
				}
				for j := i - parameter; j < i; j++ {
					array[j] = strings.ToUpper(array[j])
				}
				array[i+1] = ""
			}
			array[i] = ""
		}
	}
	return TrimEmpties(array)
}

func ToLow(array []string) []string {
	for i, val := range array {
		if (val == "(low," || val == "(LOW,") && i == 0 {
			array[i] = ""
			fmt.Println("plz enter something before (low,")
		} else if (val == "(low)" || val == "(LOW)" ) && i == 0 {
			array[i] = ""
			fmt.Println("plz enter something before (low)")
		} else if (val == "(low)" || val == "(low," || val == "(LOW)" || val == "(LOW,") && i > 0 {
			if val == "(low)" || val == "(LOW)" {
				array[i-1] = strings.ToLower(array[i-1])
			} else {
				parameter := toInt(array[i+1][:len(array[i+1])-1])
				if parameter > len(array) {
					fmt.Println("Error with length (╯°□°)╯!")
					os.Exit(1)
				}
				for j := i - parameter; j < i; j++ {
					array[j] = strings.ToLower(array[j])
				}
				array[i+1] = ""
			}
			array[i] = ""
		}
	}
	return TrimEmpties(array)
}

func Capitalize(array []string) []string {
	for i, val := range array {
		
		if (val == "(cap)" || val == "(cap,") && i > 0 {
			if val == "(cap)" {
				array[i-1] = TitleOfUs(array[i-1])
			} else {
				parameter := toInt(array[i+1][:len(array[i+1])-1])
				if parameter > len(array) {
					fmt.Println("Error with length (╯°□°)╯!")
					os.Exit(1)
				}
				for j := i - parameter; j < i; j++ {
					array[j] = TitleOfUs(array[j])
				}
				array[i+1] = ""
			}
			array[i] = ""
		}
	}
	return TrimEmpties(array)
}

func TitleOfUs(s string) string {
	str := ""
	if len(s) > 0 && (s[0] >= 'a' && s[0] <= 'z') {
		str += string(s[0] - 32)
	} else if len(s) > 0 && !(s[0] >= 'a' && s[0] <= 'z') {
		str += string(s[0])
	}
	s = s[1:]
	for i := 0; i < len(s); i++ {
		if !(unicode.IsLetter(rune(s[i]))) {
			str += string(s[i])
		} else if len(s) > 0 && (s[i] >= 'a' && s[i] <= 'z') {
			str += string(s[i])
		} else if len(s) > 0 && !(s[0] >= 'a' && s[0] <= 'z') {
			str += string(s[i] + 32)
		}
	}
	return str
}

func Ponctuation(array []string) []string {
	puncs := []string{",", ".", ";", ":", "!", "?", "!?", "..."}
	for i, val := range array {
		if start(val, puncs) && i > 0 {
			currentPunc := array[i]
			if len(val) > 1 && !start(string(val[1]), puncs) {
				array[i] = val[1:]
				array[i-1] += val[:1]
			} else {
				array[i-1] += currentPunc
			}
			array[i] = ""
		}
	}
	return TrimEmpties(array)
}

func QuotesCheck(array []string) []string {
	puncs := []string{"'", "‘", "’", "\""}
	var checker bool
	var wordCount int
	length := len(array) - 1
	for i, val := range array {
		if strings.HasPrefix(array[0], "\"") && i != 0 {
			if strings.HasSuffix(array[i-1], ":") && start(val, puncs) {
				checker = true
			}
		}
		if checker && checkAlpha(val) {
			wordCount++
		}
	}

	for i, val := range array {
		var prev, next string
		current := array[i]
		if i != 0 && i < length {
			prev = array[i-1]
			next = array[i+1]
		}
		if start(val, puncs) && i > 0 && i < length {
			if checkAlpha(prev) && checkAlpha(next) {
				array[i] = prev + current + next
			}
			if strings.HasSuffix(prev, ":") {
				array[i] += next
				array[i+1] = ""
			}
		}
		if start(val, puncs) && i == length && !checkAlpha(prev) {
			if strings.HasPrefix(array[length], "‘") {
				array[length] = "’"
			}
			array[length-1] += array[length]
			array[length] = ""
		}
	}
	if wordCount != 1 && checker {
		array[0] = array[0][1:]
		array[length-1] = array[length-1][:len(array[length-1])-1]
	}
	return TrimEmpties(array)
}

func CheckA(array []string) []string {
	for i, val := range array {
		checks := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U", "h", "H"}
		var next string
		length := len(array) - 1
		if i != 0 && i < length {
			next = array[i+1]
		}
		if val == "a" && start(next, checks) {
			array[i] = "an"
		}
	}
	TrimEmpties(array)
	return array
}

func start(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return n
}

func TrimEmpties(array []string) []string {
	var new_arr []string
	for _, word := range array {
		if word != "" {
			new_arr = append(new_arr, word)
		}
	}
	return new_arr
}

func checkAlpha(s string) bool {
	lastChar := s[len(s)-1]
	if (lastChar > 'a' && lastChar < 'z') || (lastChar > 'A' && lastChar < 'Z') {
		return true
	}
	return false
}
