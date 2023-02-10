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

/*3rd Segment*/
/*
package main

import (
	"fmt"
	"sync"
)

// DataRepository interface defines the operations for a CRUD service
type DataRepository interface {
	Create(key string, value interface{}) error
	Read(key string) (interface{}, error)
	Update(key string, value interface{}) error
	Delete(key string) error
}

// MapDataRepository implements the DataRepository interface using a map
type MapDataRepository struct {
	data map[string]interface{}
	mu   sync.Mutex
}

// Create adds a new key-value pair to the map
func (m *MapDataRepository) Create(key string, value interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; ok {
		return fmt.Errorf("Key %s already exists", key)
	}
	m.data[key] = value
	return nil
}

// Read retrieves the value for a given key from the map
func (m *MapDataRepository) Read(key string) (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if value, ok := m.data[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("Key %s does not exist", key)
}

// Update updates the value for a given key in the map
func (m *MapDataRepository) Update(key string, value interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; ok {
		m.data[key] = value
		return nil
	}
	return fmt.Errorf("Key %s does not exist", key)
}

// Delete removes a key-value pair from the map
func (m *MapDataRepository) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; ok {
		delete(m.data, key)
		return nil
	}
	return fmt.Errorf("Key %s does not exist", key)
}

func main() {
	repo := &MapDataRepository{
		data: make(map[string]interface{}),
	}

	repo.Create("key1", "value1")
	repo.Create("key2", 2)

	value, _ := repo.Read("key1")
	fmt.Println(value) // Output: value1

	repo.Update("key2", 3)
	value, _ = repo.Read("key2")
	fmt.Println(value) // Output: 3

	repo.Delete("key1")
	_, err := repo.Read("key1")
	fmt.Println(err) // Output: Key key1 does not exist
}

*/

/*Segment 4
In this example, the DataStore interface defines the operations that can be performed on the data store. 
The DataService struct implements the DataStore interface and uses a mutex to synchronize access to the map[string]interface{} data store. 
The main function demonstrates the use of the Create, Read, Update, and Delete operations on the data store.
*/

/*
package main

import (
	"fmt"
	"sync"
)

type DataStore interface {
	Create(key string, value interface{})
	Read(key string) interface{}
	Update(key string, value interface{})
	Delete(key string)
}

type DataService struct {
	data map[string]interface{}
	mux  sync.Mutex
}

func (ds *DataService) Create(key string, value interface{}) {
	ds.mux.Lock()
	defer ds.mux.Unlock()
	ds.data[key] = value
}

func (ds *DataService) Read(key string) interface{} {
	ds.mux.Lock()
	defer ds.mux.Unlock()
	return ds.data[key]
}

func (ds *DataService) Update(key string, value interface{}) {
	ds.mux.Lock()
	defer ds.mux.Unlock()
	ds.data[key] = value
}

func (ds *DataService) Delete(key string) {
	ds.mux.Lock()
	defer ds.mux.Unlock()
	delete(ds.data, key)
}

func main() {
	dataService := &DataService{
		data: make(map[string]interface{}),
	}
	dataService.Create("name", "John Doe")
	dataService.Create("age", 30)
	dataService.Create("address", "123 Main St.")
	fmt.Println("Name:", dataService.Read("name"))
	fmt.Println("Age:", dataService.Read("age"))
	fmt.Println("Address:", dataService.Read("address"))
	dataService.Update("name", "Jane Doe")
	dataService.Update("age", 35)
	dataService.Update("address", "456 Main St.")
	fmt.Println("Name:", dataService.Read("name"))
	fmt.Println("Age:", dataService.Read("age"))
	fmt.Println("Address:", dataService.Read("address"))
	dataService.Delete("name")
	dataService.Delete("age")
	dataService.Delete("address")
	fmt.Println("Data:", dataService.data)
}

*/

/*Segment 5

*/

/*
package main

import (
	"sync"
)

// DataStore is an interface that defines the methods for CRUD operations.
type DataStore interface {
	Create(key string, data map[string]interface{})
	Read(key string) map[string]interface{}
	Update(key string, data map[string]interface{})
	Delete(key string)
}

// MapDataStore is an implementation of the DataStore interface that uses a map.
type MapDataStore struct {
	data map[string]map[string]interface{}
	mutex sync.Mutex
}

// NewMapDataStore returns a new instance of MapDataStore.
func NewMapDataStore() *MapDataStore {
	return &MapDataStore{
		data: make(map[string]map[string]interface{}),
	}
}

// Create adds a new entry to the map.
func (m *MapDataStore) Create(key string, data map[string]interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	m.data[key] = data
}

// Read retrieves an entry from the map.
func (m *MapDataStore) Read(key string) map[string]interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	return m.data[key]
}

// Update updates an existing entry in the map.
func (m *MapDataStore) Update(key string, data map[string]interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	m.data[key] = data
}

// Delete removes an entry from the map.
func (m *MapDataStore) Delete(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	delete(m.data, key)
}

func main() {
	dataStore := NewMapDataStore()
	
	// Example usage:
	dataStore.Create("key1", map[string]interface{}{
		"name": "John Doe",
		"age": 30,
	})
	
	data := dataStore.Read("key1")
	println(data) // Output: map[age:30 name:John Doe]
	
	dataStore.Update("key1", map[string]interface{}{
		"name": "Jane Doe",
		"age": 31,
	})
	
	data = dataStore.Read("key1")
	println(data) // Output: map[age:31 name:Jane Doe]
	
	dataStore.Delete("key1")
	
	data = dataStore.Read("key1")
	println(data) // Output: map[]
}
*/

/*Unit Testing*/
/*
package router

import (

	"fmt"
	"my-web-server/logger"
	"my-web-server/kvStore"
	"my-web-server/validations"
	"net/http"
)


const PortNumber = ":8080"

type requestHandler struct {
	store kvStore.DataStore
}
 
func InitialiseRoutes(store kvStore.DataStore) {
	httpHandler := requestHandler{
		store: store,
	}

	http.HandleFunc("/read", httpHandler.Retrieve)      //get
	http.HandleFunc("/create", httpHandler.CreateData)  //post
	http.HandleFunc("/update", httpHandler.UpdateData)  //put
	http.HandleFunc("/delete", httpHandler.DeleteData)  //delete
	logger.Info(fmt.Sprintf("Starting application on port %s", PortNumber))
	err := http.ListenAndServe(PortNumber, nil)
	logger.Error(fmt.Sprintf("Error with application on port %s : %v", PortNumber, err))
}

//Retrieve Gets/Displays all items
func (h *requestHandler) Retrieve(w http.ResponseWriter, r *http.Request) {
	logger.Info("User called the retrieve method")

	key := r.URL.Query().Get("key")
	//fmt.Println(key)

	if !validations.ValidateRequestMethod(w, r, http.MethodGet){
		return
	}

	value, _ := h.store.ReadProduct(key)

		switch b := value.(type) {
		case string:
			fmt.Fprintf(w, "%q string \n", b)
		case float64:
			fmt.Fprintf(w, "%v float \n", b)
		case bool:
			fmt.Fprintf(w, "%v bool \n", b)
		default:
			fmt.Fprintf(w, "%v default \n ", b)
	}
}


//CreateData Creates new items
func (h *requestHandler) CreateData(w http.ResponseWriter, r *http.Request) {
	logger.Info("User called the createData method")
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if !validations.ValidateRequestMethod(w, r, http.MethodPost){
		return
	}
	
	err := h.store.CreateProduct(key, value)

	if err != nil {
		fmt.Fprintf(w, "%v \n", err.Error())
		w.WriteHeader(http.StatusConflict)
	}
	w.WriteHeader(http.StatusCreated)

}

//UpdateData Updates existing item
func (h *requestHandler) UpdateData(w http.ResponseWriter, r *http.Request) {
	logger.Info("User called the updateData method")
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if !validations.ValidateRequestMethod(w, r, http.MethodPut){
		return
	}
	
	err := h.store.UpdateProduct(key, value)
	fmt.Println(err)

	if err != nil {
		fmt.Fprintf(w, "%v \n", err.Error())
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusNoContent)
}


//DeleteData Deletes an existing item
func (h *requestHandler) DeleteData(w http.ResponseWriter, r *http.Request) {
	logger.Info("User called the deleteData method")
	key := r.URL.Query().Get("key")

	err := h.store.DeleteProduct(key)

	if err != nil {
		fmt.Fprintf(w, "%v \n", err.Error())
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusNoContent)
}
*/


/*T-Result*/

/*
package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrieve(t *testing.T) {
	store := kvStore.DataStore{
		Data: make(map[string]interface{}),
	}
	handler := &requestHandler{
		store: store,
	}
	req, err := http.NewRequest(http.MethodGet, "/read?key=testKey", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.Retrieve(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"\"Key not found\"" string `
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateData(t *testing.T) {
	store := kvStore.DataStore{
		Data: make(map[string]interface{}),
	}
	handler := &requestHandler{
		store: store,
	}
	req, err := http.NewRequest(http.MethodPost, "/create?key=testKey&value=testValue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.CreateData(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestUpdateData(t *testing.T) {
	store := kvStore.DataStore{
		Data: make(map[string]interface{}),
	}
	store.CreateProduct("testKey", "testValue")
	handler := &requestHandler{
		store: store,
	}
	req, err := http.NewRequest(http.MethodPut, "/update?key=testKey&value=newValue", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.UpdateData(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestDeleteData(t *testing.T) {
	store := kvStore.DataStore{
		Data: make(map[string]interface{}),
	}
	store.CreateProduct("testKey", "testValue")
	handler := &requestHandler{
		store: store,

*/
