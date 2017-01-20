#!/bin/bash
set -e
. setup.sh

check_arg $1 serviceId

#tag::public[]
serviceId=$1

bucket="external-public-access-test"
file="NASA-GDEM-10km-colorized.tif"

job='{
    "type": "execute-service",
    "data": {
        "serviceId": "'"$serviceId"'",
        "dataInputs": {
            "test": {
                "content": "{ \"source\": { \"domain\": \"s3.amazonaws.com\", \"bucketName\": \"'"$bucket"'\", \"fileName\": \"'"$file"'\" }, \"function\": \"crop\", \"bounds\": { \"minx\": -140.00, \"miny\": 10.00, \"maxx\": -60.00, \"maxy\": 70.00 } }",
                "type": "body",
                "mimeType": "application/json"
            }
        },
        "dataOutput": [ {
            "mimeType":"image/tiff",
            "type":"raster"
        } ]
    }
}'

$curl -X POST -d "$job" $PZSERVER/job
#end::public[]
