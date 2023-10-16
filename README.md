# Achilles: A Go Web Application Boilerplate

Achilles is a minimalist Go web application boilerplate designed to help you start your web projects efficiently. It comes with built-in features for logging, configuration management, and common web application tasks, allowing you to focus on building your application's core functionality.

## Real Use Cases

- **Start Coding Instantly**: Achilles provides a ready-to-use project structure, allowing you to dive right into coding without worrying about Go project layout and structuring.

- **Structured Logging**: Achilles uses [Logrus](https://github.com/sirupsen/logrus) for flexible logging with structured JSON output. Log rotation based on file size and days is handled through [Lumberjack](https://github.com/natefinch/lumberjack), ensuring efficient log management.

- **Configuration Made Easy**: Customize your application by editing the configuration file in the `config` directory. Achilles leverages [Viper](https://github.com/spf13/viper) for easy configuration management, so you can quickly load and access your application's settings.

- **Common Web Application Tasks**: Achilles includes middleware for common web application tasks like request/response processing and error recovery, saving you from writing repetitive code.

## Getting Started

1. **Clone this repository.**

2. **Install Go** if you haven't already: [Go Installation Guide](https://golang.org/doc/install).

3. **Customize Configuration**: Modify `config.json` in the `config` directory to tailor it to your application's requirements.

4. **Build and Run Your Application**:

   ```bash
   go build ./entrypoints/http/main.go
   ./main
