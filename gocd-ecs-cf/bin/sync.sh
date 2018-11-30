#!/bin/bash

APP_VER=$GO_PIPELINE_LABEL
AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
AWS_DEFAULT_REGION=$AWS_DEFAULT_REGION

echo "APP_VER -> $APP_VER"
echo "BUCKET_NAME -> $BUCKET_NAME"

echo "Uploading CloudFormation template & scripts to S3 -> $BUCKET_NAME ..."
aws s3 cp `pwd`/.. s3://$BUCKET_NAME/ --recursive
echo "Done."
