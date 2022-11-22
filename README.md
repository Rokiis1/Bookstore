# Table of contents

- [Overview](#overview)
- [Technologies](#technologies)
- [API Routes](#api-routes)
- [Run Locally](#run-locally)
- [Demo](#demo)

# Overview

Create bookstore api where was used Fiber for to handle operations such as routing/endpoints, middleware, server request. And Gorm for ORM

# Technologies

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Render](https://img.shields.io/badge/Render-%46E3B7.svg?style=for-the-badge&logo=render&logoColor=white)


# API Routes

#### Get all books

```http
  GET /api/books
```

#### Get book

```http
  GET /api/get_books/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`| `string`| **Required**. id |

#### Create new book

```http
  POST /api/create_books
```
#### Delete book

```http
  DELETE api/delete_books/:id
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. id |


# Run Locally

Clone the project

```bash
  git clone https://github.com/Rokiis1/api-server-bookstore
```

Start the server

```bash
  go run main.go
```

# Demo

[API Link](https://api-server-bookstore.onrender.com)
