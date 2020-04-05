# crud-posts

A simple CRUD application built with Golang + VueJS + SQLite. 

This project is made of 2 separate Docker containers that holds:


 - Golang backend (Gorilla Mux)
 - VueJS frontend (Nuxt.js)

The entry point for a user is a website which is available under the address: http://localhost:8080
## Features
 - Authentication with JWT
 - List posts
 - See post details
 - Delete post (only can be deleted by author)
 - Create post
 
![alt text](https://i.hizliresim.com/qmtUxK.png "Login")
![alt text](https://i.hizliresim.com/GAP6nQ.png "Dashboard")
![alt text](https://i.hizliresim.com/DW32Mb.png "Post details")




### Prerequisites

In order to run this application you need to install two tools: **Docker** & **Docker Compose**.

### How to run it?

An entire application can be ran with a single command in a terminal:

```
$ docker-compose up -d
```

If you want to stop it use following command:

```
$ docker-compose down
```

Rest API runs on http://localhost:8000
Frontend app runs on http://localhost:8080

### Demo Credentials

 - Username: ``mustafayumurtaci``
 - Password: ``123qwe``


#### [crud-posts-api (Rest API)](https://github.com/mstfymrtc/go-posts-api)

This is a Gorilla Mux (Go) based application that connects with a
database that and expose the REST endpoints that can be consumed by
frontend. It supports multiple HTTP REST methods like GET, POST, PUT and
DELETE for two resources - users & posts.
Authentication is provided by JWT and SQLite used as database.


#### [crud-posts-app (Frontend)](https://github.com/mstfymrtc/vue-posts-app)

This one is a Nuxt.js application that consumes Rest API above.
### Todos


 - Add Swagger UI for API
 - Migrate from SQLite to another database, because gorm (ORM) does not support foreign keys with auto migration
 - Add auth middleware for VueJS application
 - Show logout button only if user logged in
 - Only post author can delete the post (backend only)


## License

Distributed under the Apache License 2.0. See ``LICENSE`` for more information.

