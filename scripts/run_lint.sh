#!/bin/bash

set -e -x
golangci-lint run --timeout=3m
exit $?
