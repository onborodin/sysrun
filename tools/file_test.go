/*
 * Copyright 2020 Oleg Borodin  <borodin@unix7.org>
 */


package tools

import (
    "testing"
    //"github.com/stretchr/testify/assert"
    //"fmt"
    "strings"
    "errors"

    "github.com/GehirnInc/crypt"
    _ "github.com/GehirnInc/crypt/sha256_crypt"
)

func TestCreateHash(t *testing.T) {
    hash, err := createHash("123456")
    if err != nil {
        t.Error(err)
    }
    if len(hash) != 59 {
        t.Error(err)
    }
}

//$5$rfBd67ti3SMt$V4Gi.A19xLBWMSZDFw4JvfotiQdF.nDmMsFB3DyPvI0

func CheckHash(hash, password string) error {
    arr := strings.Split(hash, "$")
    if len(arr) < 3 {
        return errors.New("incorrect hash structure")
    }
    hashType := arr[1]
    hashSalt := arr[2]
    if hashType != "5" {
        return errors.New("incorrect hash type")
    }
    crypt := crypt.SHA256.New()
    newHash, _ := crypt.Generate([]byte(password), []byte("$" + hashType + "$" + hashSalt))

    if hash != newHash {
        return errors.New("password is incorrect")
    }
    return nil
}

func TestCheckHash(t *testing.T) {
    hash := "$5$rfBd67ti3SMt$V4Gi.A19xLBWMSZDFw4JvfotiQdF.nDmMsFB3DyPvI0"
    password := "123456"

    err := CheckHash(hash, password)
    if err != nil {
        t.Error(err)
    }
}
