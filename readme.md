# Go & Gin Authentication Using Docker and  Postgres

## Overview
This project is a web application built using [Go](https://golang.org/) and the [Gin](https://github.com/gin-gonic/gin) framework. It includes features such as user authentication and database integration with PostgreSQL. Docker is used to containerize the application and manage dependencies.

## Features
- User registration and login endpoints
- PostgreSQL database integration
- Dockerized setup for easy deployment
- Environment-based configuration using `.env`

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.23 or higher)
- [Docker](https://www.docker.com/) and Docker Compose
- PostgreSQL client tools (optional for local debugging)

## Getting Started

### 1. Clone the Repository
```bash
git clone <repository-url>
cd <repository-name>
```

### 2. Configure Environment Variables
Create a `.env` file in the root of the project and add the following configurations:
```env
DATABASE_HOST=db
DATABASE_PORT=5432
DATABASE_USERNAME=admin
DATABASE_PASSWORD=admin
DATABASE_NAME=mydatabase
SERVER_ADDR=0.0.0.0:8080
```

### 3. Build and Run the Application
Using Docker Compose:
```bash
docker-compose up --build
```

This will:
- Build the Go application.
- Start the application and PostgreSQL database in separate containers.

The application will be accessible at `http://localhost:8080`.

### 4. API Endpoints
The following API endpoints are available:

#### **POST /auth/register**
Registers a new user.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
    "message": "User registered successfully."
  }
  ```

#### **POST /auth/login**
Logs in a user and returns a JWT token.
- **Request Body**:
  ```json
  {
    "email": "john@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
    "token": "<JWT_TOKEN>"
  }
  ```

### 5. Accessing the Database
To connect to the PostgreSQL database locally:
```bash
psql -h localhost -U admin -d mydatabase
```

## Project Structure
```
.
├── config         # Configuration files and database connection logic
├── controller     # Handlers for API endpoints
├── models         # Database models
├── routes         # API route definitions
├── Dockerfile     # Docker image definition
├── docker-compose.yml # Docker Compose setup
├── main.go        # Entry point for the application
└── .env.example   # Example environment variables file
```

## Technologies Used
- **Go**: Programming language
- **Gin**: HTTP web framework for Go
- **PostgreSQL**: Relational database
- **Docker**: Containerization

## Troubleshooting

### Common Issues
- **Application not accessible**: Ensure the Go server binds to `0.0.0.0` in the `SERVER_ADDR` environment variable.
- **Database connection failure**: Verify the PostgreSQL service is running and the `.env` file contains the correct credentials.

### Logs
Use the following command to view application logs:
```bash
docker-compose logs app
```

## Contributing
1. Fork the repository.
2. Create a new branch for your feature/bugfix.
3. Commit your changes and push the branch.
4. Open a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments
- [Gin Documentation](https://github.com/gin-gonic/gin)
- [Docker Documentation](https://docs.docker.com/)

