create-migration:
	goose postgres "user=root dbname=hrtools sslmode=disable host=localhost port=5432" create $(NAME) $(TYPE)

migrate-up:
	goose postgres "user=root password=root dbname=hrtools sslmode=disable host=localhost port=5432" up