# FizzBuzz

# Specifications
```
Write a simple fizz-buzz REST server.

"The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"".
The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...""."

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request
```

# Installation and Setup

### Requirements
- GNU Make
- Docker
- Golang (if not using Docker)


## Usage

The Makefile allows you to perform the following actions:
1. Test
2. Fmt
3. Lint
4. Install
5. Build
6. Run

### Building the Docker Image

To build the Docker image, run:
```sh
docker build -t fizzbuzz .
```
or just run `make build`

### Running the API
Then, to start the API on port 8080:
```sh
 docker run -p 8080:8080 fizzbuzz
```
or just run `make run`

Instead of you can run `make execute` which will compile and build the image first (no cache enabled)

## Running Tests
In order to test the code, just run `make test`.
