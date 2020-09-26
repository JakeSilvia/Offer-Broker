#!/usr/bin/env bash

gcloud auth activate-service-account --key-file=organic-bivouac-290421-56b1ec812736.json
gcloud config set project organic-bivouac-290421

npm build
npm run prod
gcloud app deploy app.yaml --promote --quiet
