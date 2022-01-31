#!/bin/bash


go run *.go \
    --bind=localhost:4000 \
    --hostname=localhost \
    --http \
    --target=localhost:8080 \
    --verbose
