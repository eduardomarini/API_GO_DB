package models

import "time"

type Usuarios struct {
	IDUsuario     uint       `json:"IS_usuario" gorm:"column:id;primaryKey"`
	Nome          string     `json:"Nome" gorm:"column:nome"`
	Email         string     `json:"Email" gorm:"column:email"`
	Telefone      string     `json:"Telefone" gorm:"column:telefone"`
	Data_registro *time.Time `json:"Data_registro" gorm:"column:data_registro"`
}
