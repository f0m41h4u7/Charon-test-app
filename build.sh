#!/bin/bash
docker build --tag charon-registry:5000/test-app:v5 .
docker push charon-registry:5000/test-app:v5
