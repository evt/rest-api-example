#!/bin/bash
export GOOGLE_APPLICATION_CREDENTIALS=./serviceaccount.json

# Local postgres
export PG_URL=postgres://postgres:postgres@localhost/test?sslmode=disable
export PG_MIGRATIONS_PATH=file://../db//migrations
