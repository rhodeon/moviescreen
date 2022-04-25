#!/bin/bash
migrate -path ./../../migrations/ -database ${TEST_DB_DSN} up
