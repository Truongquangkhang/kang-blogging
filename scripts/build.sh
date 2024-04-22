#!/bin/bash
set -e

readonly service="$1"

cd "./internal/$service"
go get .
go build -o ./app-runtime .
