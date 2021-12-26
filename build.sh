#!/bin/bash

amazon-linux-extras install golang1.11
GOARCH=amd64 GOOS=linux go build output/server/pages/main.go -ldflags="-s -w"
mv output .output