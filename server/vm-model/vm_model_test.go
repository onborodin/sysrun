/*
 * Copyright 2020 Oleg Borodin  <borodin@unix7.org>
 */


package vmModel

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
    vm := New()
    vms, err := vm.List()
    if err != nil {
        t.Error(err)
    }
    assert.NotEqual(t, nil, vms, "Receive null VM list")
}

func TestStart(t *testing.T) {
    vm := New()
    err := vm.Start(12345)
    if err != nil {
        t.Error(err)
    }
}

func TestStop(t *testing.T) {
    vm := New()
    err := vm.Stop(12345)
    if err != nil {
        t.Error(err)
    }
}
