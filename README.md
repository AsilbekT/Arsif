# Arsif

**Arsif** is a Go framework designed to help developers quickly scaffold and build scalable Go applications with a modular architecture. It provides a command-line interface (CLI) that allows you to generate a structured project layout, similar to popular frameworks like Django and Ruby on Rails, but tailored for Go development.

## Features

- **Modular Architecture**: Organizes your project into clean, maintainable modules (apps) for different functionalities.
- **Easy Setup**: Quickly scaffold a new project with predefined folder structures.
- **Built-in CLI Tool**: Includes a CLI tool (`arsif`) to create projects and manage common tasks.
- **Flexible Configuration**: Easily manage configuration files (`config.yml`).
- **Integrated Dependency Injection**: Pre-configured support for dependency injection through Go's interfaces.
- **Ready for Microservices**: Support for message brokers, Kafka integration, and Docker/Kubernetes deployment.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/AsilbekT/arsif.git
    ```

2. Navigate to the project directory:

    ```bash
    cd arsif
    ```

3. Install dependencies:

    ```bash
    go mod download
    ```

## Usage

### 1. Initialize a New Project

Use the `arsif` CLI to scaffold a new project structure:

```bash
./arsif startproject <projectname>
```
This will generate a new project directory with the following structure:

```
<projectname>/
├── cmd/
│   └── <projectname>/
├── internal/
│   ├── apps/
│   ├── core/
│   └── infrastructure/
├── pkg/
├── configs/
├── deployments/
├── docs/
├── tests/
└── README.md
```

2. Project Layout
cmd/: The main application entry point.
internal/: Core business logic, services, and apps.
pkg/: Reusable packages across multiple projects.
configs/: Configuration files (e.g., environment-specific settings).
deployments/: Deployment configuration (Docker, Kubernetes).
tests/: Unit and integration tests.
docs/: Project documentation.

3. Running the Application
After setting up your project, you can run the application with:
```bash
go run ./cmd/<projectname>/
```

4. Build the Project
To build the project into a binary:


5. Docker Support
Build and run the project using Docker:

```bash
docker build -t <projectname> .
docker run -p 8080:8080 <projectname>
```
Contributing
Contributions are welcome! If you'd like to contribute, please fork the repository, create a new branch, and submit a pull request. Feel free to open issues for any bugs or feature requests.

License

This project is licensed under the MIT License. See the LICENSE file for more details.


