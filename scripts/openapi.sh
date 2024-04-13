#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"
readonly role="$4"

readonly in_oapi_dir="api/openapi"
readonly in_oapi_file="$in_oapi_dir/$service/$service.yaml"

readonly out_oapi_types_file="openapi_types.gen.go"
readonly out_oapi_api_file="openapi_api.gen.go"
readonly out_oapi_client_file="openapi_client.gen.go"

mkdir -p "$output_dir"

echo "Generate types for" "$package" "$in_oapi_file""..."
oapi-codegen --old-config-style -generate types -o "$output_dir/$out_oapi_types_file" -package "$package" "$in_oapi_file"

if [[ $role  == "client" ]]
then
    echo "Generate client for" "$package" "$in_oapi_file""..."
    oapi-codegen --old-config-style -generate client -o "$output_dir/$out_oapi_client_file" -package "$package" "$in_oapi_file"
elif [[ $role == "server" ]]
then
    echo "Generate chi-server for " "$package" "$in_oapi_file""..."
    oapi-codegen --old-config-style -generate chi-server -o "$output_dir/$out_oapi_api_file" -package "$package" "$in_oapi_file"
elif [[ $role == "both" ]]
then
    echo "Generate both client and chi-server for" "$package" "$in_oapi_file""..."
    oapi-codegen --old-config-style -generate chi-server -o "$output_dir/$out_oapi_api_file" -package "$package" "$in_oapi_file"
    oapi-codegen --old-config-style -generate client -o "$output_dir/$out_oapi_client_file" -package "$package" "$in_oapi_file"
fi
echo "OK"