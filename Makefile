GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go" -type f -not -path "./api/*"  -not -path "./vendor/*")
VETPACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/ | grep -v /examples/)

.PHONY: build
build: fmt-check
	@script/make.sh build  $(project)

.PHONY: all
all: fmt-check
	@script/make.sh build all  $(project)

.PHONY: generate
generate:
	go generate -skip=wire "./..." && wire "./..."

.PHONY: test
test:
	$(GO) test -short -failfast -coverprofile=cover.out ./...

.PHONY: cover
cover:
	$(GO) tool cover -func=cover.out -o cover_total.out
	$(GO) tool cover -html=cover.out -o cover.html

.PHONY: fmt
fmt:
	@$(GOFMT) -w $(GOFILES) && echo "\033[34m[Code] format perfect!\033[0m";

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi; \
	echo "\033[34m[Code] format perfect!\033[0m";

vet:
	$(GO) vet $(VETPACKAGES)

.PHONY: deploy
deploy:
	@script/make.sh deploy $(project)

.PHONY: deployall
deployall:
	@script/make.sh deploy all $(project)

.PHONY: gitpull
gitpull:
	@echo "Checking remote changes..."
	$(eval BRANCH := $(shell git branch -r | grep -w origin | sed 's|/|-|g'))
	@if [ -n "$(BRANCH)" ]; then \
		echo "Pulling changes..."; \
		git pull; \
	fi

.PHONY: gitadd
gitadd:
	git add .

.PHONY: protolock
protolock:
	@echo "protofile check"
	@[ ! -f proto.lock ] && protolock init || protolock status && protolock commit

# proto generate
.PHONY: proto
proto: gitpull protolock genproto   gitadd

protorm:
	@rm -rf proto.lock && protolock init

# proto regenerate
.PHONY: protoinit
protoinit: protorm proto

.PHONY: protoclean
protoclean:
	@for file in `find api -name '*.pb*.go'`; do \
	  rm $$file; \
	done;

.PHONY: genproto
genproto:
	@script/make.sh proto $(api)

.PHONY: mock
mock:
	@script/make.sh mock $(api)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: swagger
swagger:
	@script/make.sh swagger $(api)

.PHONY:genrsa
genrsa:
	# openssl 3.0 默认生成的是pkcs8格式，如果需要默认生成pkcs1则需要添加-traditional参数
	openssl genrsa -out rsa_private_key.pem  2048
	openssl rsa -in rsa_private_key.pem -pubout -out public_key.pem

