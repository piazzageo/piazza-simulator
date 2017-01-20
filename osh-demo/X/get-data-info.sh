#!/bin/bash
set -e
. setup.sh

check_arg $1 jobId

# tag::public[]
dataId=$1

$curl -X GET $PZSERVER/data/$dataId
# end::public[]
