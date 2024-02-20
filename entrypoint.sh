#!/usr/bin/env bash

set -e

function main() {
    cd ${ZENDB_HOME} || exit
    ./zendb
}

main
