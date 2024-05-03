# Bitcoin Last Traded Price API

This API retrieves the Last Traded Price (LTP) of Bitcoin for the currency pairs BTC/USD, BTC/CHF, and BTC/EUR using the Kraken public API.

## Prerequisites

- Go 1.16 or later
- Docker (optional)

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/atkachyshyn/ltp.git
cd ltp
```

## Building and Running the Application

### Building the Application

To build the application, you can use the Makefile:

```bash
make build
```

This command compiles the Go source code located in the `cmd` directory and creates an executable named `main` in the project root directory.

### Running the Application

To start the server on your local machine:

```bash
make run
```

This command will first ensure that the application is built and then start the server. The server listens on the configured port, defaulting to 8080.

## Using Docker

### Building the Docker Container

To build the Docker container with the application, run:

```bash
make docker-build
```

This command builds a Docker image tagged `bitcoin-ltp`.

### Running the Docker Container

To run the Docker container locally:

```bash
make docker-run
```

This command runs the Docker container based on the `bitcoin-ltp` image, exposing the application on port 8080 of your host machine.

## API Usage

Once the application is running, you can fetch the latest traded prices using the following endpoint:

```
http://localhost:8080/api/v1/ltp
```

This endpoint returns a JSON response with the latest prices of BTC in USD, EUR, and CHF.

### Example with `curl`

To request the LTP data using `curl` from the command line:

```bash
curl http://localhost:8080/api/v1/ltp
```

This command sends a GET request to the API and displays the JSON response in the terminal.

## Testing

To run tests, use the Makefile:

```bash
make test
```

This command runs all defined tests in the project, ensuring that the functionality meets the expected behavior.

## Cleaning Up

To clean up the build artifacts and other generated files:

```bash
make clean
```

This command removes any compiled executables and other temporary files created during the build process.

## Authors

- Andriy Tkachyshyn - Initial work

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
