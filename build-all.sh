#!/bin/sh

FILE=yolo.go
TARGET_FILE=yolo

echo "Building for Windows"
GOOS=windows GOARCH=amd64 go build -o bin/$TARGET_FILE-amd64.exe $FILE
zip -r bin/$TARGET_FILE-amd64.zip bin/$TARGET_FILE-amd64.exe
echo "Building for OSX"
GOOS=darwin GOARCH=amd64 go build -o bin/$TARGET_FILE-amd64-darwin $FILE
zip -r bin/$TARGET_FILE-amd64-darwin.zip bin/$TARGET_FILE-amd64-darwin
echo "Building for Linux"
GOOS=linux GOARCH=amd64 go build -o bin/$TARGET_FILE-amd64-linux $FILE
zip -r bin/$TARGET_FILE-amd64-linux.zip bin/$TARGET_FILE-amd64-linux
