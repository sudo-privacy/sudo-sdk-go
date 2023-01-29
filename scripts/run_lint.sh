#!/bin/bash

set -e
golangci-lint run --timeout=3m
