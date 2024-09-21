# GoBet

## Overview

GoBet is a project designed to simulate a betting environment where users can place bets on various events, track odds, and manage their financial transactions. This system provides a comprehensive backend solution built using Go, with Kafka for real-time event processing.

## Features

- **User Registration**: Users can create accounts and manage their profiles.
- **Event Management**: Users can view and bet on various events (e.g., sports matches, horse races).
- **Betting System**: Users can place bets with real-time odds and view their betting history.
- **Result Tracking**: The system tracks the outcomes of events and determines bet results.
- **Transaction Management**: Users can manage their balances through deposits, withdrawals, and winnings.

## Architecture

The system architecture consists of several key components:

1. **User**: Represents a bettor with an account, balance, and betting history.
2. **Event**: Represents events on which users can place bets.
3. **Bet**: Represents individual bets placed by users, linking them to events and odds.
4. **Result**: Stores the outcome of events to determine bet results.
5. **Odds**: Represents the odds for various outcomes of events.
6. **Transaction**: Tracks all financial transactions related to user accounts.

### Data Model

The following is a simplified data model for the system:

- **User**

  - `id`: UUID
  - `name`: VARCHAR
  - `email`: VARCHAR
  - `password`: VARCHAR
  - `balance`: DECIMAL(18, 2)
  - `createdAt`: TIMESTAMP
  - `updatedAt`: TIMESTAMP

- **Event**

  - `id`: UUID
  - `name`: VARCHAR
  - `startTime`: TIMESTAMP
  - `status`: ENUM (PENDING, IN_PROGRESS, COMPLETED)
  - `createdAt`: TIMESTAMP
  - `updatedAt`: TIMESTAMP

- **Bet**

  - `id`: UUID
  - `userId`: UUID (FK)
  - `eventId`: UUID (FK)
  - `amount`: DECIMAL(18, 2)
  - `oddsId`: UUID (FK)
  - `status`: ENUM (PENDING, WON, LOST, CANCELLED)
  - `createdAt`: TIMESTAMP
  - `updatedAt`: TIMESTAMP

- **Result**

  - `id`: UUID
  - `eventId`: UUID (FK)
  - `outcome`: VARCHAR
  - `createdAt`: TIMESTAMP

- **Odds**

  - `id`: UUID
  - `eventId`: UUID (FK)
  - `description`: VARCHAR
  - `value`: DECIMAL(18, 2)
  - `createdAt`: TIMESTAMP

- **Transaction**
  - `id`: UUID
  - `userId`: UUID (FK)
  - `type`: ENUM (DEPOSIT, WITHDRAW, WIN, LOSS)
  - `amount`: DECIMAL(18, 2)
  - `createdAt`: TIMESTAMP

## Technologies

- **Go**: For building the backend services.
- **Kafka**: For handling real-time event processing and messaging.
- **PostgreSQL**: For data storage.

## Setup Instructions

### Prerequisites

- Go (1.23 or later)
- PostgreSQL
- Kafka (including Zookeeper)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/gobet.git
   cd gobet
   go mod tidy
   ```
2. **Run the app**
   ```bash
   go run .
   ```
