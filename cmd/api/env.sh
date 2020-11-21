#!/bin/bash
export HTTP_ADDR=localhost:8080

### File upload settings

# google cloud storage (for file content)
#export GOOGLE_APPLICATION_CREDENTIALS=./serviceaccount.json
#export GC_BUCKET=my-cool-bucket

# path for local file storage
export FILE_PATH=../../files

### DB settings

# Postgres settings
#export PG_URL=postgres://postgres:postgres@localhost/test?sslmode=disable
#export PG_MIGRATIONS_PATH=file://../../store/pg/migrations

# MySQL settings
export MYSQL_ADDR=127.0.0.1:3306
export MYSQL_USER=api
export MYSQL_PASSWORD=api
export MYSQL_DB=apiserver
export MYSQL_MIGRATIONS_PATH=file://../../store/mysql/migrations

# Logger settings
export LOG_LEVEL=debug