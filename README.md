# REST API with Golang and Gin

This repository contains a REST API built using Golang and the Gin framework. The API provides a robust and scalable architecture for handling HTTP requests and responses efficiently. 
# RestAPI

This project is a RESTful API built with Go and the Gin web framework. It provides a simple event registration system where users can create, update, delete, and register for events.

## Features

- User authentication and authorization
- CRUD operations for events
- User registration for events
- Token-based authentication using JWT

## Tech Stack

- Go
- Gin web framework
- SQLite for database
- bcrypt for password hashing
- JWT for token generation and verification

## Project Structure

The project is structured into several packages:

- `models`: Contains the data models (User, Event) and their associated methods for database operations.
- `routes`: Contains the route handlers for the API endpoints.
- `middlewares`: Contains middleware functions for tasks such as user authentication.
- `utils`: Contains utility functions for tasks such as password hashing and token generation.

## Setup and Run

1. Clone the repository to your local machine.
2. Ensure you have Go installed on your machine.
3. Navigate to the project directory and run `go mod download` to download the necessary dependencies.
4. Run `go run main.go` to start the server.

The server will start on `http://localhost:8080`.

## API Endpoints

- `POST /signup`: Endpoint for user signup. Expects a JSON body with `email` and `password`.
- `POST /login`: Endpoint for user login. Expects a JSON body with `email` and `password`.
- `GET /events`: Fetches all events.
- `GET /events/:id`: Fetches a specific event by ID.
- `POST /events`: Creates a new event. Requires authentication.
- `PUT /events/:id`: Updates a specific event. Requires authentication.
- `DELETE /events/:id`: Deletes a specific event. Requires authentication.
- `POST /events/:id/register`: Registers the authenticated user for a specific event.
- `DELETE /events/:id/register`: Cancels the authenticated user's registration for a specific event.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
## Getting Started

To get started with this project, clone the repository using the following command:

```bash
git clone https://github.com/hvbhanot/RestAPI
