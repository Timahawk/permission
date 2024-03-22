
generate:
	go generate ./...
	atlas migrate diff initial --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "docker://postgres/16/?sslmode=disable&search_path=public"
	atlas migrate lint --dev-url="docker://postgres/16/?sslmode=disable&search_path=public" --dir="file://ent/migrate/migrations" --latest=1
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "postgres://postgres:postgres@localhost:5432/postgres?search_path=public&sslmode=disable"