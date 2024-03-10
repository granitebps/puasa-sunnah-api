#!/bin/sh
set -e

echo "Deploying application ..."

# Update codebase
git fetch origin main
git reset --hard origin/main

# Build the application
make build-app
make build-backup

# Restart the service
sudo service puasasunnah restart

echo "Application deployed!"