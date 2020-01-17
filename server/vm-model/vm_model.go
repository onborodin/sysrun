/*
 * Copyright 2020 Oleg Borodin  <borodin@unix7.org>
 */

package vmModel

import (
    "sysrun/config"
)

type Model struct{
    config *config.Config
}

func (this *Model) List() {
}

func (this *Model) Stop() {
}

func (this *Model) Start() {
}

func New(config *config.Config) *Model {
    return &Model{
        config: config,
    }
}
