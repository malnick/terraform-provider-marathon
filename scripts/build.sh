#!/bin/bash
set -e

SOURCE_DIR=$(git rev-parse --show-toplevel)
REVISION=$(git rev-parse --short HEAD)
BUILD_DIR="${SOURCE_DIR}/build"
OS=$1

function build_cmd {
	GOOS=$OS go build -a -o ${BUILD_DIR}/terraform-marathon-provider-${OS}-${REVISION} ${SOURCE_DIR}/main.go
}

function main {
    build_cmd
    if [ -n "$(which tree)" ]; then
        tree build/
    fi
}

main "$@"
