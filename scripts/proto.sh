#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"

readonly in_proto_dir="api/proto"
readonly in_proto_files="$in_proto_dir/$service"

echo "$in_proto_files"

mkdir -p "$output_dir"

function generate_proto() {
  protoc \
    "--proto_path=$in_proto_dir" \
    "--go_out=$output_dir" \
    "--go-grpc_out=$output_dir" \
    "--go-grpc_opt=require_unimplemented_servers=false" \
    "--grpc-gateway_out=$output_dir" \
    "$1"
}

# shellcheck disable=SC2044
for proto_file in $(find "$in_proto_files" -name "*.proto"); do
  echo "Generate sources for ${proto_file}"
  protoc \
    "--proto_path=$in_proto_dir" \
    "--go_out=$output_dir" \
    "--go-grpc_out=$output_dir" \
    "--go-grpc_opt=require_unimplemented_servers=false" \
    "--grpc-gateway_out=$output_dir" \
    "$1"
done
