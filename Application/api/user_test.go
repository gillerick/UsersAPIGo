package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserToJSON(t *testing.T) {
	user := User{
		Name:   "Gill",
		Age:    25,
		Gender: "Male",
		Email:  "ogayogill95@gmail.com",
		Phone:  254725034298,
	}

	json := user.ToJSON()

	assert.Equal(t, `{"name":"Gill Erick","age":25,"gender":"Male","email":"ogayogill95@gmail.com","phone":"254725034298"}`,
		string(json), "User JSON Marshalling unsuccessful")

}

func TestUserFromJSON(t *testing.T)  {
	json := []byte(`{name":"Gill Erick","age":25,"gender":"male","email":"ogayogill95@gmail.com","phone":"254725034298"}`)
	user := FromJSON(json)

	assert.Equal(t, User{Name:"Gill Erick",Age:25,Gender: "Male",Email:"ogayogill95@gmail.com",Phone:254725034298},
		user, "User JSON Unmarshalling was unsuccessful")
	assert.True(t, true, "Implement Me.")
}

