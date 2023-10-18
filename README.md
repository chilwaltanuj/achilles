# Achilles: A Minimalist Go Web Application Boilerplate

Achilles is a minimalist Go web application boilerplate designed to help you start your web projects efficiently and quickly. It comes with built-in features for logging, configuration management, and common web application tasks, allowing you to focus on building your application's core functionality.

## Project Overview

Achilles simplifies your web application development in several ways:

- **Intuitive Structuring**: If you appreciate clear and straightforward naming conventions, Achilles might be your perfect choice. Say goodbye to ambiguous structuring; It places everything right where it belongs.

- **Quick Proof of Concept (POC)**: Whether you're exploring Go for a specific use case or need a scalable POC with zero effort, Achilles has you covered.

- **Structured Logging**: Harness the flexibility of Logrus for logging, delivering structured JSON output. With log rotation based on file size and days handled through Lumberjack, your log management becomes more efficient.

- **Configuration Made Easy**: Customize your application by editing the `config.json` file in the `config` directory. Achilles leverages Viper for straightforward configuration management, making it a breeze to load and access your application's settings.

- **Common Web Application Tasks**: Achilles includes middleware for everyday web application tasks, such as request/response processing and error recovery, saving you from repetitive code.


## Getting Started

To begin using Achilles, follow these steps:

1. **Clone this Repository**: Start by cloning this repository to your local machine.

2. **Install Go**: If you haven't already, install Go by following the [Go Installation Guide](https://golang.org/doc/install).

3. **Customize Configuration**: Tailor the `config.json` file in the `config` directory to match your application's specific requirements.

4. **Build and Run Your Application**: Use the following commands to build and run your application:

   ```bash
   go build ./entrypoints/http/main.go
   ./main
   ```

## Real Use Cases

Here are some real-world scenarios where Achilles can be an excellent fit:

- **Quick Prototyping**: Use Achilles to create quick prototypes of web applications or microservices with minimal overhead.

- **Simplified Development**: Developers who appreciate straightforward structuring and clear naming conventions will find Achilles a breath of fresh air.

- **Customization**: Easily customize your application's behavior by tweaking the `config.json` file in the `config` directory.

- **Faster Development**: Achilles handles common web application tasks and provides a structured starting point, enabling faster development.

## Contribution

We welcome contributions to Achilles! If you'd like to improve the project, report issues, or submit pull requests, please refer to our [Contribution Guidelines](CONTRIBUTING.md).

## License

Achilles is open-source software distributed under the [MIT License](LICENSE). You're welcome to use, modify, and distribute it in your projects.

## Acknowledgments

Achilles is built upon several open-source libraries and tools. We want to extend our gratitude to the creators and maintainers of these fantastic projects.

We appreciate your interest in Achilles and hope it accelerates your Go web application development.
