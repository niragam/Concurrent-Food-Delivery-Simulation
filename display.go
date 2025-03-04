/*
307832097 Roy Kosary
311319313 Nir Agam
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// StartHTTPServer starts an HTTP server that displays the current orders.
func StartHTTPServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html>
		<head><meta charset="utf-8"><title>Orders</title></head>
		<body>
		  <h1>Orders</h1>
		  <pre id="output"></pre>
		  <script>
		    let intervalId;
		    async function fetchOrders() {
		      try {
		        const res = await fetch('/orders/data');
		        const data = await res.json();
		        console.log('Fetched orders:', data);
		        const outputEl = document.getElementById('output');
		        outputEl.textContent = JSON.stringify(data, null, 2);

		        if (data.includes("DONE")) {
		          clearInterval(intervalId);
		          console.log("All done. Stopped fetching.");
		        }
		      } catch (err) {
		        console.error(err);
		      }
		    }
		    fetchOrders();
		    intervalId = setInterval(fetchOrders, 1000);
		  </script>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/orders/data", func(w http.ResponseWriter, r *http.Request) {
		completedOrdersMu.Lock()
		snapshot := make([]Order, len(completedOrders))
		copy(snapshot, completedOrders)
		doneState := allDone
		completedOrdersMu.Unlock()

		var output []interface{}
		for _, o := range snapshot {
			entry := fmt.Sprintf("%s: %s %d", o.Restaurant, o.FoodType, o.OrderNumber)
			output = append(output, entry)
		}
		if doneState {
			output = append(output, "DONE")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(output)
	})

	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}
