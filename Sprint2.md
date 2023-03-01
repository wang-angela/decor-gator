# Sprint 2

## API Documentation

### users.go:

#### Struct:
Creates a user struct that holds an id as an int and username, password, and email all as strings.

#### Functions:

##### getUsers(w http.ResponseWriter, r \*http.Request):
Returns all of the users contained in the database. Returns an error if there are no users contained in the database. Is stored as a GET function.

##### getUser(w http.ResponseWriter, r \*http.Request):
Returns a specific user contained in the database given by email. Returns an error if that user is not contained in the database. Is stored as a GET function.

##### createUser(w http.ResponseWriter, r \*http.Request):
Creates a specific user using the information stored in that JSON body. Returns an error if that user is already contained in the database. Is stored as a POST function.

##### updateUser(w http.ResponseWriter, r \*http.Request):
Updates a specific user using the information stored in that JSON body. Returns an error if that user is not contained in the database. Is stored as a PUT function.

##### deleteUser(w http.ResponseWriter, r \*http.Request):
Delete a specific user contained in the database given by email. Returns an error if that user is not contained in the database. Is stored as a DELETE function.

### posts.go:

#### Struct:
Creates a post struct that holds an id as an int and the user that posted the furniture, furniture type, and post title all as strings.

#### Functions:

##### getPosts(w http.ResponseWriter, r \*http.Request):
Returns all of the posts contained in the database. Returns an error if there are no posts contained in the database. Is stored as a GET function.

##### getPost(w http.ResponseWriter, r \*http.Request):
Returns a specific post contained in the database given by id. Returns an error if that post is not contained in the database. Is stored as a GET function.

##### createPost(w http.ResponseWriter, r \*http.Request):
Creates a specific post using the information stored in that JSON body. Returns an error if that post is already contained in the database. Is stored as a POST function.

##### updatePost(w http.ResponseWriter, r \*http.Request):
Updates a specific post using the information stored in that JSON body. Returns an error if that post is not contained in the database. Is stored as a PUT function.

##### deletePost(w http.ResponseWriter, r \*http.Request):
Deletes a specific post contained in the database given by id. Returns an error if that post is not contained in the database. Is stored as a DELETE function.

### images.go:

#### Struct:
Creates a image struct that holds an id as an int and the image byte data as a string.

#### Functions:

##### getImages(w http.ResponseWriter, r \*http.Request):
Returns all of the images contained in the database. Returns an error if there are no images contained in the database. Is stored as a GET function.

##### getPost(w http.ResponseWriter, r \*http.Request):
Returns a specific image contained in the database given by id. Returns an error if that image is not contained in the database. Is stored as a GET function.

##### createPost(w http.ResponseWriter, r \*http.Request):
Creates a specific image using the information stored in that JSON body. Returns an error if that image is already contained in the database. Is stored as a POST function.

##### updatePost(w http.ResponseWriter, r \*http.Request):
Updates a specific image using the information stored in that JSON body. Returns an error if that image is not contained in the database. Is stored as a PUT function.

##### deletePost(w http.ResponseWriter, r \*http.Request):
Deletes a specific image contained in the database given by id. Returns an error if that image is not contained in the database. Is stored as a DELETE function.

### handlers.go

#### Functions:

##### JWTCreateToken() (string, error)
Creates a JWT token that expires in one hour.

##### ValidateToken(next func(w http.ResponseWriter, r \*http.Request)) http.Handler
Makes sure that the correct token is created before displaying the home screen. Prints an error message otherwise.

##### GetJwt(w http.ResponseWriter, r \*http.Request)
Creates a JWT token if the correct access key is passed through the JSON headers.

##### Home(w http.ResponseWriter, r \*http.Request)
Prints a simple message to verify that the token is valid.

### password.go

#### Functions:

##### encrypt(password string) string
Encrypts a given password into encrypted hash data for user privacy.

##### comparePassword(password, hash string) boo
Compares a given password to its corresponding encrypted hash data.
