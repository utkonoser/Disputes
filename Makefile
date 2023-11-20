APP_NAME=disputes
PROJECT_STRUCTURE_API=api
PROJECT_STRUCTURE_CLIENTS=pkg/api
PROJECT_STRUCTURE_IMPLEMENTATIONS=internal/app/api
FROM_SERVICES_TO_ROOT_REL=$(shell echo $(PROJECT_STRUCTURE_CLIENTS) | perl -F/ -lane 'print "../"x scalar(@F)')
IMPLEMENTATION_TYPE_NAME="Implementation"
SERVICES=$(shell ls -1 $(PROJECT_STRUCTURE_API) | grep \.proto | sed s/\.proto//)

# We always use go in module mode
export GO111MODULE=on
GO_EXEC=go

BIN?=./bin/$(APP_NAME)
BIN_MIGRATOR=./bin/migrator
LOCAL_BIN:=$(CURDIR)/bin
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")
PKGS=$(shell go list -f '{{.Dir}}' ./... | grep -v /vendor/)
GENMOCKS = $(shell (find internal -name 'genmock.go'; find test -name 'genmock.go') | sort -u)

PKGMAP:=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types

BUILD_ENVPARMS:=CGO_ENABLED=0

$(LOCAL_BIN):
	@mkdir -p $@
$(LOCAL_BIN)/%: ; $(info $(M) building $(REPOSITORY)…)
	$Q tmp=$$(mktemp -d); \
		(GOPATH=$$tmp GO111MODULE=off go get $(REPOSITORY) && cp $$tmp/bin/* $(LOCAL_BIN)/.) || ret=$$?; \
		rm -rf $$tmp ; exit $$ret

GOFMT = gofmt

GOLINT = $(LOCAL_BIN)/golint
$(LOCAL_BIN)/golint: REPOSITORY=golang.org/x/lint/golint

GOMOCKERY = $(LOCAL_BIN)/mockery
$(LOCAL_BIN)/mockery: REPOSITORY=github.com/vektra/mockery/.../

GOSWG = $(LOCAL_BIN)/swagger
$(LOCAL_BIN)/swagger: REPOSITORY=github.com/go-swagger/go-swagger/cmd/swagger

# Detecting OS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    OS = LINUX
    REPLACE = sed -i"" -e
endif
ifeq ($(UNAME_S),Darwin)
    OS = OSX
    REPLACE = sed -i "" -e
endif


.PHONY: bin-deps
bin-deps: ; $(info $(M) install bin depends…) @ ## Install bin depends
	$(info #Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) $(GO_EXEC) get github.com/gogo/protobuf/protoc-gen-gofast
	GOBIN=$(LOCAL_BIN) $(GO_EXEC) get github.com/utrack/clay/v2/cmd/protoc-gen-goclay

.PHONY: deps
deps: ; $(info $(M) install depends…) @ ## Install depends
	$(info #Install dependencies...)
	$(GO_EXEC) mod tidy

.PHONY: test
test: deps lint ; $(info $(M) run tests…) @ ## Run tests
	$(info #Running tests...)
	$(GO_EXEC) test ./...

.PHONY: fast-test
fast-test: deps ; $(info $(M) run fast tests…) @ ## Run fast tests for deploy
	$(info #Running fast tests...)
	$(GO_EXEC) test --tags=ci ./...

.PHONY: fast-build
fast-build: deps ; $(info $(M) run fast build…) @ ## Run fast build for deploy
	$(info #Building app...)
	$(BUILD_ENVPARMS) $(GO_EXEC) build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/$(APP_NAME)/main.go

.PHONY: build
build: fast-test fast-build

.PHONY: build-migrator
build-migrator: ; $(info $(M) build migrator…) @ ## Build migrator
	$(info #Building migrator...)
	$(BUILD_ENVPARMS) $(GO_EXEC) build -ldflags "$(LDFLAGS)" -o $(BIN_MIGRATOR) ./cmd/migrator/main.go

.PHONY: generate-mocks
generate-mocks: $(LOCAL_BIN) | $(GOMOCKERY) ; $(info $(M) go generate-mocks…) @ ## Generation Mocks
	$(Q) for f in $(GENMOCKS); do \
		PATH=$(LOCAL_BIN):"$(PATH)" $(GO_EXEC) generate $$f;\
	done

.PHONY: lint
lint: $(LOCAL_BIN) | $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$(Q) ERR=$$($(GOLINT) $(PKGS) | grep -v '\(should have comment\|comment on exported\)'); \
	    if [ -n "$$ERR" ]; then \
	    	echo "$@ errors:\n$$ERR" && \
	        exit 1; \
	    fi && \
	    echo "$@: done"

.PHONY: generate
generate: bin-deps ; $(info $(M) go generate…) @ ## Generation Services
	$(Q) for srv in $(SERVICES); do \
	    echo "Generate $(CURDIR)/$(PROJECT_STRUCTURE_CLIENTS)/$$srv" && \
	    echo "Implementation $(FROM_SERVICES_TO_ROOT_REL)../../$(PROJECT_STRUCTURE_IMPLEMENTATIONS)/$$srv" && \
		mkdir -p $(CURDIR)/$(PROJECT_STRUCTURE_CLIENTS)/$$srv && \
		cd $(CURDIR)/$(PROJECT_STRUCTURE_CLIENTS)/$$srv && \
		protoc --plugin=protoc-gen-goclay=$(LOCAL_BIN)/protoc-gen-goclay \
			--plugin=protoc-gen-gofast=$(LOCAL_BIN)/protoc-gen-gofast \
			-I$(FROM_SERVICES_TO_ROOT_REL)../api/:$(CURDIR)/vendor.pb \
			--gofast_out=$(PKGMAP),plugins=grpc:. \
			--goclay_out=$(PKGMAP),impl=true,impl_service_sub_dir=false,impl_path=$(FROM_SERVICES_TO_ROOT_REL)../$(PROJECT_STRUCTURE_IMPLEMENTATIONS)/$$srv,impl_type_name_tmpl=$(IMPLEMENTATION_TYPE_NAME):. \
			$(FROM_SERVICES_TO_ROOT_REL)../$(PROJECT_STRUCTURE_API)/$$srv.proto && \
		$(REPLACE) "s/,omitempty//g" ./*.pb.go > /dev/null && \
		if grep -q RegisterEnum ./$$srv.pb.go; then \
           ( \
               echo "package $$srv"; \
               echo; \
               echo 'import "github.com/gogo/protobuf/proto"'; \
               echo; \
               echo 'func init() {'; \
               grep RegisterEnum ./$$srv.pb.go | sed s/_name,\ /_name,\ types.ExtendEnumWithLowerKeys\(/ | sed s/\)$$/\)\)/ ; \
               echo '}' \
           ) >./$$srv.pb.enum.go; \
           grep RegisterEnum $$srv.pb.go | \
               awk '{printf "\n// MarshalJSON is a custom marshaller function\nfunc (x "} {printf substr($$3, 0, match($$3, "_value")-1)} {printf ") MarshalJSON() ([]byte, error) { return []byte(`\"` + x.String() + `\"`), nil }\n"}' >> ./$$srv.pb.enum.go; \
        else \
         rm -f ./$$srv.pb.enum.go; \
        fi; \
	done

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ": .*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'