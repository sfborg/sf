# sf project justfile

# Variables
PROJ_NAME := "sf"
RELEASE_DIR := "/tmp"
TEST_OPTS := "-count=1 -p 1 -shuffle=on -coverprofile=coverage.txt -covermode=atomic"

NO_C := "CGO_ENABLED=0"
FLAGS_SHARED := NO_C + " GOARCH=amd64"
FLAGS_LINUX := FLAGS_SHARED + " GOOS=linux"
FLAGS_MAC := FLAGS_SHARED + " GOOS=darwin"
FLAGS_MAC_ARM := "GOARCH=arm64 GOOS=darwin"
FLAGS_WIN := FLAGS_SHARED + " GOOS=windows"

GOCMD := "go"
GOTEST := GOCMD + " test"
GOVET := GOCMD + " vet"
GOCLEAN := GOCMD + " clean"

# Colors
GREEN := `tput -Txterm setaf 2`
YELLOW := `tput -Txterm setaf 3`
WHITE := `tput -Txterm setaf 7`
CYAN := `tput -Txterm setaf 6`
RESET := `tput -Txterm sgr0`

# Get version from git
version := `git describe --tags`
ver := `git describe --tags --abbrev=0`
date := `date -u '+%Y-%m-%d_%H:%M:%S%Z'`

# LD flags with version and build date
flags_ld := "-ldflags \"-X github.com/sfborg/" + PROJ_NAME + "/pkg.Build=" + date + " -X github.com/sfborg/" + PROJ_NAME + "/pkg.Version=" + version + "\""
flags_rel := "-trimpath -ldflags \"-s -w -X github.com/sfborg/" + PROJ_NAME + "/pkg.Build=" + date + "\""

# Default recipe (runs when just is called without arguments)
default: install

# Show this help
help:
    @echo ''
    @echo 'Usage:'
    @echo '  {{YELLOW}}just{{RESET}} {{GREEN}}<target>{{RESET}}'
    @echo ''
    @echo 'Targets:'
    @just --list --unsorted

# Display current version
version:
    @echo {{version}}

# Download dependencies
deps:
    {{GOCMD}} mod download

# Install tools
tools: deps
    @cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

# Build binary
build:
    {{NO_C}} {{GOCMD}} build \
        -o {{PROJ_NAME}} \
        {{flags_ld}} \
        .

# Build binary without debug info and with hardcoded version
buildrel:
    {{NO_C}} {{GOCMD}} build \
        -o {{PROJ_NAME}} \
        {{flags_rel}} \
        .

# Build and install binary
install:
    {{NO_C}} {{GOCMD}} install {{flags_ld}}

# Build and package binaries for a release
release: buildrel
    {{GOCLEAN}}
    {{FLAGS_LINUX}} {{GOCMD}} build {{flags_rel}}
    tar zcvf {{RELEASE_DIR}}/{{PROJ_NAME}}-{{ver}}-linux.tar.gz {{PROJ_NAME}}
    {{GOCLEAN}}
    {{FLAGS_MAC}} {{GOCMD}} build {{flags_rel}}
    tar zcvf {{RELEASE_DIR}}/{{PROJ_NAME}}-{{ver}}-mac.tar.gz {{PROJ_NAME}}
    {{GOCLEAN}}
    {{FLAGS_MAC_ARM}} {{GOCMD}} build {{flags_rel}}
    tar zcvf {{RELEASE_DIR}}/{{PROJ_NAME}}-{{ver}}-mac-arm.tar.gz {{PROJ_NAME}}
    {{GOCLEAN}}
    {{FLAGS_WIN}} {{GOCMD}} build {{flags_rel}}
    zip -9 {{RELEASE_DIR}}/{{PROJ_NAME}}-{{ver}}-win-64.zip {{PROJ_NAME}}.exe
    {{GOCLEAN}}

# Clean all the files and binaries generated
clean:
    rm -rf ./out

# Run the tests of the project
test:
    {{GOTEST}} {{TEST_OPTS}} ./...

# Run the tests of the project and export the coverage
coverage:
    {{GOTEST}} -cover -covermode=count -coverprofile=profile.cov ./...
    {{GOCMD}} tool cover -func profile.cov
