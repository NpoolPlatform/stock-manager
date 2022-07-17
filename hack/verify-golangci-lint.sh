#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

VERSION=v1.42.1
URL_BASE=https://raw.githubusercontent.com/golangci/golangci-lint
URL=$URL_BASE/$VERSION/install.sh

if [[ ! -f .golangci.yml ]]; then
    echo 'ERROR: missing .golangci.yml in repo root' >&2
    exit 1
fi

if ! command -v gofumpt; then
    go install mvdan.cc/gofumpt@latest
fi

export PATH=$PATH:bin:/tmp/golangci-lint
if ! command -v golangci-lint; then
    mkdir -p /tmp/golangci-lint
    curl -sfL $URL | sh -s $VERSION
    cp bin/golangci-lint /tmp/golangci-lint
fi

golangci-lint version
golangci-lint linters
golangci-lint run "$@"
