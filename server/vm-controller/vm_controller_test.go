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
    "github.com/jmoiron/sqlx"
    "github.com/stretchr/testify/assert"

    "agent/server/pguser-model"
    "agent/config"
)

func TestNone(t *testing.T) {
}

