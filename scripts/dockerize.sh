#!/bin/sh

readonly image_tag="$1"
readonly image_repository="$2"
readonly type="$3"
readonly service="$4"
readonly push_image="$5"

buildDir="$type"
if [ "$type" = "app" ]; then
  buildDir="internal"
fi


docker build \
  --no-cache \
  --progress=plain \
  --tag "$image_repository":"$image_tag" \
  --file "./docker/$type/Dockerfile" \
  --build-arg "SERVICE=$service" \
  "./$buildDir"

if [ "$push_image" == "true" ]; then
  docker push "$image_repository":"$image_tag"
fi
