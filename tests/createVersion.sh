#!/bin/bash
curl -X POST http://localhost:8000/services/$1/versions -d '{"version":"1.0.0"}'