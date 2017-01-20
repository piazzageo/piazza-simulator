#!/bin/bash
set -e

export PZSERVER=piazza.geointservices.io

# API key
if [ -f ~/.pzkey ] ; then
    export PZKEY=`cat ~/.pzkey | grep $PZSERVER | cut -f 2 -d ":" | cut -d \" -f 2`
else
    : "${PZKEY?"PZKEY is not set"}"
fi

# curl boilerplate
curl="curl -S -s -u $PZKEY:"" -H Content-Type:application/json"
curl_multipart="curl -S -s -u $PZKEY:"" -H Content-Type:multipart/form-data;boundary=thisismyboundary"


# cmd line arg helper
check_arg() {
    arg=$1
    name=$2
    if [ "$#" -ne "2" ];
    then
        echo "Internal error: check_arg has wrong numbner of args ($#): $@"
        exit 1
    fi
    if [ "$arg" == "" ]
    then
        echo "usage error: \$name missing"
        exit 1
    fi
}
