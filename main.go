package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

// Request structure to parse incoming JSON payload
type SortRequest struct {
	ToSort [][]int `json:"to_sort"`
}

// Response structure
type SortResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

// func main() {
// 	fmt.Print("\tmain process started\n") ////////////////////////////////////////////////

// 	http.HandleFunc("/process-single", ProcessSingle)
// 	http.HandleFunc("/process-concurrent", ProcessConcurrent)

// 	fmt.Print("\tStarting server at localhost:8000\n") ////////////////////////////////////////////////
// 	http.ListenAndServe("localhost:8000", nil)

// 	fmt.Print("\tmain process finished\n") ////////////////////////////////////////////////
// }

func main() {
	fmt.Print("\tmain process started\n") ////////////////////////////////////////////////
    http.HandleFunc("/process-single", ProcessSingle)
    http.HandleFunc("/process-concurrent", ProcessConcurrent)

    // Enable CORS (allow all origins)
    corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == http.MethodOptions {
            return
        }
        http.DefaultServeMux.ServeHTTP(w, r)
    })

    http.Handle("/", corsHandler)

	fmt.Print("\tCORS started, Starting server at localhost:8000\n") ////////////////////////////
    http.ListenAndServe(":8000", nil)

	fmt.Print("\tmain process finished\n") ////////////////////////////////////////////////
}

func ProcessSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\tProcessSingle called\n") ////////////////////////////////////////////////

	var request SortRequest
	var response SortResponse

	// Parse JSON payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Process and measure time for sequential sorting
	startTime := time.Now()
	fmt.Println("Start Time:", startTime) // Print start time for debugging
	response.SortedArrays = SequentialSort(request.ToSort)
	endTime := time.Now()
	fmt.Println("End Time:", endTime) // Print end time for debugging
	response.TimeNs = endTime.Sub(startTime).Nanoseconds()

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ProcessConcurrent(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\tProcessConcurrent called\n") ////////////////////////////////////////////////

	var request SortRequest
	var response SortResponse

	// Parse JSON payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Process and measure time for concurrent sorting
	startTime := time.Now()
	fmt.Println("Start Time:", startTime) // Print start time for debugging
	response.SortedArrays = ConcurrentSort(request.ToSort)
	endTime := time.Now()
	fmt.Println("End Time:", endTime) // Print end time for debugging
	response.TimeNs = endTime.Sub(startTime).Nanoseconds()

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SequentialSort(arrays [][]int) [][]int {
	// fmt.Print("\tSequentialSort called\n") ////////////////////////////////////////////////
	for i := range arrays {
		sort.Ints(arrays[i])
	}
	// fmt.Print(arrays, " Seq Sorted Array\n") ////////////////////////////////////////////////
	return arrays
}

func ConcurrentSort(arrays [][]int) [][]int {
	// fmt.Print("\tConcurrentSort called\n") ////////////////////////////////////////////////
	var wg sync.WaitGroup

	for i := range arrays {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sort.Ints(arrays[i])
		}(i)
	}

	wg.Wait()
	// fmt.Print(arrays, " Con Sorted Array\n") ////////////////////////////////////////////////
	return arrays
}
