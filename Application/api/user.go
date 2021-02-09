package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//User type: name, age, gender, email, phone number
//`json: field` is added to direct encoding/json on how the marshalling is to be done - to lowercase
type User struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Phone int `json:"phone"`
}

//Initializing a user map
var users = map[string]User{}

//Marshals the User type
func (u User)  ToJSON() []byte{
	ToJSON, err := json.Marshal(u)
	if err != nil{
		panic(err)
	}
	return ToJSON
}

//Unmarshal the User type
func FromJSON(data []byte) User {
	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil{
		panic(err)
	}
	return user

}

func writeJSON(w http.ResponseWriter, i interface{})  {
	u, err := json.Marshal(i)
	if err != nil{
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(u)
}

//Handles the API functionality
func RegistrationHandler(w http.ResponseWriter, r *http.Request){
	switch method := r.Method; method {
	case http.MethodGet:
		users := AllUsers()
		writeJSON(w, users)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		user := FromJSON(body)
		name, created := CreateUser(user)
		if created{
			w.Header().Add("Location", "/api/register/"+name)
			w.WriteHeader(http.StatusCreated)
		}else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Method unsupported"))


	}

}


//Function to fetch all users
func AllUsers() []User {
	records := make([]User, len(users))
	count := 0
	for _, user := range users {
		records[count] = user
		count ++
	}
	return records
}

//Function fetches a single user
func GetUser(name string) (User, bool) {
	user, found := users[name]
	return user, found
}


//Function creates a new user if they don't exist
func CreateUser(user User) (string, bool) {
	_, found := users[user.Name]
	if found{
		return "", false
	}
	users[user.Name] = user
	return user.Name, true
}


