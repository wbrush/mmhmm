# mmhmm_test : notes_mgr
This service was written as a take home coding test for a position at mmhmm.

## Test Status
This test is mostly completed. I would have liked to spend more time on the unit tests and some time on the integration tests. This was built using a framework that I came up with a few years ago and have 
been adding to ever since. With this framework, I prefer to write unit tests to test individual functions that process data. Since there weren't really any of those here, there aren't many unit tests. I 
prefer to do testing of the REST endpoints with integration tests. This tests the full functionality of the endpoint including interacting with the DB. 

## Test Task
GOAL
We would like you to create the backend for a simple multi-user note taking application that allows each
user to create, read, update and delete plain text notes. Your work should include a RESTful API and a
persistence layer. You don’t have to build a front-end for the service.
The only constraint on the project is that you should use the Go programming language. Beyond that,
you may use any tools and frameworks that you wish. All of your work should, of course, be your own
original creation, but you’re free to use any sources you wish to research and learn.
To keep the project relatively simple, we have a couple of recommendations:
● Don’t worry about deployment as part of this project - you don’t need to show us how the
software runs on AWS or GCP. You should, however, plan to talk about how you’d deploy and
operate an API like this when we review your work.
● Feel free to simplify the persistence layer by using a simple in-memory database like
go-memdb. Again, plan to talk about what a production approach to persistence might look like.
You may work on this project on your own time. There is no time limit on the project, but we wouldn’t
expect it to take more than 4 hours. When you’re ready, we will set up a 45-minute call for you to walk
us through the code, explain your architecture, and discuss how we might modify it as requirements
change. Please share your code with us in advance of this call.

## Design tradeoffs

- I created a separate User DB since I figured that the system would probably require the user to login and this data would be stored there. However since this did not require the user to login, I've 
simplified the data structure to only the values I needed for this exercise.

- In an effort to track versions, this service allows the compiler to insert the git repo commit version and the build date/time in the executable. I usually build in unix so the command is:
go build -ldflags "-X main.commit=`git rev-parse --short HEAD` -X main.builtAt=`date +%FT%T%z`"

This will insert the current values in to the executable and the commit can be used as a version tracker. Unfortunately, this command works only in unix (or unix shell like Ming64) and I haven't found the correct syntax for windows yet.

# Development Notes

## Unit Testing
Some unit tests have been written for the packages in this service. 

## Integration Testing
Since the project stated not to worry about deployment, this was not done. However, the framework allows for running integration tests in a docker container with the DB running in another container as part of the deployment pipeline. These
tests is where I usually test REST endpoints that hit the DB since I prefer to know that the code is working with the DB and not using a mocked return value that I feel is the right data.

## Developer Notes
These are notes that I recommend providing to the next dev who might be working on this.

### Status
The service is mostly functional and runs locally. The following endpoints have been verified as working:
 - /api/info : returns system information that is useful in debugging some issues
 - /api/ping : returns nothing but is useful in determining connectivity
 - /api/help : returns the swagger page which documents the endpoints functionality
 - /api/v1/users : performs various operations on the user DB and records
 - /api/v1/notes : performs various operations on the note DB and records

## Required Packages
This service uses go-modules so the modules are documented in the go.mod file in the mainline. However, key modules are listed here for completeness.

### Swagger
Framework used for compiling the swagger comments into a swagger.json file.

go get -u github.com/go-swagger/go-swagger/cmd/swagger

### Gorilla
Framework used for handling the REST endpoints and server functions.

go get -u github.com/gorilla/mux
go get -u github.com/gorilla/handlers

### logrus
Framework for logging.

go get -u github.com/sirupsen/logrus

# Deployment
Need to document (and/or update) the deployment process here

## Process
This would spell out the deployment steps once developed and tested.

## Running locally
To run this locally, you would need to have the go compiler installed. I can add the windows executable to the repo if needed. To run locally:
1. run the command "go get github.com/wbrush/mmhmm" from the command line
2. if on unix, run " go build -ldflags "-X main.commit=`git rev-parse --short HEAD` -X main.builtAt=`date +%FT%T%z`" "
3. if on windows, run "go build"
4. if on unix, run "./mmhmm" 
5. if on windows, run "mmhmm.exe"
6. you should see startup messages similar to:

time="2021-08-15 17:36:26.349242" level=info msg="Log level was set to debug"
time="2021-08-15 17:36:26.349763" level=info msg=------------------------------
time="2021-08-15 17:36:26.349763" level=info msg="Starting notes-mgr"
time="2021-08-15 17:36:26.349763" level=info msg="Version:4bda05c; Build Date:2021-08-15T17:35:52-0500"
time="2021-08-15 17:36:26.349763" level=info msg=------------------------------
time="2021-08-15 17:36:26.743940" level=debug msg="DB version is 2 for shard []\n"
pg: 2021/08/15 17:36:26 table.go:1040: DEPRECATED: use pg:"..." struct field tag instead of sql:"..."
time="2021-08-15 17:36:26.783445" level=info msg="Starting module HTTP REST API"
time="2021-08-15 17:36:26.783972" level=info msg="Starting REST Server on port 8000..."

7. The API docemented above should be available at "localhost:8000"
