package config

import "log"

type App struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
