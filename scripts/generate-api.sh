#!/bin/bash

echo "Generating UAPI from OpenAPI spec"

# Check if argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <openapi-spec-file>"
    exit 1
fi

# Check if openapi-generator is installed
if ! command -v openapi-generator &> /dev/null
then
    echo "openapi-generator could not be found"
    echo "Please install openapi-generator-cli"
    echo "https://openapi-generator.tech/docs/installation"
    exit
fi

# Clearing old generated files
echo "clearing old generated files"
find ./api -name "*.go" -type f -delete

# Generate UAPI
openapi-generator generate \
    -i $1 \
    -g go \
    -o ./api \
    --package-name api \
    --git-user-id "yohan460" \
    --git-repo-id "go-jamf-api" \
    --global-property=apiTests=false \
    --type-mappings=integer=int64 \
    --additional-properties=packageName=api,enumClassPrefix=true,structPrefix=true,generateInterfaces=true,isGoSubmodule=true \
    --skip-validate-spec
returncode=$?
echo "Return code: $returncode"
if [[ "${returncode}" != "0" ]]; then
    echo "Failed to generate UAPI"
    exit 1
fi

# Perform cleanup
rm ./api/git_push.sh
rm ./api/go.mod
rm ./api/go.sum
rm ./api/.gitignore

# Patch the defaultValue generation
echo "Patching defaultValue generation"
find ./api -type f -name "*.go" -exec sed -i '' 's/var defaultValue \[\]\(.*\) = \[\(.*\)\]/defaultValue := []\1{\2}/g' {} \;
returncode=$?
if [[ "${returncode}" != "0" ]]; then
    echo "Failed to patch defaultValue generation"
    exit 1
fi

# Patch the generic object generation
echo "Patching generic object generation"
find ./api -type f -name "*.go" -exec sed -i '' 's/MapmapOfStringinterface[\{][\}]/Generic/g' {} \;;
returncode=$?
if [[ "${returncode}" != "0" ]]; then
    echo "Failed to patch generic object generation"
    exit 1
fi

# Patch reserved go file name suffixes
echo "Patching reserved go file name suffixes"
for file in api/*ios.go; do
    mv "$file" "${file%ios.go}ios_.go"
done
