# !/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzr-labs/photon:ute1777663306 --platform linux/amd64,linux/arm64 --push
