# !/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzr-labs/photon:ute1775062971 --platform linux/amd64,linux/arm64 --push
