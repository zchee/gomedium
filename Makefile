GO_SRCS = $(shell find . -type f \( -name '*.go' -and -not -iwholename '*vendor*' -and -not -iwholename '*testdata*' \) )
PACKAGE_NAME = $(notdir $(CURDIR))

all: build

.PHONY: build
build: bin/$(PACKAGE_NAME)

bin/$(PACKAGE_NAME): $(GO_SRCS)
	go build -v -x -o ./bin/$(PACKAGE_NAME)

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: vendor/clean
vendor/install:
	go install -v -x $(go list ./vendor/...)

.PHONY: vendor/clean
vendor/clean:
	find ./vendor/github.com/google/flatbuffers -maxdepth 1 -type d -not -name 'flatbuffers' -and -not -name 'go' -exec rm -fr {} ";" || true  # for flatbuffers
	find ./vendor -type f -name '*_test.go' -delete  # test files
	find ./vendor -type d -name 'testdata' -exec rm -fr {} ";" || true
	find ./vendor -type f \( -not -name '*.go' -not -name 'README*' -and -not -name 'readme.*' -and -not -name 'LICENSE*' \) -delete  # delete unused files expect *.go, README and License

.PHONY: debug
debug:
	@echo $(PACKAGE_NAME)
	@echo $(CMD_NAME)
