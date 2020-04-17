#!/bin/bash

# Determine the build script's actual directory, following symlinks
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BUILD_DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

# Derive the project name from the directory
PROJECT="$(basename $BUILD_DIR)"

cd $BUILD_DIR
mkdir -p bin 

go build -o bin/$PROJECT main.go

EXIT_STATUS=$?

if [ $EXIT_STATUS == 0 ]; then
    echo "Build succeeded"
else
    echo "Build failed"
fi 

exit $EXIT_STATUS