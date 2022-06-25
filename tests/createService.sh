#!/bin/bash
curl -X POST http://localhost:8000/services -d '{"name":"test service 1", "description":"test description"}'