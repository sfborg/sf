# sf project justfile

# Variables
APP := "sf"
ORG := "github.com/sfborg/"
BUILD_DIR := "out/" 
RELEASE_DIR := BUILD_DIR + "releases/" 
TEST_OPTS := "-count=1 -p 1 -shuffle=on -coverprofile=coverage.txt -covermode=atomic"

NO_C := "CGO_ENABLED=0"
X86 := "GOARCH=amd64"
ARM := "GOARCH=arm64"
LINUX := "GOOS=linux"
MAC := "GOOS=darwin"
WIN := "GOOS=windows"

# Colors
GREEN := `tput -Txterm setaf 2`
YELLOW := `tput -Txterm setaf 3`
WHITE := `tput -Txterm setaf 7`
CYAN := `tput -Txterm setaf 6`
RESET := `tput -Txterm sgr0`

# Get version from git
VERSION := `git describe --tags`
VER := `git describe --tags --abbrev=0`
DATE := `date -u '+%Y-%m-%d_%H:%M:%S%Z'`

# LD flags with version and build date
FLAGS_LD := "-trimpath -ldflags '-X " + ORG + APP + \
    "/pkg/sf.Build=" + DATE + " -X " + ORG + APP + \
    "/pkg/sf.Version=" + VERSION + "'"
FLAGS_REL := "-trimpath -ldflags '-s -w -X " + ORG + APP + \
    "/pkg/sf.Build=" + DATE + "'"

# Default recipe (just install)
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
    @echo {{VERSION}}

# Clean up and sync dependencies
tidy:
    @go mod tidy
    @go mod verify

# Install tools
tools: tidy
    @go install tool
    @echo "✅ tools of the project are installed"


# Build binary
build:
    {{NO_C}} go build -o  {{BUILD_DIR}}{{APP}} {{FLAGS_LD}} 
    @echo "✅ {{APP}} built to {{BUILD_DIR}}{{APP}}"

# Build binary without debug info and with hardcoded version
buildrel:
    {{NO_C}} go build -o {{BUILD_DIR}}{{APP}} {{FLAGS_REL}} 
    @echo "✅ {{APP}} release binary built to {{BUILD_DIR}}{{APP}}"

# Build and install binary
install:
    {{NO_C}} go install {{FLAGS_LD}}
    @echo "✅ {{APP}} installed to ~/go/bin/{{APP}}"

# Build and package binaries for a release
release: buildrel
    @echo "Building releases for Linux, Mac, Windows (Intel and Arm)"
    @mkdir -p {{RELEASE_DIR}}

    {{NO_C}} {{LINUX}} {{X86}} go build {{FLAGS_REL}} -o {{RELEASE_DIR}}{{APP}} 
    tar zcvf {{RELEASE_DIR}}{{APP}}-{{VER}}-linux-amd64.tar.gz {{RELEASE_DIR}}{{APP}}
    rm {{RELEASE_DIR}}{{APP}}

    {{NO_C}} {{LINUX}} {{ARM}} go build {{FLAGS_REL}} -o {{RELEASE_DIR}}{{APP}} 
    tar zcvf {{RELEASE_DIR}}{{APP}}-{{VER}}-linux-arm64.tar.gz {{RELEASE_DIR}}{{APP}}
    rm {{RELEASE_DIR}}{{APP}}

    {{NO_C}} {{MAC}} {{X86}} go build {{FLAGS_REL}} -o {{RELEASE_DIR}}{{APP}}
    tar zcvf {{RELEASE_DIR}}{{APP}}-{{VER}}-mac-amd64.tar.gz {{RELEASE_DIR}}{{APP}}
    rm {{RELEASE_DIR}}{{APP}}

    {{NO_C}} {{MAC}} {{ARM}} go build {{FLAGS_REL}} -o {{RELEASE_DIR}}{{APP}}
    tar zcvf {{RELEASE_DIR}}{{APP}}-{{VER}}-mac-arm64.tar.gz {{RELEASE_DIR}}{{APP}}
    rm {{RELEASE_DIR}}{{APP}}

    {{NO_C}} {{WIN}} {{X86}} go build {{FLAGS_REL}} -o {{RELEASE_DIR}}{{APP}}.exe 
    cd {{RELEASE_DIR}} && zip -9 {{APP}}-{{VER}}-win-amd64.zip  {{APP}}.exe
    rm {{RELEASE_DIR}}{{APP}}.exe

    {{NO_C}} {{WIN}} {{ARM}} go build {{FLAGS_REL}} -o {{RELEASE_DIR}}{{APP}}.exe 
    cd {{RELEASE_DIR}} && zip -9 {{APP}}-{{VER}}-win-arm64.zip {{APP}}.exe
    rm {{RELEASE_DIR}}{{APP}}.exe

    @echo "✅ Release binaries created in {{RELEASE_DIR}}" 
# Clean all the files and binaries generated
clean:
    @rm -rf ./{{BUILD_DIR}}

# Run the tests of the project
test:
    go test {{TEST_OPTS}} ./...

# Run the tests of the project and export the coverage
coverage:
    @go test -p 1 -cover -covermode=count -coverprofile=profile.cov ./...
    @go tool cover -func profile.cov
