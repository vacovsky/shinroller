# Shinroller - a tokenization tool for web.config files

A tool for extracting, tokenizing, and sorting key/value pairs from DotNet (web.config) configuration files.

This tools is intended to be used alongside https://github.com/vacoj/fasturtle.  They perform comlimentary tasks, but keeping them separate seemed the best idea.

<a href="https://github.com/vacoj/fasturtle">Fasturtle</a> uses the products (tokenized configs and json values files) of shinroller to detokenize.  I have used this app combination successfully for these tasks.

## Installation

```bash
go get -u github.com/cvacoj/shinroller
go get -u github.com/vacoj/fasturtle # recommended partner app
```

## Usage

### A bash script and usage explained (uncommented bash script is TestingData/roll-shins.sh)

```bash
#!/bin/bash
for f in inputs/* ; do
    ./shinroller \

    # file which you want to be tokenized.
    --input="$f" \

    # where we want the json (unconverged) token file to be placed
    --output-path="tokens/$(basename $f).json" \

    # where we want the tokenized versions of the input to be placed.
    # Don't make this the same directry as your input or it will be overwritten
    --tokenized-output-path="tokenized_files/$(basename $f)" \

    # extractes "something" because "key" is a valid prefix <add key="something" value="myvalue">
    --key-names="name,key" \

    # extractes "blahblahblah" because "value" is a valid prefix <add key="something" value="blahblahblah">
    --value-names="hostAddress,connectionString,value" \

    # skips over <add key="something" value="Microsoft.Something.etc">
    --exclude="Microsoft.,System." \

    # chars to use as buffers around the tokens
    --buffer-left="{{" \
    --buffer-right="}}"
done

# this will converge and sort all token files, so that any files which
# share key-values have that value moved to an "environment.json" file,
# and any unique key/value pairs stay in the json for the original input
# token file.
shinroller \
    # location of json token files from previous loop
    --converge="tokens/" \

    # location where you want to place the converged tokens.  This is
    # essentially the final product.
    --output-path="converged_tokens/"

```

### Folder Structure

```text
TestingData
├── converged_tokens
│   ├── environment.json
│   ├── input1.xml.json
│   └── input2.xml.json
├── inputs
│   ├── input1.xml    <---- the only original files
│   └── input2.xml    <---- the only original files
├── roll-shins.sh
├── tokenized_files
│   ├── input1.xml
│   └── input2.xml
└── tokens
    ├── input1.xml.json
    └── input2.xml.json
```
