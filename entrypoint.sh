#!/usr/bin/env bash

set -e

function main() {
    cd ${CANODB_HOME} || exit
    ./canodb
}

main
