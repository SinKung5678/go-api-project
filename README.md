# Golang API

This repository contains the Postman Collection for the Golang API project.

## API Endpoints

### Get Users

- **URL:** `/users`
- **Method:** `GET`
- **Parameters:**
  - `page`: Page number (default: 1)
  - `limit`: Number of users per page (default: 10)
- **Response:**
  ```json
  [
      {
          "first_name": "John",
          "last_name": "Doe",
          "email": "john@example.com",
          "age": 25,
          "address": "123 Main St",
          "gender": "Male",
          "phone": "555-1234"
      },
      {
          "first_name": "Jane",
          "last_name": "Smith",
          "email": "jane@example.com",
          "age": 30,
          "address": "456 Elm St",
          "gender": "Female",
          "phone": "555-5678"
      }
      // Add more user data as needed
  ]

