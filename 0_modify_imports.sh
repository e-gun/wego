#!/usr/bin/env sh
perl -pi -w -e 's/github.com\/ynqa\/wego/github.com\/e-gun\/wego/g;' *.go
perl -pi -w -e 's/github.com\/ynqa\/wego/github.com\/e-gun\/wego/g;' */*.go
perl -pi -w -e 's/github.com\/ynqa\/wego/github.com\/e-gun\/wego/g;' */*/*.go
perl -pi -w -e 's/github.com\/ynqa\/wego/github.com\/e-gun\/wego/g;' */*/*/*.go