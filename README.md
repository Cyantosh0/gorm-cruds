# gorm-cruds

gorm-cruds is a Go application that demonstrates CRUD (Create, Read, Update, Delete) operations using the GORM ORM library. It provides a simple and straightforward implementation of basic CRUD functionality, serving as a reference or starting point for developers working with Go and GORM.

## Features

- **User Management**: Create, retrieve, update, and delete user records.
- **Database Integration**: Seamless integration with popular databases supported by GORM (e.g., PostgreSQL, MySQL, SQLite).
- **Environment Configuration**: Utilizes the Viper library for easy environment configuration management.
- **CLI Interface**: Includes a command-line interface (CLI) powered by Cobra for executing various commands and operations.

## Getting Started

### Prerequisites

- Go programming language (version 1.16 or later)
- Appropriate database driver (e.g., `github.com/lib/pq` for PostgreSQL)

### Installation

1. Clone the repository:
```git clone https://github.com/Cyantosh0/gorm-cruds.git```

2. Navigate to the project directory:
```cd gorm-cruds```

3. ### Configuration

1. Create a `.env` file in the project root directory based on the `.env.example` file.
2. Update the environment variables in the `.env` file with your database connection details.

### Running the Application

Start the application: ```go run main.go```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

This README file provides an overview of the gorm-cruds application, including its features, installation instructions, configuration guidelines, and usage examples. It also encourages contributions and specifies the project's license.
