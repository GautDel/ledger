#!/bin/bash

# Load environment variables from .env file
source .env

# Run goose status command
goose $GOOSE_DRIVER $DB_URL status
