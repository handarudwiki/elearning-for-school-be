# Makefile untuk membaca file .env untuk URL database

# Memuat variabel lingkungan dari file .env
ifneq ("$(wildcard .env)","")
    include .env
    export
endif

MIGRATIONS_DIR = database/migrations

# Perintah untuk menjalankan migrasi
up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DATABASE_URL) up

down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DATABASE_URL) down

# Perintah untuk membuat migrasi baru
# Menggunakan variabel `name` yang dikirimkan melalui perintah `make`
new:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

run:
	@go run main.go


