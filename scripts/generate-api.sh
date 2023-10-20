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

# Generate UAPI
cd api
openapi-generator generate \
    -i $1 \
    -g go \
    --package-name api \
    --git-user-id "yohan460" \
    --git-repo-id "go-jamf-api" \
    --additional-properties=packageName=uapi,enumClassPrefix=true,generateInterfaces=true,structPrefix=true

returncode=$?
echo "Return code: $returncode"
if [[ "${returncode}" != "0" ]]; then
    echo "Failed to generate UAPI"
    exit 1
fi

# Patch the defaultValue generation
# find . -type f -print | xargs -0 sed -i 's/var defaultValue \[\](.*) = \[(.*)\]/defaultValue := []$1{$2}/g' {} ';'
# returncode=$?
# if [[ "${returncode}" != "0" ]]; then
#     echo "Failed to patch defaultValue generation"
#     exit 1
# fi