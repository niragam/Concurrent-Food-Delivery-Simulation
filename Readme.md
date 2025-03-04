How to run:
1. Make sure Go is installed on your system 
2. Edit the Configuration (config.json) as needed:
    Number of producers, their FoodType, how many orders, queue sizes, etc.
    Zone worker counts and queue sizes.
    The HTTP server port.
3.Build and Run:
In terminal run inside the exercise folder:
go mod init ex4
go mod tidy
go run .

4. Access the Display:
Visit http://localhost:8080/orders
Check the developer console to see the orders logged.