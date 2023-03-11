#!/bin/sh
docker build -t go_wa_rest -f deploy/Dockerfile .
docker run -d go_wa_rest