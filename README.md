# Mini-Aspire API
Mini-Aspire is an app that allows authenticated users to go through a loan application.

## Installation

### Dependencies
- [PHP 8.2.5](https://www.php.net/manual/en/install.php)
- [Composer](https://getcomposer.org/download/)
- [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)

### Setup Mini-Aspire
1. Clone this repository
```
git clone git@github.com:rakhmatullahyoga/mini-aspire.git
```
2. Install libraries
```
composer install
```
3. Set environment variables and adjust values
```
cp .env.example .env
```
4. Setup database
```
php artisan migrate --seed
```

## Running the application
Start the REST API server by running the following command
```
php artisan serve
```

### Unit Test
Run unit test by running the following command
```
php artisan test
```

### Test API
1. Import the [Postman collection](Mini-Aspire.postman_collection.json) attached in the repository
2. Hit Login API with the given users credentials in the [database seeder file](database/seeders/DatabaseSeeder.php)
3. Test all scenarios
