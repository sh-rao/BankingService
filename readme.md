# Banking Application
# Prelude
A simple banking application called **_BankingApp_** written in Go(https://golang.org/) with the following functionality:
- Create a customer, and an associated account
- Deposit, withdraw and check balance for customers
- Check bank's balance

# Design
**_Banking_** domain entities:
* Customer - represents a person/organisation that holds an Account, has-a relationship in the form of composition.  
* Account - represents the money(asset) owned by a customer
NOTE: From robustness perspective, customer has reference to Account and vice versa. This is intentional to ensure interfaces on
Account cannot be used without a `customerId`.

**_Banking_** domain services:
* CustomerService is responsible for
  * Creating customers
  * Validating customers
  * Exposing the following interfaces for managing account via Deposit, Withdraw & CurrentBalance by customerId
* AccountService is responsible for
  * Creating an account for a given customer identified by customerId
  * Validating customer against the account
  * Validating amount value to be positive for Deposit and Withdraw and to be lesser than or equal to the balance for Withdraw
  * Updating the bank balance via the injected BankingService for any successful Deposit or Withdraw operations
* BankService is responsible for
  * Responsible for maintaining bank's balance
  * Exposing interfaces to update the bank's balance and retrieve the bank's balance

With high cohesion and low coupling design principle, using DI(Dependency Injection technique)
  - `bankService` is injected to `accountService` to update balance.
  - `accountService` is injected to `customerService` to expose Deposit, Withdraw and CurrentBalance operations via `customerService`
    with the intention that every operation has to be initiated with `customer` as the primary context

# Implementation Notes
- There are certain idiosyncrasies of Go language and hence variable names or code style will look and feel
  different from other languages. A quick read of this - https://golang.org/doc - may assist reading of the code.
- I am not a big believer is adding comments in the source code, the reason being source code itself should be
  self-explanatory. You won't see any comments in the source code unless there is a very good reason. :-)

# Assumptions
* For simplicity purposes amount has been restricted to two decimal places.
* No transaction domain entity has been considered as there is no requirement for persistence
* As this is not implemented as a web server / API service, `main.go` sets up a few customers and exercises
  all operations exposed by `customerService`. So this essentially acts as integration tests and hence there is no
  explicit integration tests have been written.
  
### What could have been done better
I hate to use time as an excuse but given an extra few hours or so, I would improvise the current design and implementation with the following:

- The whole application could have been designed as RESTful API service(s), with the following resources
  - `customers` accessed via /bank/customers/{id}
  - `accounts` accessed via /bank/customers/{id}/accounts/{id}
  - `bank` as the root resource for accessing and updating the balance
   - GET /bank?field=balance to retrieve **balance**
   - PUT /bank with body containing the **amount** to update the **balance**
- Completely decouple `accountService` and `bankService` by using asynchronous event-driven architecture/design.
- Logging and error handling can be improved by passing context and using context logger where ever necessary.

# Prerequisites
- Make sure you have installed the latest version of Golang from https://golang.org/
  This service has been built and tested with go1.13 darwin/amd64 (on MacOS Mojave v10.14.6)
  in GoLand 2020.1 IDE.

# How to run the service
- Extract the archive into your local target folder.
  e.g.
  ~~~
  unzip BankingApp.zip -d target-folder
  ~~~
  
- From the project root folder (e.g. BankingApp), run this command to download all the dependencies
  ~~~
  go get -u ./...
  ~~~
  
- Run the service from the project root folder (e.g. BankingApp) by running the following command
  ~~~
  go run main.go
  ~~~
  
 - Output should look similar to this
   ~~~
   Customer 8b8e3b59-67b5-4573-b267-2e2e3a2bef71 balance: 5.00 
   Customer dfa387e1-fb0e-4824-b818-e6dd0a959263 balance: 5.00 
   Customer bb59c52a-6010-4097-a087-2da41b6b9ece balance: 5.00 
   Customer b09938d4-728b-4649-a38c-2e592f8093a6 balance: 0.00 
   Customer ee91e3a3-ecd9-4065-acf1-6d4591655978 balance: 20.00 
   Bank balance: 35.00
   ~~~
  
  # Running unit tests
  Units tests can be run by the following command from the project's root folder (e.g. BankingApp)
  ~~~
  go test ./... -v
  ~~~
  
