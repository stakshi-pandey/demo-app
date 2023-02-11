package controllers

import (
	"bytes"
	"log"
	"os"

	"demo-app/initializers"
	"demo-app/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	migrate()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	router.POST("/add-post", CreatePost)
	return router
}

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initializers.ConnectToDB()
	migrate()
}

func migrate() {
	// migrator := initializers.DB.Migrator()
	// migrator.DropTable(&models.Post{})
	initializers.DB.AutoMigrate(&models.Post{})
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {

	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	_, _ = http.DefaultClient.Do(request)
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		Title: "PARTYPARTYPARTYYEAH",
		Body:  "BYJK",
	}

	writer := makeRequest("POST", "/add-post", post)
	t.Log("\nREQ SRVER HTTP   :::   ", writer)
	t.Log("\n\n\n", writer)
	assert.Equal(t, 200, writer.Code)

}
