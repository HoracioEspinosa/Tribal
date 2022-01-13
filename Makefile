#!/bin/bash

build-dev:
	docker-compose --env-file .env up --build --remove-orphans -d
