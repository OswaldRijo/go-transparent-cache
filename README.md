# Golang Challenge Cache
Project that implements a transparent cache using concurrency/parallelism

## Getting started 

The directory is as follows:

    .
    ├── controller                      # Part of example server
        └── controller.go               # Part of example server
    ├── pkg                             
        ├── cache.go                    # Solution
        └── cache_test.go               # Solution's tests
    ├── response                        # Part of example server
        └── builder.go
    ├── server                          # Part of example server
        └── router.go
    ├── service                         # Part of example server
        └──price.go
    ├── view                            # Part of example server
        ├── serializer.go
        └── view.go
    ├── Server.postman_collection.json  # Part of example server
    ├── Dockerfile                      # Part of example server
    ├── .gitignore
    ├── main.go                         # Part of example server
    ├── README.md 
    └── script.sh                       # Installation Script

As you may see, the solution is located in pkg directory and there is a server app developed to implement the solution 
as an example. 

The solution presented here was developed in golang and has involved concepts such as concurrency and parallelism to 
optimize some expensive process in matter of time. As an example, I decide to build a pretty simple server just to show you
how the cache developed will behave, and why did I decide that? just for fun! why not right?

Further in the text, is explained how to run the tests case and how to deploy the example app to see the cache library in action


### Pre-requisites

You will need Golang be installed, and for best experience upper 1.5 version, 
because in previous versions the default value for GOMAXPROCS is set to 1, which means that by default "go routines" will 
be concurrent and not parallel.


## Run the solution Tests⚙️

Open a terminal and navigate to the pkg folder with command:
```
cd pkg
```
After that, run in the terminal:
```
go test -v
```

The command above will print all the results of every test in the tests case


### Installing and Deploying Example Server App (Ubuntu)

In root folder directory is a file named as "script.sh" that is intended to run in Ubuntu machine. This script will set up everything that server app needs. 
For running the script, first it's needed to set the script as executable, hence open a terminal and runs the command 
in project root folder.
```
chmod +x ./script.sh
```
then to execute the installation script run the following command:
```
sudo ./script.sh
```
finally, wait for the script finishes and voilá! the server app will be running on port 8080 (make sure you have this port 
available)



##Installation and Deploying Example Server App (NO Ubuntu dist)

The "main.go" file located in root directory is in charge of starting the server, the next command will build the app, so
open a terminal and execute:
```
go build .
```
After that, execute the file that the command above generates in root directory called "Golang-challenge" as follows:

```
./Golang-challenge
```

and that's it!



##Collection to make Requests to the server

Import "Server.postman_collection.json" file in postman to get ready and start to make requests to the server.

## Third party

* [Mux](https://github.com/gorilla/mux/) - Library to build http servers