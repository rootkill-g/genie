# genie 🧞

A Web application to cater user management.

Provides public APIs for:
1. User registration
2. Fetch list of all registered users
3. Fetch information of a specific user
4. Update User info
5. Delete a specific User
6. Post creation
7. Post retrieval

## Tech stack 🤖

* Go
* Gin Web Framework
* Gorm
* Gorm's SQLite Driver

## Get Started 🚀

To run the API server, you first have to install the go dependencies:
```bash
go get github.com/gin-gonic/gin
go get gorm.io/driver/sqlite
go get gorm.io/gorm
```

After the above step, run:
```bash
go run main.go
```
The Gin server will start running on port `8080`

## Register Users 📓

To register users, we can use cURL program. Let's register three users:
```bash
curl -X POST -H 'Content-Type: application/json' -d '{"name":"Elliot Alderson", "email":"elliot@fcociety.com"}' http://localhost:8080/users/
curl -X POST -H 'Content-Type: application/json' -d '{"name":"Tyrell Wellick", "email":"tyre@fsociety.com"}' http://localhost:8080/users/
curl -X POST -H 'Content-Type: application/json' -d '{"name":"Angela Moss", "email":"angela@fsociety.com"}' http://localhost:8080/users/
```

## Get all users 👥

To get list of all registered users:
```bash
curl http://localhost:8080/users/
```

## Get specific user 👤

To get a specific user's information, we can pass the user ID along in the URI:
```bash
curl http://localhost:8080/users/2
```

## Update user information 📓

To update any user's information, we can use the `PUT` verb in the curl for users URI with user ID, and pass along the updated data:
```bash
curl -X PUT -H 'Content-Type: application/json' -d '{"name":"Tyrell Wellick", "email":"tyrell@god.com"}' http://localhost:8080/users/2
```

## Delete a user record ❌

To delete a user record from the database, use the `DELETE` verb and pass the user ID in the URI:
```bash
curl -X DELETE http://localhost:8080/users/2
```

## Create a New Post 📮

To create a new Post by a specific user, we pass the user ID in the data along with the post's title and body:
```bash
curl -X POST -H 'Content-Type: application/json' -d '{"title":"Keeping up with Angela","body":"This is a new newsletter that Im starting to post updates about my life. PS: Angela", "user_id":3}' http://localhost:8080/posts/
```

## Read Posts 📮

Want to read the posts created by the users? Here's how you can do so:
```bash
curl http://localhost:8080/posts/
```

---
Enjoy the Genie APIs
