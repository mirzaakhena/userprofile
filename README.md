# UserProfile

This project helped by tools https://github.com/mirzaakhena/gogen for generate some boiler platecode


## Usecase and Requirement
Those are only "fiction usecase" added by me and it may not cover all the negative case

### RegisterUser
* Email must be unique
* All field must not empty
* Activation User Token is created (may improve by using cache)
* Password is hashed
* Email is send to user for activation

### Activation
* activate user by sending the activation token

### Login
* Input Email and Password
* Return token (possibly jwt token)

### ShowAllUser
* No need input
* Will return all user with ID and Email information

### ShowUser
* Need UserID as an Input
* Will return specific user with ID, Email, and Address information

### UpdateUser
* Need UserID and Address as an Input
* Will update the address only


## How to run

To run with memory as database (under gateway inmemory) you can use this command
```
$ go run main.go simplememory
```

To run with real database (SQLite) (under gateway indatabase) you can use this command
```
$ go run main.go usedatabase
```

