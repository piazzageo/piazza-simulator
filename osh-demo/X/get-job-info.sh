#!/bin/bash
set -e
. setup.sh

check_arg $1 jobId

# tag::public[]
jobId=$1

$curl -XGET $PZSERVER/job/$jobId
# end::public[]
