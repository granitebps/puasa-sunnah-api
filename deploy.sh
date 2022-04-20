#!/bin/sh
set -e

echo "Deploying application ..."

# Update codebase
git fetch origin main
git reset --hard origin/main

# Generate docs
make doc-deploy

# Build the application
make build

# Restart the service
sudo service psn restart

echo "Application deployed!"