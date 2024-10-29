package models

import "time"

type Autores struct {
	IDAutor         uint       `json:"ID_autor" gorm:"column:id;primaryKey"`
	Nome            string     `json:"Nome" gorm:"column:nome"`
	Nacionalidade   string     `json:"Nacionalidade" gorm:"column:nacionalidade"`
	Data_nascimento *time.Time `json:"Data_nascimento" gorm:"column:data_nascimento"`
}
