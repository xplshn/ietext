#!/bin/sh

# Function to build each command in the cmds directory
build_commands() {
    for dir in ./cmds/*; do
        if [ -d "$dir" ]; then
            echo "Building $dir"
            (cd "$dir" && go build)
        fi
    done
}

# Function to copy built executables to the built directory
copy_executables() {
    mkdir -p ./built
    find ./cmds -type f -executable | while read -r executable; do
        cp "$executable" ./built/
    done
}

# Function to clean up the built directory and remove binaries in cmds
clean_up() {
    case "$1" in
        "full")
            rm -rf ./built
            find ./cmds -type f -executable -exec rm {} +
            ;;
        *)
            echo "Invalid argument. Use 'full' to clean everything."
            ;;
    esac
}

# Function to update dependencies
update_deps() {
    for dir in ./cmds/*; do
        if [ -d "$dir" ]; then
            echo "Updating deps for: $dir"
            (cd "$dir" && go get -u ; go mod tidy)
        fi
    done
}

# Main script execution
case "$1" in
        ""|"build")
        build_commands
        copy_executables
        ;;
    "clean")
        clean_up "$2"
        ;;
    "deps")
        update_deps
        ;;
        *)
        echo "Usage: $0 {--help} {build|clean|deps}"
        exit 1
        ;;
esac
