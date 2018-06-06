#!/bin/bash

# load the config
set -o allexport
source ./run.env
set +o allexport

# Run the kafka producer or consumer
./kafka consume