#!/bin/bash

usage() {
    echo "Usage: $0 [major|minor|patch]"
    echo "Example: $0 major"
    exit 1
}

if [ "$#" -ne 1 ]; then
    usage
fi

SERVICE_NAME="message-service"
PART=$1

VERSION_FILE=".version"

cd ../

if [ ! -f "$VERSION_FILE" ]; then
    echo "0.0.0" > "$VERSION_FILE"
fi

current_version=$(cat "$VERSION_FILE")
echo "Current version for $SERVICE_NAME: $current_version"

bump_version() {
    local version=$1
    local part=$2

    IFS='.' read -r -a version_parts <<< "$version"

    case $part in
        major)
            version_parts[0]=$((${version_parts[0]}+1))
            version_parts[1]=0
            version_parts[2]=0
            ;;
        minor)
            version_parts[1]=$((${version_parts[1]}+1))
            version_parts[2]=0
            ;;
        patch)
            version_parts[2]=$((${version_parts[2]}+1))
            ;;
        *)
            echo "Error: Invalid part to bump ('$part'). Use 'major', 'minor', or 'patch'."
            exit 1
            ;;
    esac

    new_version="${version_parts[0]}.${version_parts[1]}.${version_parts[2]}"
    echo "$new_version" > "$VERSION_FILE"
    echo "Version bumped to $new_version"
}

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 [major|minor|patch]"
    exit 1
fi

bump_version "$current_version" "$1"
