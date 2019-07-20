#!/usr/bin/env bash

file="$1"

if [[ "${IS_CI}" = "true" ]] ; then
    bash <(curl -s https://codecov.io/bash) -f "${file}"
fi
