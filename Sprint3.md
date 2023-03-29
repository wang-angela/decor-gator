# Sprint 3

## Backend Accomplishments
1. Was able to start sending emails through AWS SES service.
2. Made unit tests to assure that these emails are sent successfully.

## Backend Shortcomings
1. Cannot send emails that aren't manually approved by Amazon Web Service.
2. Have not gotten AWS to host images through buckets yet.
3. Needs to be able to send JWT token to frontend.

## Updated API Documentation

### users.go:

#### Struct:
Creates a user struct that holds an id as an int and username, password, and email all as strings.

#### Functions:

##### getUsers(w http.ResponseWriter, r \*http.Request):
Returns all of the users contained in the database. Returns an error if there are no users contained in the database. Is stored as a GET function.

##### getUser(w http.ResponseWriter, r \*http.Request):
Returns a specific user contained in the database given by email. Returns an error if that user is not contained in the database. Is stored as a GET function.

##### createUser(w http.ResponseWriter, r \*http.Request):
Creates a specific user using the information stored in that JSON body. Returns an error if that user is already contained in the database. Is stored as a POST function. Also sends an email to the user confirming account creation (see email.go for SendWelcomeEmail() function).

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

### email.go:

#### Functions:

##### SendWelcomeEmail(destinationEmails []string):

Sends an email from decorgators@gmail.com confirming that the user signed up with our service. For now, the only emails we can send are to those we manually approve on our Amazon Web Service account. We cannot fix this until we get approval to leave the sandbox from Amazon.

##### SendForgotPasswordEmail(destinationEmails []string):

Sends an email from decorgators@gmail.com for user to reset their password. Similar issues from SendWelcomeEmail().

## Frontend Goals
Chris: Username --> Uploading image --> 
Joanne: Use the GUI create About page

### Transform Website using MUI
- Make front page (About page)
- Login, Sign up, Post, Buy

### Implement username for users
- Include in Sign-up page
   - Check if the username is unqiue
- Change Log in page to login with username instead of email
- Send username instead of email when creating new post

### Finish making Buy Page & Post Page structure and make Rent Page
#### Post Page Features:
- Title, Description, Type (Drop-down Menu), Buy/Rent, (Multiple) Images, Price, Location (?)

#### Buy Page Features: 
- Display Posts (obvious) by Newst (default)
- Sort by Chair, Sofa, Bed, Table, Electronics, Lighting, Stroage, Kitchen, Other 
- Search bar to search posts by Buy/Rent, Title, Type, Price (Low to high, high to low), Location (?)
   - Three fields - Search by input string, search by drop-down with Buy/Rent, Price, Newest/Oldest, and search by drop-down furniture type
- When clicked on a post that same user made - Put edit button
- Implement Edit page using Post Page
- User Icon -> User Page

#### In User Page:
- Show my post
- Change username, email, password, first and last name


