# This workflow will build a golang project
# For more information see: https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go

name: redirection_service_build

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:

  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Build
      run: go build -v ./src/redirection_service/...

    - name: Test
      run: go test -v --coverprofile redirection_service_cover.out ./src/redirection_service/...
    
    - name: Coverage
      run: go tool cover --func=redirection_service_cover.out
