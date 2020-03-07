#!/bin/bash
docker build --tag charon-registry:5000/test-app:v1.1 .
docker push charon-registry:5000/test-app:v1.1
