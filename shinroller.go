package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"bytes"

	"strings"

	"regexp"
)

func main() {
	args := flagInit()
	tokens := [][]byte{}

	if *args.converge != "" {
		convergeFromPath(*args.converge, *args.outputPath)

	} else {
		var output []byte

		ensureFileExists(*args.inputPath, "--input")
		input := loadFile(*args.inputPath)
		tokens = append(tokens, extractTokenNames(input, "("+strings.Join(args.keyNames, "|")+")=\".*\"\\s*?("+strings.Join(args.valueNames, "|")+")=\".*\"")...)

		// spew.Dump(strings.Join(args.keyNames, "|"))
		// spew.Dump(strings.Join(args.valueNames, "|"))
		// spew.Dump(strings.Join(args.excluded, "|"))

		cleanTokens := [][]byte{}
		for _, t := range tokens {
			r1 := regexp.MustCompile("(" + strings.Join(args.keyNames, "|") + ")=\"")
			r2 := regexp.MustCompile("\"\\s*?(" + strings.Join(args.valueNames, "|") + ")=\".*\"")

			t = r1.ReplaceAll(t, []byte(""))
			t = r2.ReplaceAll(t, []byte(""))

			cleanTokens = append(cleanTokens, t)
		}
		tokenValueMap := map[string]string{}
		for _, t := range cleanTokens {
			for _, argggg := range tokens {
				re := fmt.Sprintf("("+strings.Join(args.keyNames, "|")+")=\"%s\"\\s*("+strings.Join(args.valueNames, "|")+")=", t)
				re3 := regexp.MustCompile(re)
				if re3.Match(argggg) {
					argggg = re3.ReplaceAll(argggg, []byte(""))
					argggg = bytes.Replace(argggg, []byte("\""), []byte(""), -1)
					tfound := false
					for _, exclude := range args.excluded {
						if strings.HasPrefix(string(argggg), exclude) {
							tfound = true
						}
					}
					if !tfound {
						tokenValueMap[string(t)] = string(argggg)
					}
				}
			}
		}

		output, _ = json.MarshalIndent(tokenValueMap, "", "    ")

		// send tokens to file
		outputToFile(*args.outputPath, output)

		// send tokenized config to file
		outputToFile(*args.tokenizedoutputPath, []byte(tokenize(input, tokenValueMap)))
	}
}

func tokenize(input []byte, tokenMap map[string]string) string {
	inputS := string(input)
	for k, v := range tokenMap {
		if v != "false" && v != "true" {
			inputS = strings.Replace(inputS, "\""+v+"\"", "__"+k+"__", 1)
		}
	}
	return inputS
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func loadFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	checkError(err)
	return file
}

func outputToFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0644)
	checkError(err)
}

func ensureFileExists(file, use string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("Error: File \"%s\" does not exist. Please provide a valid file path for %s.\n", file, use)
		os.Exit(1)
	}
}

func outputToStdout(data []byte) {
	fmt.Print(string(data))
}
