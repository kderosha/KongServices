#!/bin/bash
curl -X POST http://localhost:8000/services -d '{"name":"$1", "description":"test description"}'