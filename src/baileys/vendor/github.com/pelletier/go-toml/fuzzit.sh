#!/bin/bash
set -xe

# fuzz only in one configuration
# there's no benefit to fuzzing with different go versions
if [ -z ${WITH_FUZZ} ]; then
    exit 0
fi

# go-fuzz doesn't support modules yet, so ensure we do everything
# in the old style GOPATH way
export GO111MODULE="off"

# install go-fuzz
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

# target name can only contain lower-case letters (a-z), digits (0-9) and a dash (-)
# to add another target, make sure to create it with `fuzzit create target`
# before using `fuzzit create job`
TARGET=toml-fuzzer

go-fuzz-build -libfuzzer -o ${TARGET}.a github.com/pelletier/go-toml
clang -fsanitize=fuzzer ${TARGET}.a -o ${TARGET}

# install fuzzit for talking to fuzzit.dev service
# or latest version:
# https://github.com/fuzzitdev/fuzzit/releases/latest/download/fuzzit_Linux_x86_64
wget -q -O fuzzit https://github.com/fuzzitdev/fuzzit/releases/download/v2.4.23/fuzzit_Linux_x86_64
chmod a+x fuzzit

# upload fuzz target for long fuzz testing on fuzzit.dev server
# or run locally for regression, depending on --type
if [ "${TRAVIS_PULL_REQUEST}" == "false" ]; then
	TYPE=fuzzing
else
	TYPE=local-regression
fi

# TODO: change kkowalczyk to go-toml and create toml-fuzzer target there
./fuzzit create job --type $TYPE go-toml/${TARGET} ${TARGET}
