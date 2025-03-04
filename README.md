# Food Delivery System Simulation

This project is a concurrent simulation of a food delivery system built in Go. It leverages goroutines, channels, and an HTTP server to model a real-world food order processing pipeline.

## Overview

The simulation consists of four main components:

1. **Producers (Restaurants)**  
   - Each restaurant generates food orders continuously until it completes its workload.
   - Orders are formatted as:  
     ```
     Restaurant <ID>: <FoodType> <OrderNumber>
     ```
   - When finished, the restaurant sends a `DONE` signal.

2. **Dispatcher**  
   - Collects orders from all restaurants using a **fan-in** pattern.
   - Routes orders to specific delivery zones (Pizza, Burger, Sushi) using **fan-out**.
   - Passes a `DONE` signal to all zones when all producers are finished.

3. **Zone Managers**  
   - Each zone (e.g., `PizzaZone`, `BurgerZone`, `SushiZone`) processes its assigned orders.
   - Uses a **token passing mechanism** to control concurrency:
     - Workers must acquire a token before processing an order.
     - Once processing is complete, the token is released for another worker.
   - Once all orders are processed, the zone signals completion.

4. **Display Manager (HTTP Server)**  
   - Displays completed orders in real time.
   - The `/orders` endpoint provides a JSON array of completed orders.
   - Logs processed orders to the browser console.
   - When all orders are processed, `"DONE"` is included in the response.

## Key Concepts

### Token Passing Mechanism

- **Purpose:** Controls concurrency to prevent overloading the system.
- **Implementation:**
  - A channel acts as a token pool.
  - Workers acquire a token before processing and return it after completion.
  - Ensures fair access and prevents excessive concurrent tasks.

### Concurrency Patterns

- **Fan-In:** Merges multiple producer channels into one.
- **Fan-Out:** Dispatches orders to multiple zone managers for parallel processing.
- **Worker Pools:** Limits concurrent processing using tokens.



## Getting Started

### Prerequisites

Before running the project, ensure that you have:

- [Go](https://golang.org/dl/) installed on your system.
- The required dependencies managed using Go modules.

### Installation & Setup

1. **Clone the Repository (if applicable):**
   ```sh
   git clone <repository-url>
   cd <project-folder>
   ```

2. **Initialize Go Modules (if needed):**
   ```sh
   go mod init ex4
   go mod tidy
   ```

3. **Update Configuration (Optional):**  
   Modify `config.json` to adjust the system settings, such as:
   - Number of producers and their queue sizes.
   - Number of workers per zone.
   - Dispatcher-to-zone communication queue size.
   - HTTP server port.

### Running the Simulation

1. **Run the Application:**
   ```sh
   go run .
   ```

2. **Monitor the Process:**
   - Orders are processed in real-time.
   - The system will print logs in the terminal as orders are generated, dispatched, processed, and completed.

### Viewing the Results

#### Web Interface
- Open your browser and visit:
  ```
  http://localhost:8080/orders
  ```
- The page will display a real-time list of completed orders.

#### Browser Console Logging
- Orders are also logged to the browser console.
- Open **Developer Tools** (F12 or right-click → "Inspect") and navigate to the **Console** tab to see real-time updates.

### Stopping the Simulation

To terminate the program, press:
```sh
Ctrl + C
```
in the terminal.

### Configuration

The system parameters can be customized in `config.json`. You can modify:

- **Producers:**
  - Number of restaurants.
  - Types of food they produce.
  - Number of orders they generate.
  - Queue sizes for each producer.

- **Zones:**
  - Number of workers assigned to each delivery zone.
  - Zone queue size.

- **HTTP Server:**
  - Port number for the display server.

Example `config.json`:

```json
{
  "Producers": [
    { "Restaurant": "Restaurant 1", "FoodType": "Pizza",  "Orders": 50, "QueueSize": 5 },
    { "Restaurant": "Restaurant 2", "FoodType": "Burger", "Orders": 100, "QueueSize": 8 },
    { "Restaurant": "Restaurant 3", "FoodType": "Sushi",  "Orders": 40, "QueueSize": 3 }
  ],
  "Zones": [
    { "Name": "PizzaZone",  "Workers": 2 },
    { "Name": "BurgerZone", "Workers": 4 },
    { "Name": "SushiZone",  "Workers": 1 }
  ],
  "ZoneQueueSize": 20,
  "HTTPServerPort": 8080
}
```
After modifying `config.json`, restart the simulation for changes to take effect.