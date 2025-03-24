# Tweets API Changelog

## VERSION 0.4.1
_Dependencies Fixes._
<hr>

### Tweets
* The tweet service now depends on the follows service for followers list lookup(update)

<hr>


## VERSION 0.4.0
_Service Layer._
<hr>

### Tweets
* A new service layer was added to make the internal architecture more clear (update)
* The application logic errors are now defined on the domain layer (update)

### Follows
* A new service layer was added to make the internal architecture more clear (update)
* The application logic errors are now defined on the domain layer (update)

<hr>


## VERSION 0.3.1
_Documentation Update. Code Refactoring._
<hr>

### API Documentation
* Response model added to the new tweet endpoint documentation (update)

### Tweets
* The domain component was refactored to use a individual package for tweets (update)
* The repository component was refactored to use a individual package for tweets (update)
* The controller component was refactored to use a individual package for tweets (update)

### Follows
* The domain component was refactored to use a individual package for follows (update)
* The repository component was refactored to use a individual package for follows (update)
* The controller component was refactored to use a individual package for follows (update)

### Router
* The gin engine was set to run in production mode as default (update)

### General Changes
* Some tweaks and code revisions to use better practices (update)

<hr>

## VERSION 0.3.0
_Dockerization_
<hr>

### DOCKER
* Docker compose added for container execution (new)
* Dockerfiel added for container execution (new)

<hr>

## VERSION 0.2.0
_Documentation_
<hr>

### Database Documentation
* Collections documentation added (new)

### API Documentation
* API documentation added in swagger format (new)

### ROUTER
* Added CORS support on the gin router (update)

<hr>


## VERSION 0.1.0
_Initial Version_
<hr>

### Tweets
* The new Tweet API endpoint is now implemented (new)

### Follows
* The follow user API endpoint is now implemented (new)

<hr>

