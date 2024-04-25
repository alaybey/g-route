# Modular Route Handling For Golang

This is a simple example of a Go web application demonstrating modular route handling using the net/http with golang 1.22+

## Overview

This project illustrates how to organize a Go web application with multiple controllers, each responsible for handling specific routes. The goal is to keep the codebase clean, modular, and easy to maintain.

## Features

- **Modular Route Handling:** Controllers are organized into separate files, making it easy to manage routes for different parts of the application.
- **Centralized Route Definition:** Routes are defined in a central location (routes.go), where controllers register themselves with the router.
- **Controller Interface:** All controllers implement a BaseController interface, ensuring consistency and structure across the application.
## Structure 
        .
        ├── controllers
        │ ├── BaseController.go
        │ └── *Controller.go # OtherControllers
        ├── main.go
        ├── README.md
        └── LICENCE.md

## Advantages
just add all controllers with single line like `controller.RegisterRoutes(mux)`
- **Less Code Complexity:** Just define a controller/handler on own file 
- **Simplified Maintenance**: Organizing routes into separate modules makes it easier to manage and maintain the codebase.
- **Enhanced Scalability**: Can scale individual modules independently, allowing for better resource allocation and improved performance.
- **Improved Developer Experience**: Developers can work on specific modules without needing to understand the entire application, reducing cognitive load and improving productivity.
- **Flexibility**: Modules can be added, updated, or removed with minimal impact on other parts of the application, providing flexibility for future enhancements.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License
This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.
