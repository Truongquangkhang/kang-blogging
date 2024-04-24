#!/bin/bash
set -e

readonly service="$1"

cd "./internal/$service"
go build -o ./app-runtime .
