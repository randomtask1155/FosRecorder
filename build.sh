#!/bin/bash

export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
go install FosRecorder
