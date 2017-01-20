#!/bin/sh
set -e

rm -fr glide.* vendor vendors
glide init --non-interactive
glide install --strip-vendor

