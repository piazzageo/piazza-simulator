#!/bin/bash

set -e

. setup.sh

$curl -XGET $PZSERVER
