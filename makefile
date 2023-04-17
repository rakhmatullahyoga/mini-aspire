setup-db:
	@php artisan migrate --seed

setup-env:
	@cp .env.example .env

install:
	@composer install

setup: install setup-env setup-db

serve: setup
	@php artisan serve

test: setup
	@php artisan test
