/*
 * Copyright 2019 Oleg Borodin  <borodin@unix7.org>
 */

package vmController

import (
    "net/http"
    "fmt"
    "errors"
    "log"

    "github.com/gin-gonic/gin"

    "sysrun/config"
    "sysrun/server/vm-model"
)

type Controller struct {
    config *config.Config
    vm *vmModel.Model
}

type Response struct {
    Error       bool        `json:"error"`
    Message     string      `json:"message,omitempty"`
    Result      interface{} `json:"result,omitempty"`
}

func sendError(context *gin.Context, err error) {
    if err == nil {
        err = errors.New("undefined")
    }
    log.Printf("%s\n", err)
    response := Response{
        Error: true,
        Message: fmt.Sprintf("%s", err),
        Result: nil,
    }
    context.JSON(http.StatusOK, response)
}

func sendOk(context *gin.Context) {
    response := Response{
        Error: false,
        Message: "",
        Result: nil,
    }
    context.JSON(http.StatusOK, response)
}

func sendMessage(context *gin.Context, message string) {
    log.Printf("%s\n", message)
    response := Response{
        Error: false,
        Message: fmt.Sprintf("%s", message),
        Result: nil,
    }
    context.JSON(http.StatusOK, response)
}

func sendResult(context *gin.Context, result interface{}) {
    response := Response{
        Error: false,
        Message: "",
        Result: result,
    }
    context.JSON(http.StatusOK, &response)
}

func (this *Controller) List(context *gin.Context) {
    sendOk(context)
}

func (this *Controller) Start(context *gin.Context) {
    sendOk(context)
}

func (this *Controller) Delete(context *gin.Context) {
    sendOk(context)
}

func New(config *config.Config) *Controller {
    return &Controller{
        config: config,
        vm: vmModel.New(config),
    }
}
