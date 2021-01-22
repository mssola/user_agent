#!/bin/bash
# Copyright (C) 2012-2021 Miquel Sabaté Solà <mikisabate@gmail.com>
# This file is licensed under the MIT license.
# See the LICENSE file.

# Idea taken from: https://github.com/golang/lint/issues/263#issuecomment-590070899
#
# The thing here is that I named the package with an underscore, and that was
# done way before coding conventions and such were enforced with tools like
# golint, so these kinds of rules were not taken into consideration.
#
# For compatibility reasons I'm wary of renaming this package, so I'll just
# ignore this rule.

# Hack so we work around old versions of Go which are no longer supported by
# linting tools that we use in this project.

case "${TRAVIS_GO_VERSION:2:1}" in
    "4" | "5" | "6" | "7" | "8")
        echo "This version of Go is not going to use golint."
        exit 0
        ;;
esac

which golint >/dev/null 2>/dev/null || (echo "ERROR: golint not found." && false)

# Patterns to be ignored from the go lint output
IGNORED_PATTERNS=(
    "don't use an underscore in package name"
)

# Patterns joined into a regular expression
REGEX=$(printf "|(%s)" "${IGNORED_PATTERNS[@]}")
REGEX=${REGEX:1}

# Execute go lint on all the files and filter output by the regualr expression
output=$( ( (golint ./... | egrep -v "$REGEX") 2>&1 ) | tee /dev/fd/2);
if [ -z "$output" ]
then
    exit 0
else
    exit 1
fi
