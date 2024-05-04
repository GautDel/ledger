#!/bin/bash

# Load environment variables from .env file
source .env

cd ./migrations

# Run goose status command
goose $GOOSE_DRIVER $DB_URL down
