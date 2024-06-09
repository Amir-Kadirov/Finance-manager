CURRENT_DIR=$(shell pwd)

# Paths
SWAGGER_CMD=swag
SWAGGER_OUTPUT_DIR=api/docs
SWAGGER_MAIN_FILE=api/router.go

# Targets
.PHONY: all gen-swagger clean-swagger install-swag

all: gen-swagger

gen-swagger: install-swag
	rm -rf ${SWAGGER_OUTPUT_DIR}
	mkdir -p ${SWAGGER_OUTPUT_DIR}
	${SWAGGER_CMD} init -g ${SWAGGER_MAIN_FILE} -o ${SWAGGER_OUTPUT_DIR}

install-swag:
	@if ! [ -x "$$(command -v ${SWAGGER_CMD})" ]; then \
		echo "Swag is not installed. Installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	export PATH=$(go env GOPATH)/bin:$$PATH

clean-swagger:
	rm -rf ${SWAGGER_OUTPUT_DIR}
