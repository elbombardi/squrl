#!/bin/bash

current_dir="$(pwd)"
dir=$(mktemp -d) 
git clone https://github.com/go-swagger/go-swagger $dir 
cd "$dir"	
go install "$dir"/cmd/swagger
cd "$current_dir"   