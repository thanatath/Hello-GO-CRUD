# Simple Golang CRUD with Gorm

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)

## About <a name = "about"></a>

Simple Go lang with CRUD Backend API

## Getting Started <a name = "getting_started"></a>


### Prerequisites

What things you need to install the software and how to install them.

```
git clone https://github.com/thanatath/Hello-GO-CRUD.git
```

.env file

```
DB_HOST
DB_NAME
DB_USER
DB_PORT
DB_PASSWORD
```

Table `users` in Database
````
name
lastname
age
````


### Installing

A step by step series of examples that tell you how to get a development env running.

```
go run .\main.go
```

 

## API <a name = "API"></a>

- /
- /create_user
- /get_user
- /update_user/:id
- /del_user/:id