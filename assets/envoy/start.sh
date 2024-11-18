#!/bin/bash 
cd "$(dirname "$0")"

envoy -c ./config.yml &> ./envoy.log  &