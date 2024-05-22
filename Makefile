postgresinit:
	docker run --name pet_shop -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
createdb:
	docker exec -it pet_shop createdb --username=root --owner=root pet_shop
dropdb:
	docker exec -it pet_shop dropdb pet_shop

.PHONY: postgresinit createdb dropdb 
