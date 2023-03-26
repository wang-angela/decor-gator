# Sprint 3

## Backend Accomplishments
1. Created methods that send welcome email to user when they sign up.

## Backend Shortcomings
1. Cannot send emails that aren't manually approved by Amazon Web Service.

## API Documentation (Continued from Sprint 2.md)

### email.go

#### Functions

##### SendWelcomeEmail(destinationEmails []string)

Sends an email from decorgators@gmail.com confirming that the user signed up with our service. For now, the only emails we can send are to those we manually approve on our Amazon Web Service account. We cannot fix this until we get approval to leave the sandbox from Amazon.

##### SendForgotPasswordEmail(destinationEmails []string)

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


