#!/usr/bin/env bash
set -euo pipefail

echo "🚀 Starting Toppira backend..."

echo "📄 Generating Swagger documents..."
swag init -o ./docs -g ./cmd/http/main.go --pd

echo "🧱 Generating repositories..."
go run ./cmd/gen

echo "🧪 Running application..."
go run ./cmd/http
