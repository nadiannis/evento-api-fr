<div align="center">
  <br>
  <h1>Evento</h1>
  <p>🎫 A simple REST API for ordering event tickets online 🎫</p>
  <br>
</div>

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Entities](#entities)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
- [Tech Stack](#tech-stack)
- [Run Locally](#run-locally)
- [Screenshots](#screenshots)

## Description

[`^ back to top ^`](#table-of-contents)

**Evento** is a simple REST API for ordering event tickets online. This API is written in Go. It is created as a submission for the third-week exam in the Go phase of the Backend Development Training.

## Features

[`^ back to top ^`](#table-of-contents)

- View list of customers & their orders.
- Add a new customer.
- View a customer.
- Change balance amount.
- View list of events with the tickets available.
- View an event with the tickets available.
- View list of tickets.
- View a ticket.
- View list of orders.
- Order a ticket.

## Entities

[`^ back to top ^`](#table-of-contents)

There are 5 entities: **Customer**, **TicketType**, **Event**, **Ticket**, & **Order**.

**Customer**

- id: `int64`
- username: `string`
- balance: `float64`

**TicketType**

- id: `int64`
- name: `TicketTypeName`
- price: `float64`

**Event**

- id: `int64`
- name: `string`
- date: `timestamp`

**Ticket**

- id: `int64`
- event_id `int64`
- ticket_type_id `int64`
- quantity `int`

**Order**

- id: `int64`
- customer_id: `int64`
- ticket_id: `int64`
- quantity: `int`
- total_price: `float64`
- created_at: `timestamp`

## Database Schema

[`^ back to top ^`](#table-of-contents)

```mermaid
erDiagram
    Customer ||--o{ Order : order
    Customer {
        int64 id PK
        string username
        float64 balance
    }
    Ticket }o--|| TicketType : has
    TicketType {
        int64 id PK
        string name
        float64 price
    }
    Event ||--|{ Ticket : has
    Event {
        int64 id PK
        string name
        datetime date
    }
    Order }o--|| Ticket : has
    Ticket {
        int64 id PK
        int64 event_id FK
        int64 ticket_type_id FK
        int quantity
    }
    Order {
        int64 id PK
        int64 customer_id FK
        int64 ticket_id FK
        int quantity
        float64 total_price
        datetime created_at
    }
```

## API Endpoints

[`^ back to top ^`](#table-of-contents)

| **Method** | **Pattern**                 | **Description**                                 |
| ---------- | --------------------------- | ----------------------------------------------- |
| GET        | /api/customers              | View list of customers & their orders.          |
| POST       | /api/customers              | Add a new customer.                             |
| GET        | /api/customers/:id          | View a customer.                                |
| PATCH      | /api/customers/:id/balances | Add balance amount.                             |
| GET        | /api/events                 | View list of events with the tickets available. |
| GET        | /api/events/:id             | View an event with the tickets available.       |
| GET        | /api/tickets                | View list of tickets.                           |
| GET        | /api/tickets/:id            | View a ticket.                                  |
| GET        | /api/orders                 | View list of orders.                            |
| POST       | /api/orders                 | Order a ticket.                                 |

## Tech Stack

[`^ back to top ^`](#table-of-contents)

- Language: [Go 1.22](https://go.dev)
- Web Framework: [Gin](https://gin-gonic.com)
- DBMS: [PostgreSQL](https://www.postgresql.org)
- Database Migration: [migrate](https://github.com/golang-migrate/migrate)

## Run Locally

[`^ back to top ^`](#table-of-contents)

- Make sure you have [Go 1.22](https://go.dev), [PostgreSQL](https://www.postgresql.org), [migrate](https://github.com/golang-migrate/migrate) installed on your computer. Run these commands to check whether the tools are already installed. The terminal will output the version number if it is installed.

  ```bash
  go version
  ```

  ```bash
  psql --version
  ```

  ```bash
  migrate -version
  ```

- Connect to the PostgreSQL server by providing a user name & password.

  ```bash
  psql -U postgres
  ```

  Then create a database. You can name it as `evento`.

  ```sql
  CREATE DATABASE evento;
  ```

- Clone the repo.

  ```bash
  git clone https://github.com/nadiannis/evento.git
  ```

  ```bash
  cd evento
  ```

- Install the dependencies.

  ```bash
  go mod tidy
  ```

- Apply migrations.

  Provide the DSN (data source name): `<db-driver>://<db-username>:<db-password>@localhost:5432/<db-name>`.

  ```bash
  migrate -path ./migrations -database pgx://postgres:pass1234@localhost:5432/evento up
  ```

- Run the server.

  By default, the web server will run on port 8080 & the DSN is `postgres://postgres:pass1234@localhost:5432/evento`.

  ```bash
  go run ./cmd
  ```

  You can change the port & the DSN by running the server with flag.

  This command will run the server on port 4000 & connect to the database with DSN `postgres://myusername:mypassword@localhost:5432/mydbname`.

  ```bash
  go run ./cmd -port 4000 -db-dsn postgres://myusername:mypassword@localhost:5432/mydbname
  ```

## Screenshots

[`^ back to top ^`](#table-of-contents)

### `POST /api/customers` - add a new customer

<details>
<summary>Request</summary>

![Add a new customer (request)](./docs/img/customer-add-request.png)

</details>

<details>
<summary>Response</summary>

![Add a new customer (response)](./docs/img/customer-add-response.png)

</details>

<details>
<summary>Log</summary>

![Add a new customer (log)](./docs/img/customer-add-log.png)

</details>

### `PATCH /api/customers/:id/balances` - add balance amount

<details>
<summary>Request</summary>

![Add balance amount (request)](./docs/img/customer-balance-update-request.png)

</details>

<details>
<summary>Response</summary>

![Add balance amount (response)](./docs/img/customer-balance-update-response.png)

</details>

<details>
<summary>Log</summary>

![Add balance amount (log)](./docs/img/customer-balance-update-log.png)

</details>

### `GET /api/customers` - view list of customers & their orders

<details>
<summary>Request</summary>

![View all customers (request)](./docs/img/customer-list-request.png)

</details>

<details>
<summary>Response</summary>

![View all customers (response)](./docs/img/customer-list-response.png)

</details>

<details>
<summary>Log</summary>

![View all customers (log)](./docs/img/customer-list-log.png)

</details>

### `GET /api/customers/:id` - view a customer

<details>
<summary>Request</summary>

![View a customer (request)](./docs/img/customer-get-request.png)

</details>

<details>
<summary>Response</summary>

![View a customer (response)](./docs/img/customer-get-response.png)

</details>

<details>
<summary>Log</summary>

![View a customer (log)](./docs/img/customer-get-log.png)

</details>

### `GET /api/events` - view list of events with the tickets available

<details>
<summary>Request</summary>

![View list of events (request)](./docs/img/event-list-request.png)

</details>

<details>
<summary>Response</summary>

![View list of events (response)](./docs/img/event-list-response.png)

</details>

<details>
<summary>Log</summary>

![View list of events (log)](./docs/img/event-list-log.png)

</details>

### `GET /api/events/:id` - view an event with the tickets available

<details>
<summary>Request</summary>

![View an event (request)](./docs/img/event-get-request.png)

</details>

<details>
<summary>Response</summary>

![View an event (response)](./docs/img/event-get-response.png)

</details>

<details>
<summary>Log</summary>

![View an event (log)](./docs/img/event-get-log.png)

</details>

### `GET /api/tickets` - view list of tickets

<details>
<summary>Request</summary>

![View list of tickets (request)](./docs/img/ticket-list-request.png)

</details>

<details>
<summary>Response</summary>

![View list of tickets (response)](./docs/img/ticket-list-response.png)

</details>

<details>
<summary>Log</summary>

![View list of tickets (log)](./docs/img/ticket-list-log.png)

</details>

### `GET /api/tickets/:id` - view a ticket

<details>
<summary>Request</summary>

![View a ticket (request)](./docs/img/ticket-get-request.png)

</details>

<details>
<summary>Response</summary>

![View a ticket (response)](./docs/img/ticket-get-response.png)

</details>

<details>
<summary>Log</summary>

![View a ticket (log)](./docs/img/ticket-get-log.png)

</details>

### `POST /api/orders` - order a ticket (success)

<details>
<summary>Request</summary>

![Order a ticket (success) (request)](./docs/img/order-success-request.png)

</details>

<details>
<summary>Response</summary>

![Order a ticket (success) (response)](./docs/img/order-success-response.png)

</details>

<details>
<summary>Log</summary>

![Order a ticket (success) (log)](./docs/img/order-success-log.png)

</details>

<details>
<summary>Ticket quantity deducted</summary>

![Order a ticket (success) (ticket quantity deducted)](./docs/img/order-success-ticket-deducted.png)

</details>

<details>
<summary>Balance deducted</summary>

![Order a ticket (success) (balance deducted)](./docs/img/order-success-balance-deducted.png)

</details>

### `POST /api/orders` - order a ticket (insufficient ticket quantity)

<details>
<summary>Request</summary>

![Order a ticket (insufficient ticket quantity) (request)](./docs/img/order-insufficient-ticket-request.png)

</details>

<details>
<summary>Response</summary>

![Order a ticket (insufficient ticket quantity) (response)](./docs/img/order-insufficient-ticket-response.png)

</details>

<details>
<summary>Log</summary>

![Order a ticket (insufficient ticket quantity) (log)](./docs/img/order-insufficient-ticket-log.png)

</details>

<details>
<summary>Ticket quantity not deducted</summary>

![Order a ticket (success) (ticket quantity not deducted)](./docs/img/order-insufficient-ticket-ticket-not-deducted.png)

</details>

<details>
<summary>Balance not deducted</summary>

![Order a ticket (success) (balance not deducted)](./docs/img/order-insufficient-ticket-balance-not-deducted.png)

</details>

### `POST /api/orders` - order a ticket (insufficient balance)

<details>
<summary>Request</summary>

![Order a ticket (insufficient balance) (request)](./docs/img/order-insufficient-balance-request.png)

</details>

<details>
<summary>Response</summary>

![Order a ticket (insufficient balance) (response)](./docs/img/order-insufficient-balance-response.png)

</details>

<details>
<summary>Log</summary>

![Order a ticket (insufficient balance) (log)](./docs/img/order-insufficient-balance-log.png)

</details>

<details>
<summary>Ticket quantity not deducted</summary>

![Order a ticket (success) (ticket quantity not deducted)](./docs/img/order-insufficient-balance-ticket-not-deducted.png)

</details>

<details>
<summary>Balance not deducted</summary>

![Order a ticket (success) (balance not deducted)](./docs/img/order-insufficient-balance-balance-not-deducted.png)

</details>

### `GET /api/orders` - view list of orders

<details>
<summary>Request</summary>

![View list of orders (request)](./docs/img/order-list-request.png)

</details>

<details>
<summary>Response</summary>

![View list of orders (response)](./docs/img/order-list-response.png)

</details>

<details>
<summary>Log</summary>

![View list of orders (log)](./docs/img/order-list-log.png)

</details>
