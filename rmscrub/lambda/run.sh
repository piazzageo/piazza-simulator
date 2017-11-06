#!/bin/sh

make

aws lambda create-function \
  --function-name RunTagger \
  --runtime python2.7 \
  --role arn:aws:iam::943964223958:role/mpg-serverless-stack-TestFunctionRole-I0YNTX7VFMN3 \
  --handler handler.RunTagger \
  --zip-file fileb://handler.zip

aws lambda update-function-code \
  --function-name RunTagger \
  --zip-file fileb://handler.zip

aws lambda invoke \
  --function-name RunScrubber \
  --invocation-type RequestResponse \
  --payload file://input.txt \
  x
