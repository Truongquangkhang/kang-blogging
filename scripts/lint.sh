#!/bin/bash
set -e

readonly service="$1"

cd "./internal/$service"
golangci-lint run
