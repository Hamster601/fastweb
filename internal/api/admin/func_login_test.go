package admin

import (
	"bytes"
	"encoding/json"
	"github.com/Hamster601/fastweb/pkg/pkglog"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_LoginNew(t *testing.T) {
	router := gin.Default()
	router.POST("/login", New(pkglog.ProjectLogger).LoginNew)
	body := map[string]string{
		"username": "admin",
		"password": "123",
	}
	w := httptest.NewRecorder()
	bodyByte, er := json.Marshal(body)
	if er != nil {
		pkglog.ProjectLogger.Fatal("marshal body failed:")
	}
	req, _ := http.NewRequest("POST", "/v1/api/login", bytes.NewReader(bodyByte))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
