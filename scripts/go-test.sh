#!/bin/sh
go test -coverprofile=coverage.out -coverpkg= \
github.com/zniptr/flowcraft/pkg/chartmanager \
github.com/zniptr/flowcraft/internal/filereader \
github.com/zniptr/flowcraft/internal/xmlparser \
&& go tool cover -html=coverage.out -o coverage.html