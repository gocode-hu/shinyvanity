# Shiny Vanity

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Shiny Vanity is a simple, Go vanity import server. It allows you to use custom Go import paths for your open-source Go modules, redirecting users and `go get` requests to your actual GitHub repositories. This is especially useful for organizations and individuals who want to present a unified import path under their own domain.

## Features

- **Go vanity import support**: Serve custom Go import meta tags for your domain.
- **Automatic redirect**: Non-`go-get` requests are redirected to the corresponding GitHub repository.
- **Configurable via environment variables**: No code changes required for new domains or organizations.
- **Docker-ready**: Easily deployable anywhere.
- **MIT Licensed**: Open source and free to use.

## How It Works

- Requests with `?go-get=1` receive an HTML page with the correct `<meta name="go-import">` tag for Go tooling.
- All other requests are redirected to the corresponding GitHub repository.

## Quick Start

### 1. Clone the repository

```sh
git clone https://github.com/your-org/shiny-vanity.git
cd shiny-vanity
```

### 2. Set environment variables

- `VANITY_DOMAIN`: The domain to use in the import path (e.g., `example.com`)
- `VANITY_ORGANIZATION`: The GitHub organization or user (e.g., `your-org`)
- `LISTEN_PORT`: The port to listen on (e.g., `8080`)

You can use a `.env` file or set them in your environment:

```sh
export VANITY_DOMAIN=example.com
export VANITY_ORGANIZATION=your-org
export LISTEN_PORT=8080
```

### 3. Run the server

```sh
go run ./cmd/shinyvanity/main.go
```

Or build and run:

```sh
go build -o shiny-vanity ./cmd/shinyvanity
./shiny-vanity
```

### 4. Example Usage

For a Go module at `example.com/foo/bar`, a request to:

```
GET http://example.com/foo/bar?go-get=1
```

will return:

```
<meta name="go-import" content="example.com/foo/bar git https://github.com/your-org/foo/bar">
```

A browser request to `http://example.com/foo/bar` will redirect to `https://github.com/your-org/foo/bar`.

## Docker

A `Dockerfile` is provided for easy containerization:

```sh
docker build -t shiny-vanity .
docker run -e VANITY_DOMAIN=example.com -e VANITY_ORGANIZATION=your-org -e LISTEN_PORT=8080 -p 8080:8080 shiny-vanity
```

## Testing

- **Unit tests**: Run with `go test ./...`
- **Integration tests**: See `tests/integration_test.go` for end-to-end server tests.

## Project Structure

```
cmd/shinyvanity/      # Main entrypoint
internal/handler/    # Vanity handler logic
main.go              # (root) for legacy/compatibility
Dockerfile           # Container build
LICENSE              # MIT License
```

## Contributing

Pull requests and issues are welcome! Please open an issue to discuss your ideas or report bugs.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

