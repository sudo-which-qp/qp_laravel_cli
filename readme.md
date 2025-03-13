## QP Laravel CLI

<p>This is just a CLI that clones my Laravel template for me to use when trying to build any laravel app. I was struggling with always trying to scaffold a laravel project setting up all I need in the project which mostly eats up my productive time</p>

<p>So I decided to build this CLI using Go Lang to always clone the repo that I have already setup and made public. It clones it and run the laravel pub get command to make sure the project is just ready for me to start building apps or anything with laravel instead of trying to setup my folder structure.</p>

<h4>To run this if you want to use this code for your laravel template:</h4>
<h6 style="color:red">Make sure Makefile can run on your PC</h6>

## Simple steps

1. make build
2. make move (running this the first name will require sudo privileges)
3. qp_laravel create project-name

## Any issues?

1. make clean-build

<p>You can change qp_laravel to what you want, the second step is for moving the build into your local bin folder, this is for Mac OS as of the moment</p>

<p>The last step there is for creating the laravel app with my template which is structure to how I build API's with laravel, you can clone and structure to how you want yours to be or just Clone the Go project and edit the template repo to yours, that's just it</p>
