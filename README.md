# Gyad - Get Your API Data


## Dev Setup
**Step 1: Clone the repository**
```bash
git clone https://github.com/ximmanuel/gyad.git
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

**Rollback/Delete specific migratrion**:
```bash
go run cmd/migration/migration.go delete bober
```
