package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	want := greetings{
		Message: "வணக்கம் கமேஷ்!",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?name=Kamesh&lang=ta", nil)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if assert.NoError(t, greet(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		var got greetings
		assert.NoError(t, json.Unmarshal(res.Body.Bytes(), &got))
		assert.Equal(t, want.Message, got.Message)
	}
}
