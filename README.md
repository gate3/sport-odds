# Bet.Me Developer Task

## Introduction
The Bet.Me task is to build an app that reads data from a sports odds API, parses the data and persists it in a data store in an efficient manner.

## File Structure

The following details the file structure for the project:
- pkg : This contains the main files for the project and contains code that can be re-used elsewhere.
    - common: The common package contains shared code between the packages in the project such as 3rd party invocations
    - config: The config directory contains configuration code for the project such as environment variables or language files
    - infrastructure: Initialization of any infra that is required such as a database, redis-cache or elk stack
        - db: This contains the database initialization code as well as repository pattern functions for the different entities
    - services: This contains the business logic of the project and the services are responsible for connecting the required parts to get the tasks done
    
## Libraries Used

- Testify - Mocking Library
- Viper - Environment Variables
- Copier - Deep copy structs
- MongoDb Driver - Driver for working with mongodb database

## Considerations

- One of my major considerations was to focus on constantly pushing working code. This is evident in the fact that all pushes I made contained non buggy and working code. This pattern is to ensure that shippable code is pushed consistently.
- Something to also note is that i included a sample app.env file in the project to show what environment variables are needed to run the project
- I adhered by seperation of concerns by ensuring my packages are responsible for specific parts of the task

## Todo 

I had a lot of fun building this project, but there are some improvements i can still make

- More Tests: I can still write tests for the bookmaker api functions and some other parts of the code. I could also add more edge cases for the tests written for the services as well. I could also have written some integration tests.
- Mock Database: Another improvement that will make the repository functions testable would be to create a mock for the mongodb Database functions.
- Change the name of the database model fields so its not tied 1 to 1 to the api fields
- Make use of a task management tool like trello to show how I plan out my tasks
- Containerize the app by adding a docker container file

## Build Instructions

- Clone the repository using git clone https://github.com/gate3/sport-odds
- Rename the app.env.sample file to app.env
- Add the required values to the newly renamed app.env file
- Ensure all dependencies like the database are running
- 

Cheers


