#!/bin/sh

set -ex

pushd server_nocors
    go build -o server_nocors_bin server_nocors.go
	./server_nocors_bin&
popd

pushd server_cors
     go build -o server_cors_bin server_cors.go 
	./server_cors_bin&
popd 

open ui.html