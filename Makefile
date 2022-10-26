dep:
	go mod tidy
	go mod vendor

migrate:
	npm i
	npx sequelize db:migrate

run:
	go run app/api/main.go

createdb:
	npx sequelize db:create

undo:
	npx sequelize db:migrate:undo:all   