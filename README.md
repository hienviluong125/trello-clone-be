##### Prerequisites

The setups steps expect following tools installed on the system.

- Github
- Go [1.16]
- Postgresql [9.x]
##### 1. Check out the repository
```bash
git clone git@github.com:hienviluong125/trello-clone-be.git
```

##### 2. Setup environment variables
Run this command to setup env vars
```bash
export DATABASE_URL="host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai" GO_APP_ENV="development" JWT_SECRET="SECRET"
```

##### 3. Start a server
```bash
go run main.go
```
