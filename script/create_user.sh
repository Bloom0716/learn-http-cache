#!/bin/bash

curl -i -X POST localhost:8080/api/users/create \
    -H "Content-Type: application/json" \
    -d '{"name": "test user"}'
