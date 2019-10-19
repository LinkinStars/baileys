#!/usr/bin/env bash

go build -o baileys ./src/baileys/cmd/
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o baileys.exe ./src/baileys/cmd/

mkdir baileys-v1.0.0-darwin-amd64
mkdir baileys-v1.0.0-windows-amd64

tar -zcvf ./baileys-v1.0.0-darwin-amd64.tar.gz ./baileys-v1.0.0-darwin-amd64
tar -zcvf ./baileys-v1.0.0-darwin-windows.tar.gz ./baileys-v1.0.0-windows-amd64