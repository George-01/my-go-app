package main

import (
	"fmt"
	"net/http"

	"github.com/George-01/my-go-app/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/abt", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}


/*
func CreateData(rw http.ResponseWriter, r *http.Request) {
	logger.Info("User called the createData method")

	mutex.Lock()
	defer mutex.Unlock()
		
	if !validations.ValidateRequestMethod(rw, r, http.MethodPost){
		return
	}
	
	decoder := json.NewDecoder(r.Body)
	var newProducts map[string]interface{}
	err := decoder.Decode(&newProducts)
	if err != nil {
		fmt.Fprintf(rw, "Error decoding update request: %v", err)
		logger.Error(fmt.Sprintf("Error decoding update request: %v", err))
		return
	}

	for key, value := range newProducts{

		if _, ok := kvStore.Products[key]; ok{
			rw.WriteHeader(http.StatusConflict)
			fmt.Fprintf(rw, "This key: %v already exist\n", key)
			logger.Error(fmt.Sprintf("This key: %v already exist\n", key))
			return
		} 		
		kvStore.Products[key] = value
	}
	rw.WriteHeader(http.StatusCreated)
	logger.Info("User created new items")
}
*/

/*
func Retrieve(w http.ResponseWriter, r *http.Request) {
	logger.Info("User called the retrieve method")

	mutex.Lock()
	defer mutex.Unlock()

	if !validations.ValidateRequestMethod(w, r, http.MethodGet){
		return
	}

	if kvStore.Products == nil {
		http.Error(w, "No data available yet", http.StatusNotFound)
		logger.Error("No data available yet")
	}
	for k, v := range kvStore.Products{
		switch b := v.(type) {
		case string:
			fmt.Fprintf(w, "%v : %q\n", k, b)
		case float64:
			fmt.Fprintf(w, "%v : %v\n", k, b)
		case bool:
			fmt.Fprintf(w, "%v : %v\n", k, b)
		default:
			fmt.Fprintf(w, "%v : %v\n", k, b)
		}
	}
	logger.Info("list of all items")
}
*/
