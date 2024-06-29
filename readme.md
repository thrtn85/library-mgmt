## Bookstore API Documentation

### Overview

The Bookstore API provides endpoints to manage books and users. It allows you to create, retrieve, update, and delete books and users. The API uses JSON for data exchange and is built using Go, Gin, and GORM with SQLite as the database.

### Base URL

```
http://localhost:1111
```

### Endpoints

#### Books

- **Create a Book**
  - **URL**: `/books`
  - **Method**: `POST`
  - **Description**: Creates a new book.
  - **Request Body**:
    ```json
    {
      "title": "string",
      "author": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "int",
      "title": "string",
      "author": "string",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

  - **Example**:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"title\":\"The Go Programming Language\", \"author\":\"Alan A. A. Donovan\"}" http://localhost:1111/books
    ```

- **Retrieve All Books**
  - **URL**: `/books`
  - **Method**: `GET`
  - **Description**: Retrieves a list of all books.
  - **Response**:
    ```json
    [
      {
        "id": "int",
        "title": "string",
        "author": "string",
        "created_at": "timestamp",
        "updated_at": "timestamp"
      }
    ]
    ```

  - **Example**:
    ```bash
    curl http://localhost:1111/books
    ```

- **Update a Book**
  - **URL**: `/books/update/:id`
  - **Method**: `PUT`
  - **Description**: Updates an existing book by ID.
  - **Request Body**:
    ```json
    {
      "title": "string",
      "author": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "int",
      "title": "string",
      "author": "string",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

  - **Example**:
    ```bash
    curl -X PUT -H "Content-Type: application/json" -d "{\"title\":\"Updated Title\", \"author\":\"Updated Author\"}" http://localhost:1111/books/update/1
    ```

- **Delete a Book**
  - **URL**: `/books/:id`
  - **Method**: `DELETE`
  - **Description**: Deletes an existing book by ID.
  - **Response**:
    ```json
    {
      "message": "Book deleted successfully"
    }
    ```

  - **Example**:
    ```bash
    curl -X DELETE http://localhost:1111/books/1
    ```

#### Users

- **Create a User**
  - **URL**: `/users`
  - **Method**: `POST`
  - **Description**: Creates a new user.
  - **Request Body**:
    ```json
    {
      "name": "string",
      "email": "string",
      "password": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "int",
      "name": "string",
      "email": "string",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

  - **Example**:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"John Doe\", \"email\":\"john@example.com\", \"password\":\"password\"}" http://localhost:1111/users
    ```

- **Retrieve All Users**
  - **URL**: `/users`
  - **Method**: `GET`
  - **Description**: Retrieves a list of all users.
  - **Response**:
    ```json
    [
      {
        "id": "int",
        "name": "string",
        "email": "string",
        "created_at": "timestamp",
        "updated_at": "timestamp"
      }
    ]
    ```

  - **Example**:
    ```bash
    curl http://localhost:1111/users
    ```

- **Update a User**
  - **URL**: `/users/update/:id`
  - **Method**: `PUT`
  - **Description**: Updates an existing user by ID.
  - **Request Body**:
    ```json
    {
      "name": "string",
      "email": "string"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "int",
      "name": "string",
      "email": "string",
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

  - **Example**:
    ```bash
    curl -X PUT -H "Content-Type: application/json" -d "{\"name\":\"Updated Name\", \"email\":\"updated@example.com\"}" http://localhost:1111/users/update/1
    ```

- **Delete a User**
  - **URL**: `/users/:id`
  - **Method**: `DELETE`
  - **Description**: Deletes an existing user by ID.
  - **Response**:
    ```json
    {
      "message": "User deleted successfully"
    }
    ```

  - **Example**:
    ```bash
    curl -X DELETE http://localhost:1111/users/1
    ```

### Error Handling

All endpoints will return appropriate HTTP status codes and error messages in case of failures. Here are some common error responses:

- **400 Bad Request**: The request data is invalid or malformed.
- **404 Not Found**: The requested resource was not found.
- **500 Internal Server Error**: An error occurred on the server.

### Models

#### Book

- **ID**: `int`
- **Title**: `string`
- **Author**: `string`
- **CreatedAt**: `timestamp`
- **UpdatedAt**: `timestamp`

#### User

- **ID**: `int`
- **Name**: `string`
- **Email**: `string`
- **Password**: `string` (hashed)
- **CreatedAt**: `timestamp`
- **UpdatedAt**: `timestamp`

### Example Requests

#### Create a Book

```bash
curl -X POST -H "Content-Type: application/json" -d "{\"title\":\"The Go Programming Language\", \"author\":\"Alan A. A. Donovan\"}" http://localhost:1111/books
```

#### Retrieve All Books

```bash
curl http://localhost:1111/books
```

#### Update a Book

```bash
curl -X PUT -H "Content-Type: application/json" -d "{\"title\":\"Updated Title\", \"author\":\"Updated Author\"}" http://localhost:1111/books/1
```

#### Delete a Book

```bash
curl -X DELETE http://localhost:1111/books/1
```

#### Create a User

```bash
curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"John Doe\", \"email\":\"john@example.com\", \"password\":\"password\"}" http://localhost:1111/users
```

#### Retrieve All Users

```bash
curl http://localhost:1111/users
```

#### Update a User

```bash
curl -X PUT -H "Content-Type: application/json" -d "{\"name\":\"Updated Name\", \"email\":\"updated@example.com\"}" http://localhost:1111/users/1
```

#### Delete a User

```bash
curl -X DELETE http://localhost:1111/users/1
```
