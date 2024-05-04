#!/bin/bash

# Set the name for the executable
EXECUTABLE_NAME="ledger.systems"

# Build the Go application with the custom name
go build -o "$EXECUTABLE_NAME" ./cmd

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "Build successful. Running $EXECUTABLE_NAME..."
    
    # Run the executable
    ./"$EXECUTABLE_NAME"
else
    echo "Build failed."
fi

