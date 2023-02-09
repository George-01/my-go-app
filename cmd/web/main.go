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

/*
func UpdateData(rw http.ResponseWriter, r *http.Request) {
	logger.Info("User called updateData method")

	mutex.Lock()
	defer mutex.Unlock()

	mutex.Lock()
	defer mutex.Unlock()

	if !validations.ValidateRequestMethod(rw, r, http.MethodPut){
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

		if _, ok := kvStore.Products[key]; !ok{
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(rw, "This key: %v does not exist\n", key)
			logger.Error(fmt.Sprintf("This key: %v does not exist\n", key))
			return
		} 		
		kvStore.Products[key] = value
	}
	rw.WriteHeader(http.StatusOK)
}
*/

/*
func DeleteData(rw http.ResponseWriter, r *http.Request) {
	logger.Info("User called deleteData method")

	mutex.Lock()
	defer mutex.Unlock()

	if !validations.ValidateRequestMethod(rw, r, http.MethodDelete){
		return
	}

	qryData := r.URL.RawQuery
	fmt.Printf("the key is %v\n", qryData)
	
	if _, ok := kvStore.Products[qryData]; !ok{
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "This key: %v does not exist\n", qryData)
		logger.Error(fmt.Sprintf("This key: %v does not exist\n", qryData))
		return
	} 
	//delete entry
	delete(kvStore.Products, qryData)
	rw.WriteHeader(http.StatusNoContent)
	logger.Info(fmt.Sprintf("User deleted an entry with key: %v\n", qryData))
}
*/

/*
func ListData(w http.ResponseWriter, r *http.Request) {

	mutex.Lock()
	defer mutex.Unlock()

	if Products == nil {
		http.Error(w, "No data available yet", http.StatusNotFound)
		logger.Error("No data available yet")
	}
	for k, v := range Products{
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

/*
var Products = map[string]interface{} {

	"name":"Apple Mac Book Pro",
	"price": 1999.99,
	"inStock": true,
	"categories": []string{"Computers", "Laptop", "Apple"},
}
*/

/*Updated*/
/*
package main

import (
	"fmt"
)

type Repository interface {
	Create(map[string]interface{})
	Read(string) map[string]interface{}
	Update(string, map[string]interface{})
	Delete(string)
}

type MapRepository struct {
	data map[string]map[string]interface{}
}

func (m *MapRepository) Create(data map[string]interface{}) {
	id := data["id"].(string)
	m.data[id] = data
}

func (m *MapRepository) Read(id string) map[string]interface{} {
	return m.data[id]
}

func (m *MapRepository) Update(id string, data map[string]interface{}) {
	m.data[id] = data
}

func (m *MapRepository) Delete(id string) {
	delete(m.data, id)
}

type Service struct {
	repo Repository
}

func (s *Service) Create(data map[string]interface{}) {
	s.repo.Create(data)
}

func (s *Service) Read(id string) map[string]interface{} {
	return s.repo.Read(id)
}

func (s *Service) Update(id string, data map[string]interface{}) {
	s.repo.Update(id, data)
}

func (s *Service) Delete(id string) {
	s.repo.Delete(id)
}

func main() {
	repo := &MapRepository{
		data: make(map[string]map[string]interface{}),
	}
	service := &Service{
		repo: repo,
	}

	data := map[string]interface{}{
		"id":   "1",
		"name": "John Doe",
		"age":  30,
	}
	service.Create(data)

	readData := service.Read("1")
	fmt.Println("Read Data:", readData)

	data = map[string]interface{}{
		"id":   "1",
		"name": "Jane Doe",
		"age":  32,
	}
	service.Update("1", data)

	readData = service.Read("1")
	fmt.Println("Read Data after update:", readData)

	service.Delete("1")

	readData = service.Read("1")
	fmt.Println("Read Data after delete:", readData)
}

*/
/*Notes:
The Repository interface defines the methods for a CRUD service.
The MapRepository struct implements the Repository interface using the map[string]interface{} data structure.
The Service struct implements the CRUD methods using the Repository interface, which can be any implementation of the Repository interface.
In the main function, we create an instance of MapRepository and Service and demonstrate the CRUD operations.
*/

/*Segment_2*/
/*
package main

import (
	"fmt"
	"sync"
)

// Single Responsibility Principle - this interface is responsible for performing CRUD operations
type CrudOperations interface {
	Create(string, interface{})
	Read(string) (interface{}, error)
	Update(string, interface{}) error
	Delete(string) error
}

// Open/Closed Principle - this struct implements the CrudOperations interface, but can be extended to add more functionality
type CrudService struct {
	data map[string]interface{}
	mux  sync.Mutex
}

// Liskov Substitution Principle - this method satisfies the CrudOperations interface and can be used as a replacement
func (c *CrudService) Create(key string, value interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data[key] = value
}

// Liskov Substitution Principle - this method satisfies the CrudOperations interface and can be used as a replacement
func (c *CrudService) Read(key string) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	value, exists := c.data[key]
	if !exists {
		return nil, fmt.Errorf("Key not found")
	}
	return value, nil
}

// Liskov Substitution Principle - this method satisfies the CrudOperations interface and can be used as a replacement
func (c *CrudService) Update(key string, value interface{}) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, exists := c.data[key]
	if !exists {
		return fmt.Errorf("Key not found")
	}
	c.data[key] = value
	return nil
}

// Liskov Substitution Principle - this method satisfies the CrudOperations interface and can be used as a replacement
func (c *CrudService) Delete(key string) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, exists := c.data[key]
	if !exists {
		return fmt.Errorf("Key not found")
	}
	delete(c.data, key)
	return nil
}

func main() {
	crudService := &CrudService{
		data: make(map[string]interface{}),
	}
	crudService.Create("Key1", "Value1")
	crudService.Create("Key2", 2)
	value, _ := crudService.Read("Key1")
	fmt.Println(value)
	crudService.Update("Key1", "UpdatedValue1")
	value, _ = crudService.Read("Key1")
	fmt.Println(value)
	crudService.Delete("Key2")
	_, err := crudService.Read("Key2")
	fmt.Println(err)
}
``

*/
