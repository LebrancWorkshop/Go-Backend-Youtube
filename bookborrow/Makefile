container:
	docker run --name bookborrow -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=777 -d postgres:alpine

postgres:
	docker exec -it bookborrow psql -U root

.PHONY: container