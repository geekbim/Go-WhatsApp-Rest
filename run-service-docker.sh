#!/bin/sh
docker build -n go_wa_rest -t go_wa_rest -f deploy/Dockerfile .
docker run -d go_wa_rest