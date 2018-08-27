package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
	"net/http/httptest"
)

func TestMain(m *testing.M){
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplate bool) *gin.Engine {
	r := gin.Default()
	if withTemplate{
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool ){
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}
