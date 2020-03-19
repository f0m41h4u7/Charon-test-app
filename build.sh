#!/bin/bash
docker build --tag charon-registry:5000/test-app:v2.3 .
docker push charon-registry:5000/test-app:v2.3
