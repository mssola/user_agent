#!/bin/bash
# Copyright (C) 2012-2021 Miquel Sabaté Solà <mikisabate@gmail.com>
# This file is licensed under the MIT license.
# See the LICENSE file.

# Hack so we work around old versions of Go which are no longer supported by
# linting tools that we use in this project.

case "${TRAVIS_GO_VERSION:2:1}" in
    "4" | "5" | "6" | "7" | "8")
        echo "This version of Go is not going to use linting tools."
        ;;
    *)
        go get -u github.com/vbatts/git-validation
        go get -u golang.org/x/lint/golint
        ;;
esac
