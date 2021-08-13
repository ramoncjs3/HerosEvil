#!/usr/bin/env zsh

echo "build macos x64 HerosEvil..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o release/HerosEvil_macos_x64
echo "build windows x64 HerosEvil..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o release/HerosEvil_windows_x64.exe
echo "build linux x64 HerosEvil..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o release/HerosEvil_linux_x64

