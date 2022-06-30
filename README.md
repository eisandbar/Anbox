# Anbox
A demo api for anbox


## Introduction
To get a sense of how you would approach this kind of work, we’ve come with the following
simple scenario:
As a customer using Anbox Cloud I would like to implement a service to manage a set of
games and their metadata. The service should provide a REST API, store its data in a databaseand have a CLI client to interact with it.
With it, as the user I should be able to:
    - Add and manage (update, delete) games into the system, including some basic
    information about these games (title, description, age rating, publisher, genre,
    published date, logo, ..)
    - List all games and filter lists by publisher, genre, published date, ...
    - Add and manage (update, delete) users into the system, including some basic
    information (username, age, ...)
    - List all users and filter lists by age, ...
    - Allow linking games to users including additional information of how long a particular
    user has played a game (in hours)

### Step 1: Describe the expected user experience
This should include all CLI commands and their options, think of this as writing the help or
manual page of your software ahead of time, including documented examples of user
interaction.

### Step2: Describe the expected REST API
This should include the documentation of the REST API, all supported methods and what
they’ll be doing, examples of input and output data and what query parameters will be
supported and what they’ll do.

### Step 3: Describe the database structure
This should describe the database structure you have chosen to efficiently store the data the
service needs to handle.

### Step 4: Implementation
Now that you have a good idea of what everything will look like, implement a couple of the
APIs in a minimal standalone Go or Python based REST API service. It is more important to
implement a few things very well than to poorly implement the entire API. Please make sure
to add proper input validation and error handling.
We’d recommend you do this in a VCS like git, structuring your commits as you would when
contributing to a real project (don’t hesitate to rebase into fewer logical commits) and then
send us the result as a tarball or ZIP archive.
Please provide a README.md file as part of the implementation which describes how the
service can be built and set up.