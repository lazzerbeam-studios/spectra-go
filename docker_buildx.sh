#!/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzerbeam-studios/spectra:ute1765860480 --platform linux/amd64,linux/arm64 --push
