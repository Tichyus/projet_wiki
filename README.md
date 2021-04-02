<p align="center"> 
<img
      alt="Flamin.go"
      src="https://i.imgur.com/0nvlLvW.png"
      width="400"
    />
</p>

Flamin.go
====
#### ***Golang & Json - Blog API***
HETIC - MT4 Golang project - Quentin MAILLARD, Corentin BOULANOUAR, Lucas LEHOT, Cyrille BANOVSKY

For more information on using Flamin.go, see the [Flamin.go Documentation](https://github.com/Tichyus/projet_wiki/wiki).

###### Toute ressemblance du logo Flamin.go avec un logo existant ou ayant existé serait purement fortuite (c'est faux).

# Project Status
*This repo is a [HETIC school](https://www.hetic.net/) project and its purpose is purely educational.* 

*Feel free to fork the project, but be aware that development might slow down or stop completely at any time, and that we are not looking for maintainers or owner.*

# Table of Contents
- [Flamin.go](#flamingo)
      - [***Golang & Json - Blog API***](#golang--json---blog-api)
- [Project Status](#project-status)
- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Project Demo](#project-demo)
- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Quick Start](#quick-start)
- [Deployment](#deployment)
- [Documentation](#documentation)
- [Known Issues](#known-issues)
- [Built With](#built-with)
- [Team Members](#team-members)
- [Acknowledgments](#acknowledgments)
- [License](#license)

# Overview
This project should be carried out according to the following guidelines:
* Back-end only
* Golang
* Json data
* MVC structure
* MySQL Database
* jwt auth
* GORM
* with a Readme and a Documentation

As defined by the teacher, the raw features intended are (at least):
* ability to connect
* create an account
* create an article
* edit an article
* comment an article

Other evaluation criterias are:
* clean API documentation
* clean git repo management

# Project Demo
*No demo available at the moment.*

# Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) section for notes on how to deploy the project on a live system.

## Requirements
Before installing, [download and install Golang](https://golang.org/dl/). go1.15.8 or higher is required.

## Installation
First, create the simplest local database
```
mysql -u root
create database flamingo
```

## Quick Start
```
go run .
```
Congratulations.

You can now access to all GET routes like list article(s), get comments, etc.

Now get to 127.0.0.1/signup to create an account. 

Then go to 127.0.0.1/signin to enter your credentials and get your bearer token (present in the header of the response under "Authorisation").

You can now access to all the POST routes during 20 minutes (create article, update article, create comment, etc) by adding the token in your header, until it expires.

# Documentation
See the [Documentation Wiki](https://github.com/Tichyus/projet_wiki/wiki) file for documentation of classes and utility functions.

# Known Issues
- The refresh token endpoint struggles with time calculations to give a new token, so for now, you just need a token that was created by the API to hit /refresh-token and get a new valid one.
- While you have a valid token you can access all POST routes, included the ones to update/delete comments, users, and articles that should need an access authorisation based on rights of the account to edit/update/delete an element.

# Built With
* [Golang](https://golang.org/) - Open source programming language
* [GORM](https://gorm.io/index.html) - "The fantastic ORM library for Golang"
* [Mux](https://github.com/gorilla/mux) - "A powerful HTTP router and URL matcher for building Go web servers"
* [Jwt-go](https://github.com/dgrijalva/jwt-go) - A go implementation of JSON Web Tokens

# Team Members
* **Quentin Maillard** - [Tichyus](https://github.com/Tichyus)
* **Lucas Lehot** - [lucaslehot](https://github.com/lucaslehot)
* **Cyrille Banovsky** - [Ban0vsky](https://github.com/Ban0vsky)
* **Corentin Boulanouar** - [Shawnuke](https://github.com/Shawnuke)

# Acknowledgments
* Aiming to build a clear and well-structured documentation, the [Stripe API Documentation](https://stripe.com/docs/api) has been a huge reference for us.
* [OAuth 2.0](https://www.oauth.com/), the modern standard for securing access to APIs guided us in the implementation of JWT.
* Hat tip to anyone whose code was used
* Inspiration
* etc

# License
This project is licensed under the terms of the [MIT](https://opensource.org/licenses/MIT) license.
