.PHONY: run
COMPOSE=docker-compose
DOCKER=docker
DB=todo_database
DB_URI="mysql://root:my-secret-pw@tcp(127.0.0.1:33063)/todo_db"
MIGRATE_FILES=mysql/migrations
v=

init:
	brew install golang-migrate
	go install github.com/cosmtrek/air@latest&&air -v

go/run:
	cd go && go run main.go

go/run/air:
	air

go/build:
	cd go && go build -o ../tmp/main main.go

docker/up:
	$(COMPOSE) up -d --build

docker/build:
	$(COMPOSE) build

# sh/go:
# 	$(DOCKER) exec -it $(GO) ash

sh/db:
	$(DOCKER) exec -it $(DB) bash

docker/logs:
	$(COMPOSE) logs

docker/ps:
	$(COMPOSE) ps

# restart/go:
# 	$(COMPOSE) restart $(GO)

restart/db:
	$(COMPOSE) restart $(DB)

docker/down:
	$(COMPOSE) down -v

# down/go:
# 	$(COMPOSE) rm -fsv $(GO)

docker/down/db:
	$(COMPOSE) rm -fsv $(DB)

migrate/up:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} up ${v}

migrate/down:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} down ${v}

migrate/force:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} force ${v}

migrate/updown: migrate/down migrate/up
