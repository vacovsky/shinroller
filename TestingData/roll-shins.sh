#!/bin/bash

for f in inputs/* ; do
    ./shinroller \
    --input="$f" \
    --output-path="tokens/$(basename $f).json" \
    --tokenized-output-path="tokenized_files/$(basename $f)" \ 
    --key-names="name,key" \
    --value-names="hostAddress,connectionString,value" \
    --exclude="Microsoft.,System."
done

./shinroller \
    --converge="tokens/" \
    --output-path="converged_tokens/"
