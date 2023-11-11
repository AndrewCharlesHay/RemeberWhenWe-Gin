package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/aws/aws-sdk-go/aws/credentials"
)


// MockCredentialsProvider is a mock implementation of the credentials.Provider interface
type MockCredentialsProvider struct {
	mock.Mock
	credentials.Provider
}
func TestPingRoute(t *testing.T) {


	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}


func TestAWSSetup(t *testing.T) {
	// Initialize a new AWS session using the default shared config
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		t.Fatalf("Error creating AWS session: %v", err)
	}

	t.Logf("AWS setup is successful")
	t.Logf(*sess.Config.Region)
}