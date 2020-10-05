#!/bin/bash
export GOOGLE_APPLICATION_CREDENTIALS=./serviceaccount.json

# Google Cloud Postgres
export PG_PROTO=unix
export PG_ADDR=/cloudsql/helical-arcade-290815:europe-west1:wakeup/.s.PGSQL.5432
export PG_DB=wakeup
export PG_USER=wakeup
export PG_PASSWORD=wakeup
