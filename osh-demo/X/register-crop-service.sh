#!/bin/bash
set -e
. setup.sh

#tag::public[]
cropper=`echo $PZSERVER | sed -e sXpiazzaXhttp://pz-svcs-prevgenX`
cropper="$cropper/crop"

service='{
    "url": "'"$cropper"'",
    "method": "POST",
    "contractUrl": "http://example.com/contract",
    "resourceMetadata": {
        "name": "Preview Generator",
        "description": "Service that takes payload containing S3 location and bounding box for some raster file, downloads, crops and uploads the crop back up to s3.",
        "classType": "unclassified"
    }
}'

$curl -X POST -d "$service" $PZSERVER/service
#end::public[]
