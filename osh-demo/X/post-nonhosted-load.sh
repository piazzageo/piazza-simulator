#!/bin/bash
set -e
. setup.sh

check_arg $1 name

# tag::public[]
name=$1

data='{
    "type": "ingest",
    "host": false,
    "data": {
        "dataType": {
            "type": "raster",
            "location": {
                "type": "s3",
                "bucketName": "external-public-access-test",
                "fileName": "elevation.tif",
                "domainName": "s3.amazonaws.com"
            }
        },
        "metadata": {
            "name": "'"$name"'",
            "description": "mydescription"
        }
    }
}'

$curl -XPOST -d "$data" $PZSERVER/data
# end::public[]
