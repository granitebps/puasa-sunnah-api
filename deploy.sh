#!/bin/sh
set -e

echo "Deploying application ..."

# Update codebase
git fetch origin main
git reset --hard origin/main

# Generate docs
./bin/swag-linux init

# Build the application
go build main.go

# Restart the service
sudo service psn restart

echo "Application deployed!"