#!/bin/bash

# Fail if no argument is provided
if [ $# -eq 0 ]; then
    echo "No arguments provided"
    exit 1
fi


# Get the problem name from the url
url=$1
url=${url%/}
dir_name=$(basename $(dirname $url))
echo $dir_name
mkdir "$dir_name"
cd "$dir_name" || exit

# Copy template file to the directory
cp ../app.go .
cp ../app_test.go .

# init go project
go mod init app
go mod tidy
