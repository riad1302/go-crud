# Golang Project
## About the project
Fetch data from database then generate csv. After csv generate update data in database.

## Project installation instruction

    cd docker
    cp .env.example .env
    cp docker-compose.override.example.yml docker-compose.override.yml
    cd /.envs
    cp app.env.example app.env
    docker-compose build
    docker-compose up -d