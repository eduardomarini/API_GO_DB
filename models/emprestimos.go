package models

type Emprestimos struct {
	IDEmprestimo   uint     `json:"ID_emprestimo" gorm:"column:id;primaryKey"`
	IDUsuario      uint     `json:"ID_usuario" gorm:"column:id_usuario;foreignKey"`
	IDLivro        uint     `json:"ID_livro" gorm:"column:id_livro;foreignKey"`
	DataEmprestimo DateOnly `json:"Data_emprestimo" gorm:"column:data_emprestimo"`
	DataDevolucao  DateOnly `json:"Data_devolucao" gorm:"column:data_devolucao"`
}
