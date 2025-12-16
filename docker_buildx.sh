#!/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzerbeam-studios/spectra:ute1765858665 --platform linux/amd64,linux/arm64 --push
