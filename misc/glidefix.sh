#!/bin/sh
set -e

glide cache-clear
for i in pz-gocommon pz-logger pz-uuidgen pz-workflow
do
    cd $i
    rm -fr glide.* vendor
    glide init --non-interactive
    glide install --strip-vendor
    cd ..
done

