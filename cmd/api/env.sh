#!/bin/bash
# google cloud storage
export GOOGLE_APPLICATION_CREDENTIALS=./serviceaccount.json
export GC_BUCKET=businessclass-stage-tracer-logs

# Local postgres
export PG_URL=postgres://postgres:postgres@localhost/test?sslmode=disable
export PG_MIGRATIONS_PATH=file://../../db//migrations

# Logger settings
export LOG_LEVEL=debug