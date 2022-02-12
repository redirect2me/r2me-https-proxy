#!/bin/bash


set -o errexit
set -o pipefail
set -o nounset

~/go/bin/air

go run *.go \
    --bind=localhost:4000 \
    --hostname=localhost \
    --http \
    --target=localhost:8080 \
    --verbose
