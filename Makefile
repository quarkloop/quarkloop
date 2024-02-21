# GOBIN=${GOBIN:=$(go env GOBIN)}

# if [ -z "$GOBIN" ]; then
# 	GOBIN="$(go env GOPATH)/bin"
# fi

proto:
	@if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	./tools/proto/generate.bash