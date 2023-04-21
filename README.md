# Mini-Aspire API
Mini-Aspire is an app that allows authenticated users to go through a loan application. This app is a stateless app, which means the data is always stored in memory and not being stored in a persistent database.

## Installation

### Dependencies
- [Go 1.19](https://go.dev/doc/install)

### Setup Mini-Aspire
1. Clone this repository
```
git clone git@github.com:rakhmatullahyoga/mini-aspire.git
```
2. Install libraries
```
make dep
```

## Running the application
### Starting the REST API server
Start the REST API server by running the following command
```
make run
```
or, we can run immediately by running:
```
make
```

### Unit Test
Run unit test by running the following command
```
make test
```

### Test API
1. Import the [Postman collection](Mini-Aspire.postman_collection.json) attached in the repository
2. Test all scenarios
