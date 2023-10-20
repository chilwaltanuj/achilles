# Achilles: A Minimalist Go Web Application Boilerplate

Achilles is a lean, Go-based web application boilerplate crafted to streamline the initiation of your web projects. With built-in capabilities for logging, configuration management, and common web application tasks, Achilles frees you from mundane and essential chores, enabling you to channel your efforts into developing your application's core functionality efficiently.

### Disclaimer: Deviation from Go Standard Naming Conventions

Achilles follows a clear and streamlined naming structure specifically designed for the sake of clarity and minimalism. We acknowledge that the Go community values adherence to the standard Go packaging conventions, and Achilles intentionally deviates from these standards for a more intuitive and straightforward approach to structuring your code.


## Highlights

Achilles simplifies your web application development in several ways:

- **Intuitive Structuring**: If you appreciate clear and straightforward naming conventions, Achilles might be your perfect choice. Say goodbye to ambiguous structuring; It places everything right where it belongs.

- **Rapid Proof of Concept (POC)** : Whether you're exploring Go for a specific use case or demand an effortlessly scalable proof of concept, Achilles is your reliable companion.


- **Structured Logging**: Harness the flexibility of Logrus for logging, delivering structured JSON output. With log rotation based on file size and days handled through Lumberjack, your log management becomes more efficient.

- **Configuration Made Easy**: Customize your application by editing the `config.json` file in the `config` directory. Achilles leverages Viper for straightforward configuration management, making it a breeze to load and access your application's settings.

- **Common Web Application Tasks**: Achilles includes middleware for everyday web application tasks, such as request/response processing and error recovery, saving you from repetitive code.


## Getting Started

To initiate your journey with Achilles, follow these fundamental steps:

1. **Clone this Repository**: Start by cloning this repository to your local machine.

2. **Install Go**: If you haven't already, install Go by following the [Go Installation Guide](https://golang.org/doc/install).

3. **Customize Configuration**: Tailor the `config.json` file in the `config` directory to match your application's specific requirements.

4. **Build and Run Your Application**: Use the following commands to build and run your application:

   ```bash
   go build ./entrypoints/http/main.go
   ./main
   ```

## Use Cases

Here are some real-world scenarios where Achilles can be an excellent fit:

- **Swift Prototyping**: Utilize Achilles to expedite the creation of web application prototypes or microservices with minimal overhead.

- **Streamlined Development**: Developers who appreciate straightforward structuring and clear naming conventions will find Achilles a breath of fresh air.

- **Customization**: Easily customize your application's behavior by tweaking the `config.json` file in the `config` directory.

- **Accelerated Development**: Achilles simplifies the handling of common web application tasks and provides a well-structured foundation, facilitating swift development.

## License

Achilles is open-source software distributed under the [MIT License](LICENSE). You're welcome to use, modify, and distribute it in your projects.

## Acknowledgments

Achilles is built upon several open-source libraries and tools. Gratitude to the creators and maintainers of these fantastic projects.

I appreciate your interest in Achilles and hope it accelerates your Go web application development.
