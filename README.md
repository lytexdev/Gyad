# Gyad - Get Your API Data

## Overview
A lightweight backend system built in Go designed to simplify the management and access to data through APIs.

## Features
- **Database Migrations**: Easily create, manage, and rollback database migrations using simple command-line operations.
- **TODO: API Controllers**: Facilitate and manage RESTful APIs that enable clients to interact with the backend.
- **TODO: DAL**: A Database Abstraction Layer that abstracts and simplifies database interactions. 

### Prerequisites
- Go (version 1.15 or higher)
- A PostgreSQL Database

## Installation
**Step 1: Clone the repository**
```bash
git clone git@github.com:ximmanuel/Gyad.git
```

**Step 2: Run postgres docker container**
```bash
docker run --name postgresql -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres
```

**Step 3: Rename the `.env.example` file to `.env` and adjust the values**
```bash
cp .env.example .env
```

**Step 4: Run the application**
```bash
go run cmd/main.go
```

## Migration

**Create Migration**
```bash
go run ./migration create bober
```

**Run all migrations**:
```bash
go run ./migration migrate all
```
Migrations are executed one after the other based on the timestamps.

**Run specific migratrion**:
```bash
go run ./migration migrate bober
```

**Rollback specific migration**:
```bash
go run ./migration rollback bober
```

## License
Distributed under the Mozilla Public License - see the [LICENSE](LICENSE) file for details.
