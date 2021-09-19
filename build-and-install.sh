#!/bin/sh

go build -o news cmd/main.go
chmod +x news
mv news /usr/local/bin 