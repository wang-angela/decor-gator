# Sprint 2

## What Frontend accomplished in Sprint 2:
1. Connected frontend and backend to communicate with fetch
2. Created a Buy Page where users can see posts created. 
3. Created Post Page where users can submit posts.
4. Implemented unit tests for nearly all functions.

## What Frontend didn't get accomplished:
1. Constructing the Buy Page and adding more features to it (ex. display post in a box)
2. Upload images for creating posts


## What backend accomplished in Sprint 2:
1. Created a file called handlers.go that creates a token that is required to access a given page. 
2. Made a file called images.go that can store image byte data in the database.
3. Implented cors and allowed frontend to communicate with backend
4. Implemented unit tests for nearly all functions

## What backend didn't get accomplished:
1. Not able to get the user token to work in a web broswer nor be able to set a cookie after logging into an account.


## Unit Tests

### Frontend:
#### Signup and Login
* Flips login screen to signup page
* Returns error message if no sign up information entered
* Signs up with all information entered
* Alerts if sign up is attempted twice with the same email
* Alerts if login information is invalid (Wrong password)
* Logs in with correct information and redirects to Buy Page

#### Creating Post
* Click "+ Post" button and submit a post with title and furniture type
* Click "Show Post" to log post title, furniture type, and user email to console

### Backend:

#### user_test.go
* TestGetAllUsers()
* TestGetUser()
* TestCreateUser()
* TestDeleteUser()

#### posts_test.go
* TestGetAllPosts()
* TestGetPost()
* TestCreatePost()
* TestDeletePost()

#### images_test.go
* TestGetAllImages()
* TestGetImage()
* TestCreateImage()
* TestDeleteImage()

#### password_test.go
* TestEncryption()

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

### users_test.go

#### Functions:

##### initDB()
Initializes database and creates gorm transactions.

##### TestGetAllUsers (t \*testing.T)
Test getUsers() using a custom http request and checks results.

##### TestGetUser (t \*testing.T)
Test getUser() using a custom http request and checks results.

##### TestCreateUser (t \*testing.T)
Test createUser() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

##### TestUpdateUser (t \*testing.T)
Test updateUser() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

##### TestDeleteUser (t \*testing.T)
Test deleteUser() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

### posts_test.go

#### Functions:

##### TestGetAllPosts (t \*testing.T)
Test getPosts() using a custom http request and checks results.

##### TestGetPost (t \*testing.T)
Test getPost() using a custom http request and checks results.

##### TestCreatePost (t \*testing.T)
Test createPost() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

##### TestUpdatePost (t \*testing.T)
Test updatePost() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

##### TestDeletePost (t \*testing.T)
Test deletePost() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

### images_test.go

#### Functions:

##### TestGetAllImages (t \*testing.T)
Test getImages() using a custom http request and checks results.

##### TestGetImage (t \*testing.T)
Test getImage() using a custom http request and checks results.

##### TestCreateImage (t \*testing.T)
Test createImage() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

##### TestUpdateImage (t \*testing.T)
Test updateImage() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

##### TestDeleteImage (t \*testing.T)
Test deleteImage() using a http request and checks results. The changes to the database are undone through a rollback using gorm transactions.

### password_test.go

#### Functions:

##### TestEncryption (t \*testing.T)
Tests encrypt() then checks result using comparePassword() .
