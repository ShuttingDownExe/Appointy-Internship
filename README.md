# Appointy-Internship

This project was developed for the VIT Appointy Internship test. The project is complete to the best of my ability. 

## Problem Statement:

The project required me to *"Design and Develop an HTTP JSON API"* that is capable of the following:

1. Create User

2. Get an user id

3. Create a post

4. Get a post using id

5. list all posts of a user

## Implementation

I tried to implement this project in the following way:

1. Create a HTTP server to connect with the MongoDB cluster using GoLang

2. Create a login and an upload page that can be hosted on the HTTP server

3. In the login page there will be an option to sign up for first-time users

4. The login page will ensure the input from the user is not malicious i.e. removing special characters to prevent Cross-Site-Scripting attacks or SQL injection attacks.

5. The upload page is where users can upload pictures

5. While uploading, the file will be encrypted on the client side before uploading it to the DB. This can be accomplished by using the many Crpto Libraries in GoLang. 

6. The encrypted file will be stored on on the DB and accessed by the , and decrypted on the client. This prevents malicious actors from accessing the image directly from the DB

7.  User's passwords will be first checked for strength. This involves using special characters and alpha-numberic characters.

8. The passwords will be stored as a digest on the DB. 

9. Users will be able to see and edit their own posts and will only be able to see other users posts. (*Access Management*)
 
## Accomplishments:

In this project I was able to accomplish the following:

1. Create a MongoDB cluster (*Using MongoDB Atlas*)

2. Connect to the MongoDB cluster using GoLang

3. Create a basic login Page 

4. Create a basic upload page

5. Create a HTTP server using GoLang to host the login and upload pages

## Challenges Faced:

I have faced several challenges while working on this project. The biggest challenge being the time available. I am a complete beginner in GoLang and web development as I am only in 5th semester, and trying to implement the above ideas proved more challenging then initially thought. Given more time I will be able to complete the task successfully.
