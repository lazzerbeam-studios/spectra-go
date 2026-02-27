#!/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzr-labs/spectra:ute1769640310 --platform linux/amd64,linux/arm64 --push
