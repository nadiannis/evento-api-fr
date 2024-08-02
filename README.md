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
  psql -U root
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

  ```bash
  migrate -path ./migrations -database pgx://postgres:pass1234@localhost:5432/evento up
  ```

- Run the server. The web server will run on port 8080.

  ```bash
  go run ./cmd
  ```

## Screenshots

[`^ back to top ^`](#table-of-contents)

### Add a new customer

![Add a new customer](./docs/img/customer-add.png)

### View an event with the tickets available

![View an event](./docs/img/event-get.png)

### Order a ticket (success)

![Order a ticket (success)](./docs/img/order-success.png)

### Order a ticket (ticket out of stock)

![Order a ticket (ticket out of stock)](./docs/img/order-error-1.png)

### Order a ticket (insufficient balance)

![Order a ticket (insufficient balance)](./docs/img/order-error-2.png)

### Log

![Log](./docs/img/log.png)
