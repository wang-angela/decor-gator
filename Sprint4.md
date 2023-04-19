# Sprint 4

## Backend Accomplishments

1. Moved our data into MongoDB so that the database is stored remotely.
2. Create AWS bucket to host post data.

## Backend Shortcomings

1. The post data is stored properly in an AWS bucket, but is not accessible for edits, so it is simply used as a second database in addition to MongoDB.
2. Our project was not able to deploy publicly as a website, so only emails that are registered with our AWS account will recieve emails from DecorGator.

## Frontend Accomplishments

## Frontend Shortcomings

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
Tests encrypt() then checks result using comparePassword().

### email_test.go

#### Functions:

##### TestSendWelcomeEmail (t \*testing.T)
Tests SendWelcomeEmail() and checks result to ensure email was sent.

##### TestSendForgotPasswordEmail(t \*testing.T)
Tests SendForgotPasswordEmail() and checks result to ensure email was sent.
