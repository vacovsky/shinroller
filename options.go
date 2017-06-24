package main

import (
	"flag"
	"strings"
)

type flagsModel struct {
	inputPath           *string
	outputPath          *string
	bufferCharsLeft     *string
	bufferCharsRight    *string
	tokenizedoutputPath *string
	tokenize            *bool
	extract             *bool
	converge            *string
	keyNames            []string
	valueNames          []string
	excluded            []string
}

func flagInit() flagsModel {
	model := flagsModel{}

	// buffer stuff
	model.bufferCharsLeft = flag.String("buffer-left", "", "Characters used to buffer the keys within the input file on the left side of a token key.  The default value is an empty string.  Example: \"{{mykey\" (usually used in conjunction with --buffer-left).")
	model.bufferCharsRight = flag.String("buffer-right", "", "Characters used to buffer the keys within the input file on the right side of a token key.  The default value is an empty string.  Example: \"mykey}}\" (usually used in conjunction with --buffer-left).")

	// input for collecting tokens
	model.inputPath = flag.String("input", "", "Path to the input files.")
	model.outputPath = flag.String("output-path", "", "Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.")
	model.tokenizedoutputPath = flag.String("tokenized-output-path", "", "Destination path and file name for the tokenized file.  If not set, detokenized file is printed to stdout.")

	// if path is passed to --converge, that is the source for converging the json files
	model.converge = flag.String("converge", "", "Destination path and file name for the detokenized file.  If not set, detokenized file is printed to stdout.")

	// do I tokenize?
	model.tokenize = flag.Bool("tokenize", false, "should I tokenize?")

	// keys and values prefixes.  In the input file, this will be something like <add name="blah" value="someblah">
	keyNamesString := flag.String("key-names", "", "Strings to search for when detecting token keys.  Separate strings with a comma (,) and no spaces.")
	valueNamesString := flag.String("value-names", "", "Strings to search for when detecting token values.  Separate strings with a comma (,) and no spaces.")
	excludedString := flag.String("exclude", "", "Values you wish to ignore if they start with anything in this list.  Separate strings with a comma (,) and no spaces.")

	flag.Parse()

	model.excluded = strings.Split(*excludedString, ",")
	model.keyNames = strings.Split(*keyNamesString, ",")
	model.valueNames = strings.Split(*valueNamesString, ",")
	return model
}
