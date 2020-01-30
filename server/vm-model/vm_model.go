/*
 * Copyright 2020 Oleg Borodin  <borodin@unix7.org>
 */

package vmModel

import (
    "sysrun/config"
    "log"
    "os/exec"
    "strings"
    "regexp"
    "strconv"
    "errors"
    "time"
)

type Model struct{
    config *config.Config
}

type VM struct {
    Name    string      `json:"name"`
    Id      int         `json:"id"`
    Status  string      `json:"status"`
}

func (this *Model) List() ([]VM, error) {

    out, err := exec.Command("qm", "list").Output()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(out), "\n")

    vms := []VM{}

    for _, line := range lines {

        rexp, err := regexp.Compile(`\s+`)
        if err != nil {
            log.Println(err)
            return nil, err

        }
        res := rexp.Split(line, -1)
        if len(res) < 3 {
            continue
        }

        id, _ := strconv.Atoi(res[1])
        if id == 0 {
            continue
        }

        vm := VM{
            Name : res[2],
            Id : id,
            Status : res[3],
        }
        vms = append(vms, vm)
    }
    return vms, err
}

func (this *Model) Stop(id int) error {
    out, err := exec.Command("qm", "stop", strconv.Itoa(id)).CombinedOutput()
    log.Println("stop vm id=", strconv.Itoa(id), string(out))
    if err != nil {
        return errors.New(string(out))
    }
    return nil
}

func (this *Model) DummyStop(id int) error {
    time.Sleep(time.Second * 5)
    log.Println("dummy stop vm id=", strconv.Itoa(id))
    return nil
}


func (this *Model) Shutdown(id int) error {
    out, err := exec.Command("qm", "shutdown", strconv.Itoa(id)).CombinedOutput()
    log.Println("stop vm id=", strconv.Itoa(id), string(out))
    if err != nil {
        return errors.New(string(out))
    }
    return nil
}

func (this *Model) Start(id int) error {
    out, err := exec.Command("qm", "start", strconv.Itoa(id)).CombinedOutput()
    log.Println("start vm id=", strconv.Itoa(id), string(out))
    if err != nil {
        return errors.New(string(out))
    }
    return nil
}

func New() *Model {
    return &Model{
    }
}
