# Sprint 4

## Backend Accomplishments

1. Moved our data into MongoDB so that the database is stored remotely.
2. Create AWS bucket to host post data.
3. Modified the JWT token to allow the frontend to access the token to authenticate users.

## Backend Shortcomings

1. The post data is stored properly in an AWS bucket, but is not accessible for edits, so it is simply used as a second database in addition to MongoDB.
2. Our project was not able to deploy publicly as a website, so only emails that are registered with our AWS account will recieve emails from DecorGator.

## Frontend Accomplishments

1. Completed the post upload page.
2. Handled image uploading and rendering using base64.
3. Built pop-up windows for each post to display all of the post's information.
4. The search function now can filter by furniture type.
5. "My Page" feature that allows the user to change their email and password.

## Frontend Shortcomings

1. JWT token/authentication was unable to be implemented.
2. Navbar was unable to be completed.

## Updated API Documentation

### user.go:

#### Struct:
Creates a user struct that holds an id as an object ID and first name, last name, username, password, and email all as strings.

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

### post.go:

#### Struct:
Creates a post struct that holds an id as an object ID, the price as a float, and the user that posted the furniture, furniture type, image URL, and post title all as strings.

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

### jwt.go

#### Struct
Creates JWT struct that stores token as a string value and create exception struct that helps display error messages.

#### Functions:

##### CreateTokenEndpoint(w http.ResponseWrite, r \*http.Request)
Validates user. If user exists and correct password is entered, then a token string is created and return through a json file.

##### ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc
Validates JWT and allows access if token is valid.

### accessors.go

#### Functions:

##### JwtVerifyUserExists(user models.User) bool
Verifies that user exists. Used by CreateTokenEndpoint() function.

##### JwtVerifyPassword(user models.User) bool
Verifies that password is correct for a specific user in the database. Used by CreateTokenEndpoint() function.

### password.go

#### Functions:

##### encrypt(password string) string
Encrypts a given password into encrypted hash data for user privacy.

##### comparePassword(password, hash string) boo
Compares a given password to its corresponding encrypted hash data.

### email.go:

#### Functions:

##### SendWelcomeEmail(destinationEmails []string):

Sends an email from decorgators@gmail.com confirming that the user signed up with our service. For now, the only emails we can send are to those we manually approve on our Amazon Web Service account. We cannot fix this until we get approval to leave the sandbox from Amazon.

##### HelperForgotPassword(w http.ResponseWriter, r \*http.Request):

Wrapper for SendForgotPasswordEmail(), so no parameters are needed.

##### SendForgotPasswordEmail(destinationEmails []string):

Sends an email from decorgators@gmail.com for user to reset their password. Similar issues from SendWelcomeEmail(). This is stored as a PUT function.

### buckets.go:

#### Functions:

##### InitAWSSession():

Creates an AWS session with the proper location and credentials.

##### CreateBucket() (resp *s3.CreateBucketOutput)

Creates a bucket using the S3 service in AWS that we can store data to.

##### UploadObject(post models.Post) (resp \*s3.PutObjectOutput)

Sends a post object to the bucket in AWS.

##### GetObject(id string, post models.Post) error

Returns a post function given by its id.

##### DeleteObject(key string) (resp \*s3.DeleteObjectOutput)

Deletes a post given by the post id.

### users_test.go

#### Functions:

##### TestGetAllUsers (t \*testing.T)
Test getUsers() using a custom http request and checks results.

##### TestGetUser (t \*testing.T)
Test getUser() using a custom http request and checks results.

##### TestCreateUser (t \*testing.T)
Test createUser() using a http request and checks results. The changes to the database are reverted.

##### TestDeleteUser (t \*testing.T)
Test deleteUser() using a http request and checks results. The changes to the database are reverted.

### posts_test.go

#### Functions:

##### TestGetAllPosts (t \*testing.T)
Test getPosts() using a custom http request and checks results.

##### TestCreatePost (t \*testing.T)
Test createPost() using a http request and checks results. The changes to the database are reverted.

### password_test.go

#### Functions:

##### TestEncryption (t \*testing.T)
Tests encrypt() then checks result using comparePassword().

### email_test.go

#### Functions:

##### TestSendWelcomeEmail (t \*testing.T)
Tests SendWelcomeEmail() and checks result to ensure email was sent.

##### TestSendForgotPasswordEmail(t \*testing.T)
Tests SendForgotPasswordEmail() and checks result to ensure email was sent.

### Frontend Unit Tests (In Sprint 4)
*Note: Some unit tests from sprint 3 had to have modifications to reflect new changes to post creation system*
1. Displays details of a post when clicked on (the pop-up window)
     (including title, furniture type, price, post description and image)
3. Searches posts by category and/or title
     (Is able to search with category only or title only or both category and title)
4. Changes current email or password
