#!/bin/sh
docker build --name go_wa_rest -t go_wa_rest -f deploy/Dockerfile .
docker run -d go_wa_rest