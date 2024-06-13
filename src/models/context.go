package models

import "github.com/treenq/treenq-cli/src/dto"

type Context struct {
	Name   string
	Url    string
	Active bool
	Info   dto.InfoResponse
}
