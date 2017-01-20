#!/bin/bash
set -e

# This script will get your API key.
#
# $PZSERVER must be set; see setup.sh for details.
#
# usage:
#   ./getkeys.sh USERNAME "PASSWORD"
#
# Your new API key will be printed.
#
# You may wish to put the key into your ~/.pzkey file. The file contains a
# JSON map from Piazza server names to API key strings, e.g.
#
#     {
#         "pz-gateway.int.venicegeo.io" : "aa91dafd-c4ff-43b1-8f73-d5e363b9381b",
#         "piazza.example.com": "ba91fe4d-c1ff-43b1-8f76-d5e339a9381c"
#     }

if [[ $# -ne 2 ]]; then
    echo Usage: $ sh getkeys.sh USERNAME "PASSWORD"
    exit 1
fi
user=$1
pass=$2

: "${PZSERVER:?"PZSERVER is not set"}"

json=`curl -S -s -XGET -u $user:$pass https://$PZSERVER/key`
#echo "$json"

key=`echo "$json" | grep '"uuid" :' | cut -f 2 -d ":" | cut -d \" -f 2`
echo $key
