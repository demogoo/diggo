package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/data"
	"github.com/demogoo/diggo/internal/models"
	"github.com/demogoo/diggo/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestItemList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := config.NewConfig()
	service := services.NewItemService(config)
	handler := NewItemHandler(service)

	e := gin.Default()
	e.GET("/items", handler.List)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/items", nil)
	e.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	items := make([]*models.Item, 0)
	err := json.Unmarshal(body, &items)
	if err != nil {
		t.Fatal(err, "=>", string(body))
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, items[0], data.Items[0])
}
