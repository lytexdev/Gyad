# Gyad - Get Your API Data

## Overview
A lightweight backend system built in Go designed to simplify the management and access to data through APIs.
An example model, migration and controller have been provided for demonstration purposes

## Features
- **Database Migrations**: Easily create, manage, and rollback database migrations using simple command-line operations.
- **ORM**: A simple ORM that allows you to interact with the database using Go structs [(XORM)](https://xorm.io/).
- **API Controllers**: Facilitate and manage RESTful APIs that enable clients to interact with the backend.


## Installation

### Prerequisites
- Go (version 1.22 or higher)
- A PostgreSQL Database

**Clone the repository**
```bash
git clone https://github.com/lytexdev/Gyad.git
```

**Rename the `.env.example` file to `.env` and adjust the values**
```bash
cp .env.example .env
```

**Run the web server**
```bash
go run cmd/main.go
```

## Usage

### Migration
**Create Migration**
```bash
go run cmd/migration/migration.go create bober
```

**Run all migrations**:
```bash
go run cmd/migration/migration.go migrate all
```
Migrations are executed one after the other based on the timestamps.

**Run specific migratrion**:
```bash
go run cmd/migration/migration.go migrate bober
```

**Rollback specific migration**:
```bash
go run cmd/migration/migration.go rollback bober
```

**Create migration binary file**:
```bash
go build cmd/migration/migration.go
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.