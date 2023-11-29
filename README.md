# Achilles: A Minimalist Go Web Application Boilerplate

Achilles is a lean, Go ( GoLang )-based web application boilerplate crafted to streamline the initiation of your web projects. With built-in capabilities for logging, configuration management, and common web application tasks, Achilles frees you from mundane and essential chores. Channel your efforts into developing your application's core functionality efficiently.

### Disclaimer: Deviation from Go Standard Naming Conventions

Achilles follows a clear and streamlined naming structure specifically designed for the sake of clarity and minimalism. We acknowledge that the Go community values adherence to the standard Go packaging conventions, and Achilles intentionally deviates from these standards for a more intuitive and straightforward approach to structuring your code.


## Highlights

- **Intuitive Structuring**: If you appreciate clear and straightforward naming conventions, Say goodbye to ambiguous structuring; Achilles might be your perfect choice. 
- **Rapid Proof of Concept (POC)** : Ideal for exploring Go for a specific use case or an effortlessly scalable proof of concept.
- **Structured Logging**: Harness the flexibility of Logrus for logging, delivering structured JSON output. With log rotation based on file size and days handled through Lumberjack, your log management becomes more efficient.
- **Configuration Made Easy**: Customize your application by editing the `config.json` file in the `config` directory. Achilles leverages Viper for straightforward configuration management, making it a breeze to load and access your application's settings.
- **Common Web Application Tasks**: Achilles includes middleware for everyday web application tasks, such as request/response processing and error recovery, saving you from reinventing the wheel.

## Logging and Error Handling

Achilles simplifies logging and error handling, offering configurability and powerful features to enhance your development experience.

- **Incoming Requests:** Log every incoming request, providing insights into the interactions.
- **Rendered Responses:** Capture details of the response, aiding in performance analysis and debugging.
- **Error Recovery:** Logs comprehensive information on crashes, simplifying the debugging and recovery process.
- **Indentation:** Supports indentation for better readability.

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

- **`log_in_terminal_over_file`**: choice between logs in the terminal (stdout) or files.
- **`enable_indentation`**: Enable indentation for clearer log readability.
- **`log_path`**: Set the directory for log files. Default is the `log/` directory.
- **`log_max_size_mb`**: Control log rotation by setting the maximum size (in megabytes) for each log file.
- **`log_max_backups`**: Manage old log files with a maximum number of backups to retain.
- **`log_max_age_days`**: Specify the maximum number of days to retain old log files.

With Achilles, logging becomes a breeze, providing comprehensive insights, efficient log storage, and easy configuration.


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


# Adding a New Route ( v1/test/success ) to Achilles

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
 3. **add the new route handle**r to the `AddRouteHandlers` function:
```go
// AddRouteHandlers adds route handlers to the version 1 group.
func AddRouteHandlers() {
    // ... (existing routes)

    versionOne.GET("test/success", Success) // Add the new route here
}
```
 4. **Test Your Route** - build and run your application to test the new route:

```bash
go build ./entrypoints/http/main.go
./main
```

Your new route should now be accessible. Congratulations on adding a new route to Achilles!


## License

Achilles is open-source software distributed under the [MIT License](LICENSE). You're welcome to use, modify, and distribute it in your projects.

## Acknowledgments

Achilles is built upon several open-source libraries and tools. Gratitude to the creators and maintainers of these fantastic projects.

I appreciate your interest in Achilles and hope it accelerates your Go web application development.
