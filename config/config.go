/*
 * Copyright 2019 Oleg Borodin  <borodin@unix7.org>
 */

package config

import (
    "io/ioutil"
    "path/filepath"
    "github.com/go-yaml/yaml"
    "os"
)

type Config struct {
    ConfigPath          string  `yaml:"-"`
    LibDir              string  `yaml:"-"`
    PasswordPath        string  `yaml:"passwords"`
    PidPath             string  `yaml:"pidfile"`
    MessageLogPath      string  `yaml:"messagelog"`
    AccessLogPath       string  `yaml:"accesslog"`
    Port                int     `yaml:"port"`
    Debug               bool    `yaml:"debug"`
    Devel               bool    `yaml:"-"`
    StoreDir            string  `yaml:"storedir"`
    User                string  `yaml:"user"`
    Group               string  `yaml:"group"`
    CertPath            string  `yaml:"cert"`
    KeyPath             string  `yaml:"key"`

    DbHost              string  `yaml:"dbhost"`
    DbPort              int     `yaml:"dbport"`
    DbUser              string  `yaml:"dbuser"`
    DbPass              string  `yaml:"dbpass"`
}

//func (this Config) ResolveConfigPath() (string, error) {
//    fullPath, err := filepath.Abs(this.ConfigFile)
//    return fullPath, err
//}

func (this Config) Write() error {
    fileName, _ := filepath.Abs(this.ConfigPath)
    os.Rename(fileName, fileName + "~")

    var data []byte
    var err error
    if data, err = yaml.Marshal(this); err != nil {
        return err
    }
    return ioutil.WriteFile(fileName, data, 0640)
}

func (this *Config) Read() error {
    fileName, _ := filepath.Abs(this.ConfigPath)

    var data []byte
    var err error
    if data, err = ioutil.ReadFile(fileName); err != nil {
        return err
    }
    return yaml.Unmarshal(data, &this)
}

func (this *Config) GetStoreDir() (string, error) {
    return filepath.Abs(this.StoreDir)
}

func New() *Config {
    return &Config{
        ConfigPath:     "/home/ziggi/sysrun//syssrv.yml",
        LibDir:         "/home/ziggi/sysrun/",
        PasswordPath:   "/home/ziggi/sysrun//syspwd.db",
        PidPath:        "/home/ziggi/sysrun/run/syssrv.pid",
        MessageLogPath: "/home/ziggi/sysrun/log/message.log",
        AccessLogPath:  "/home/ziggi/sysrun/log/access.log",
        Port:           8080,
        Debug:          false,
        Devel:          false,
        StoreDir:       "/home/ziggi/sysrun/data",
        User:           "ziggi",
        Group:          "wheel",
        CertPath:       "/home/ziggi/sysrun//syssrv.crt",
        KeyPath:        "/home/ziggi/sysrun//syssrv.key",

        DbHost:         "localhost",
        DbPort:         5432,
        DbUser:         "postgres",
        DbPass:         "password",
    }
}
