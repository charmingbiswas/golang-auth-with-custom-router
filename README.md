# golang-auth-with-custom-router
My implementation of authentication from scratch without using any library and custom routing.

This repo is pretty easy to understand.
### router.go
 - This file contains implementation of basic router struct which can be used to register rest api with different methods. It is basic but more features than vanilla golang mux.
 - Not much to do here. You can see the implementation if you want.

### utils.go
 - This file contains all the functions which were used for various purposes.
 - There is a function to create users. Here we are using a redis database.
 - There is a function ```generateToken()``` which generates a custom JWT for the specific user, without using any external library. It uses golang's in-built crypto package.
 - Then there is a custom function called ```hashPassword``` to generate password hash for user safety.

### main.go
 - All the routes are registered here.
 - Database connections are also initialized here.
 - It is setup in such a way that server will fail to start unless the database redis servers are up and running.
 - Here we are using two instances of redis database servers, one to store user data and one to maintain user _login session_

### Makefile
 -  Comes with a Makefile to start the server.
 -  All you have to do to run the server is ```make build```
 -  Then just do ```make up``` to run the containers and start up the servers.
 -  To stop do ```ctrl+c``` and to remove all containers and related services, use ```make down```
