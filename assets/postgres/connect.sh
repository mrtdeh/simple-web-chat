#!/bin/bash

docker exec -it postgres_container psql -U apiservices -d test
