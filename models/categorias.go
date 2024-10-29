package models

type Categorias struct {
	IDCategoria uint   `json:"ID_categoria" gorm:"column:id;primaryKey"`
	Descricao   string `json:"Descricao" gorm:"column:descricao"`
}
