#!/bin/bash
set -e
. setup.sh

check_arg $1 dataId

# tag::public[]
dataId=$1

data='{
    "type": "access",
    "dataId": "'"$dataId"'",
    "deploymentType": "geoserver"
}'

$curl -X POST -d "$data" $PZSERVER/deployment
# end::public[]
