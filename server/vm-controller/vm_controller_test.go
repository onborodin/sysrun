/*
 * Copyright 2020 Oleg Borodin  <borodin@unix7.org>
 */

package vmController

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "bytes"
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"

    "sysrun/server/vm-model"
    "sysrun/config"
)

func TestList(t *testing.T) {
    configuration := config.New()
    controller := New(configuration)

    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.GET("/list", controller.List)

    request, err := http.NewRequest(http.MethodGet, "/list", nil)
    if err != nil {
        t.Fatalf("couldn't create request: %v\n", err)
    }
    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusOK, recorder.Code, "Not equal response code")

    response := Response{}
    err = json.Unmarshal(recorder.Body.Bytes(), &response)
    assert.Equal(t, response.Error, false, "Model or controller error")
    fmt.Println(recorder.Body.String())
}


func TestStart(t *testing.T) {
    configuration := config.New()
    controller := New(configuration)

    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/start", controller.Start)

    vm := vmModel.VM{
        Id: 12345,
    }

    data, err := json.Marshal(vm)
    if err != nil {
        t.Error(err)
    }

    fmt.Println("request:", string(data))

    reader := bytes.NewReader(data)
    request, err := http.NewRequest(http.MethodPost, "/start", reader)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    request.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusOK, recorder.Code, "Not equal response code")

    response := Response{}
    err = json.Unmarshal(recorder.Body.Bytes(), &response)
    assert.Equal(t, response.Error, false, "Model or controller error")
    fmt.Println("response body:", recorder.Body.String())
}

func TestStop(t *testing.T) {
    configuration := config.New()
    controller := New(configuration)

    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/stop", controller.Stop)

    vm := vmModel.VM{
        Id: 12345,
    }

    data, err := json.Marshal(vm)
    if err != nil {
        t.Error(err)
    }

    fmt.Println("request:", string(data))

    reader := bytes.NewReader(data)
    request, err := http.NewRequest(http.MethodPost, "/stop", reader)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    request.Header.Set("Content-Type", "application/json")
    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    assert.Equal(t, http.StatusOK, recorder.Code, "Not equal response code")

    response := Response{}
    err = json.Unmarshal(recorder.Body.Bytes(), &response)

    assert.Equal(t, response.Error, false, "Model or controller error")
    fmt.Println("response body:", recorder.Body.String())

}
