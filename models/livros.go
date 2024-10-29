package models

type Livros struct {
	IDLivro       uint   `json:"ID_livro" gorm:"column:id;primaryKey"`
	Titulo        string `json:"Titulo" gorm:"column:titulo"`
	IDAutor       uint   `json:"ID_autor" gorm:"column:id_autor;foreignKey"`
	IDCategoria   uint   `json:"ID_categoria" gorm:"column:id_categoria;foreignKey"`
	AnoPublicacao uint   `json:"Ano_publicacao" gorm:"column:ano_publicacao"`
	Isbn          string `json:"Isbn" gorm:"column:isbn"`
}
