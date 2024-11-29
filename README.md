# Go Products CRUD Application

A web-based CRUD (Create, Read, Update, Delete) application crafted with Go and MySQL. This project demonstrates a practical implementation of database operations through a clean web interface, focusing on product and category management.

Developped by
*Radian Try Darmawan - 5025221097*
*Albin Maurice Pierre Tardivel - 5999241037*

## Features

This application brings together essential web development concepts and database operations. At its core, it provides comprehensive product management capabilities, allowing users to create, view, modify and remove product entries. The category system enables logical organization of products, while the MySQL integration ensures reliable data persistence. The web interface presents these features through an intuitive layout, following RESTful architectural principles.

## Prerequisites

Before diving into this project, ensure you have Go 1.x installed on your system. MySQL 8.0 or higher is required for the database functionality. Familiarity with Go programming and basic SQL knowledge will be beneficial for understanding and modifying the codebase.

## Installation

The installation process consists of four main steps:

First, clone the repository to your local machine:
```bash
git clone https://github.com/radiandrmwn/go-crud-web.git
```

Then, open `config/database.go` and configure your database connection settings according to your local MySQL setup.

Finally, launch the application:
```bash
go run .
```

## Project Structure

The project follows a clean and modular architecture:

```
go-crud-web/
├── config/
│   └── database.go
├── controllers/
│   ├── categorycontroller/
│   ├── homecontroller/
│   └── productcontroller/
├── entities/
│   ├── category.go
│   └── product.go
├── models/
│   ├── categorymodel/
│   └── productmodel/
├── views/
│   ├── category/
│   ├── home/
│   └── product/
└── main.go
```
