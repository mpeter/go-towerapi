package authtoken

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthToken_Create(t *testing.T) {
	setup()
	defer teardown()

	var aUsername = "admin"
	var aPassword string = "admin"
	var aToken string = "8f17825cf08a7efea124f2638f3896f6637f8745"
	var aExpires string = "2023-09-05T21:46:35.729Z"

	createRequest := &AuthTokenCreateRequest{
		Username: aUsername,
		Password: aPassword,
	}

	mux.HandleFunc("/authtoken/", func(w http.ResponseWriter, r *http.Request) {
		v := new(AuthTokenCreateRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, createRequest) {
			t.Errorf("Request body = %+v, expected %+v", v, createRequest)
		}
		fmt.Fprintf(w, `{"token":"`+aToken+`","expires":"`+aExpires+`"}"`)
	})

	authToken, _, err := client.AuthToken.Create(createRequest)
	if err != nil {
		t.Errorf("AuthToken.Create returned error: %v", err)
	}

	expected := &AuthToken{
		Token:   aToken,
		Expires: aExpires,
	}
	if !reflect.DeepEqual(authToken, expected) {
		t.Errorf("AuthToken.Create returned %+v, expected %+v", authToken, expected)
	}
}
*/
