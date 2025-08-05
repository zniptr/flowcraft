#!/bin/sh
go test -coverprofile=coverage.out -coverpkg= \
github.com/zniptr/flowcraft/pkg/chartmanager \
github.com/zniptr/flowcraft/pkg/chartcontext \
github.com/zniptr/flowcraft/pkg/executableregistry \
github.com/zniptr/flowcraft/internal/filereader \
github.com/zniptr/flowcraft/internal/xmlparser \
github.com/zniptr/flowcraft/internal/chartinstance \
github.com/zniptr/flowcraft/internal/actions \
github.com/zniptr/flowcraft/internal/chart \
&& go tool cover -html=coverage.out -o coverage.html