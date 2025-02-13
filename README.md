# Achilles: A Minimalist Go Web Application Boilerplate

Achilles is a lean, Go (GoLang)-based web application boilerplate crafted to streamline the initiation of your web projects. With built-in capabilities for logging, configuration management, and common web application tasks, Achilles frees you from mundane and essential chores. Channel your efforts into developing your application's core functionality efficiently.

### Disclaimer: Deviation from Go Standard Naming Conventions

Achilles follows a clear and streamlined naming structure specifically designed for the sake of clarity and minimalism. We acknowledge that the Go community values adherence to the standard Go packaging conventions, and Achilles intentionally deviates from these standards for a more intuitive and straightforward approach to structuring your code.

## Table of Contents
- [Introduction](#achilles-a-minimalist-go-web-application-boilerplate)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Highlights](#highlights)
- [Configuration Management](#configuration-management)
- [Circuit Breaker Pattern](#circuit-breaker-pattern)
- [Logging and Error Handling](#logging-and-error-handling)
- [Getting Started](#getting-started)
- [Adding a New Route](#adding-a-new-route-v1testsuccess-to-achilles)
- [Using the Clients](#using-the-clients)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Prerequisites

Before sailing with Achilles, ensure you have:

- Go 1.19 or higher
- Git for version control
- A text editor or IDE (VS Code recommended)
- Basic understanding of HTTP services
- PostgreSQL (if using RDBMS features)

## Quick Start

Get a basic Achilles server running in under 5 minutes:

```bash
# Clone the repository
git clone https://github.com/yourusername/achilles.git

# Navigate to project
cd achilles

# Install dependencies
go mod download

# Run the server
go run entrypoint/http/main.go
```

## Highlights

- **Intuitive Structuring**: If you appreciate clear and straightforward naming conventions, say goodbye to ambiguous structuring; Achilles might be your perfect choice.
- **Rapid Proof of Concept (POC)**: Ideal for exploring Go for a specific use case or an effortlessly scalable proof of concept.
- **Structured Logging**: Harness the flexibility of Logrus for logging, delivering structured JSON output. With log rotation based on file size and days handled through Lumberjack, your log management becomes more efficient.
- **Configuration Made Easy**: Customize your application by editing the `config.json` file in the `config` directory. Achilles leverages Viper for straightforward configuration management, making it a breeze to load and access your application's settings.
- **Common Web Application Tasks**: Achilles includes middleware for everyday web application tasks, such as request/response processing and error recovery, saving you from reinventing the wheel.
- **Resilient Circuit Breaker**: Built-in circuit breaker pattern using Hystrix to prevent cascading failures in both HTTP and RDBMS operations.
- **Database Integration**: Ready-to-use RDBMS client with connection pooling and circuit breaker protection.
- **HTTP Client Resilience**: Robust HTTP client with configurable timeouts, retries, and circuit breaker capabilities.

## Configuration Management

Achilles provides a flexible configuration system through JSON files in the `config` directory. Different environments can have their own configuration files:

### HTTP Client Configuration

Fine-tune your HTTP client behavior in the configuration file:

```json
{
  "http_config": {
    "request_timeout_ms": "2000ms",
    "circuit_breaker_name": "default",
    "error_threshold_percentage": 50,
    "retry_count_max": 2,
    "retry_backoff_ms": "100ms",
    "retry_jitter_ms": "100ms",
    "retry_duration_max": "1000ms",
    "request_volume_threshold": 5,
    "max_concurrent_requests": 10
  }
}
```

- **`request_timeout_ms`**: Maximum duration to wait for HTTP requests
- **`circuit_breaker_name`**: Unique identifier for the circuit breaker instance
- **`error_threshold_percentage`**: Percentage of errors that trigger circuit breaker
- **`retry_count_max`**: Maximum number of retry attempts
- **`request_volume_threshold`**: Minimum requests needed before tripping circuit breaker

### RDBMS Configuration

Configure your database connection and resilience settings:

```json
{
  "rdbms_config": {
    "dsn": "postgres://user:password@localhost:5432/dbname",
    "max_open_conns": 10,
    "max_idle_conns": 5,
    "conn_max_lifetime_ms": "3600000ms",
    "circuit_breaker_name": "defaultDB",
    "error_threshold_percentage": 50,
    "request_volume_threshold": 5,
    "max_concurrent_requests": 10,
    "retry_count_max": 2,
    "retry_backoff_ms": "100ms"
  }
}
```

- **`dsn`**: Database connection string
- **`max_open_conns`**: Maximum number of open connections
- **`max_idle_conns`**: Maximum number of idle connections
- **`conn_max_lifetime_ms`**: Maximum lifetime of connections

## Circuit Breaker Pattern

Achilles implements the circuit breaker pattern to enhance system resilience:

### States
1. **Closed**: Normal operation state where requests flow through
2. **Open**: When error threshold is exceeded, requests are blocked to prevent cascade failures
3. **Half-Open**: After a cooling period, allows test requests to check if the service has recovered

### Configuration Parameters
- **Error Threshold**: Percentage of failures that trigger the circuit breaker
- **Request Volume**: Minimum number of requests needed before measuring errors
- **Sleep Window**: Duration the circuit stays open before testing recovery
- **Max Concurrent**: Maximum number of concurrent requests allowed

## Logging and Error Handling

Achilles simplifies logging and error handling, offering configurability and powerful features to enhance your development experience.

- **Incoming Requests**: Log every incoming request, providing insights into the interactions.
- **Rendered Responses**: Capture details of the response, aiding in performance analysis and debugging.
- **Error Recovery**: Logs comprehensive information on crashes, simplifying the debugging and recovery process.
- **Indentation**: Supports indentation for better readability.

### Logging Configuration

Fine-tune your logging preferences in the `config.json` file:

```json
{
  "application": "Achilles",
  "application_id": 1,
  "http_port": 8080,
  "log": {
    "log_in_terminal_over_file": true,
    "enable_indentation": true,
    "log_path": "log/",
    "log_max_size_mb": 10,
    "log_max_backups": 5,
    "log_max_age_days": 28
  }
}
```

Configure the logging behavior of Achilles in the `config.json` file:

- **`log_in_terminal_over_file`**: Choice between logs in the terminal (stdout) or files.
- **`enable_indentation`**: Enable indentation for clearer log readability.
- **`log_path`**: Set the directory for log files. Default is the `log/` directory.
- **`log_max_size_mb`**: Control log rotation by setting the maximum size (in megabytes) for each log file.
- **`log_max_backups`**: Manage old log files with a maximum number of backups to retain.
- **`log_max_age_days`**: Specify the maximum number of days to retain old log files.

With Achilles, logging becomes a breeze, providing comprehensive insights, efficient log storage, and easy configuration.

## Getting Started

**If you can't wait or don't want to do a local setup at all, instantly try this** [instant cloud setup for free](https://medium.com/stackademic/go-app-from-github-to-cloud-in-10-minutes-c6622bdda2a1).

If you want to sail through the code, follow these fundamental steps:

1. **Clone this Repository**: Start by cloning this repository to your local machine.
2. **Install Go**: If you haven't already, install Go by following the [Go Installation Guide](https://golang.org/doc/install).
3. **Customize Configuration**: Tailor the `config.json` file in the `config` directory to match your application's specific requirements.
4. **Build and Run Your Application**: Use the following commands to build and run your application:

   ```bash
   go build ./entrypoint/http/main.go
   ./main
   ```

If you are facing any issue while setup, try [setting-up workspace](https://blog.stackademic.com/go-development-vs-code-workspace-be41470ca134).

## Adding a New Route (v1/test/success) to Achilles

1. **Locate the version-specific route** handler file. For example, let's consider `v1RouteHandler/v1.go`.
2. **Add a New Route Handler Function** - Create a new Go file, e.g., success.go, to handle the logic for the new route. For example:

```go
// success.go
package v1RouteHandler
import (
    "achilles/constant"
    routeHelper "achilles/route/helper"
    "github.com/gin-gonic/gin"
)

// Success handles the new success route.
func Success(ginContext *gin.Context) {
    routeHelper.SetSuccessResponseWithOnlyMessage(ginContext, constant.HttpOk)
}
```

3. **Add the new route handler** to the `AddRouteHandlers` function:

```go
// AddRouteHandlers adds route handlers to the version 1 group.
func AddRouteHandlers() {
    // ... (existing routes)

    versionOne.GET("test/success", Success) // Add the new route here
}
```

4. **Test Your Route** - Build and run your application to test the new route:

```bash
go build ./entrypoint/http/main.go
./main
```

Your new route should now be accessible. Congratulations on adding a new route to Achilles!

## Using the Clients

Achilles provides two approaches to client initialization, with a deliberate emphasis on direct initialization.

### Direct Initialization (Recommended)

This approach follows the dependency inversion principle and is recommended because it:
- Makes dependencies explicit and visible where they're used
- Simplifies testing by allowing easy mocking of dependencies
- Reduces hidden coupling between components
- Keeps the code straightforward and maintainable
- Provides better control over client lifecycle

```go
import (
    "achilles/client"
    "achilles/model"
)

// Initialize HTTP client with config
httpConfig := model.ClientHTTPConfig{
    RequestTimeoutDuration: 2000 * time.Millisecond,
    CircuitBreakerName:    "default",
    // ... other config options
}
httpClient, err := client.NewHTTPClient(httpConfig, logger)
```

### Alternative: Dependency Container

While Achilles includes a dependency container, it's provided as an optional feature for cases where:
- You need centralized dependency management
- Your application grows to require complex dependency graphs
- You want to implement service location pattern

```go
// Available but not mandatory
container := helper.NewDependencyContainer(config)
httpClient := container.GetHTTPClient()
```

The container pattern, while powerful, can hide dependencies and make code harder to test. That's why we recommend starting with direct initialization and only moving to the container pattern when your application's complexity truly demands it.

This dual approach allows you to:
- Start simple with direct initialization
- Scale to dependency injection when needed
- Keep your codebase flexible and maintainable
- Make conscious decisions about dependency management

## License

Achilles is open-source software distributed under the [MIT License](LICENSE). You're welcome to use, modify, and distribute it in your projects.

## Acknowledgments

Achilles is built upon several outstanding open-source libraries:
- Hystrix for circuit breaker implementation
- Logrus for structured logging
- Viper for configuration management
- Gin for HTTP routing
- SQLx for enhanced database operations

I appreciate your interest in Achilles and hope it accelerates your Go web application development.
