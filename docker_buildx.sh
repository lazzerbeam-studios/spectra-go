#!/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzr-labs/photon:ute1774887950 --platform linux/amd64,linux/arm64 --push
