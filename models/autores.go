package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Autores struct {
	IDAutor         uint     `json:"ID_autor" gorm:"column:id;primaryKey"`
	Nome            string   `json:"Nome" gorm:"column:nome"`
	Nacionalidade   string   `json:"Nacionalidade" gorm:"column:nacionalidade"`
	Data_nascimento DateOnly `json:"Data_nascimento" gorm:"column:data_nascimento"`
}

type DateOnly struct {
	time.Time
}

const DateLayout = "2006-01-02"

// MarshalJSON implementa a serialização no formato "YYYY-MM-DD"
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(DateLayout) + `"`), nil
}

// UnmarshalJSON implementa a desserialização no formato "YYYY-MM-DD"
func (d *DateOnly) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(`"`+DateLayout+`"`, string(data))
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// Valuer implementa a interface do GORM para persistir no banco
func (d DateOnly) Value() (driver.Value, error) {
	// Retorna a data no formato adequado para o banco (por exemplo, '2006-01-02')
	return d.Time.Format(DateLayout), nil
}

// Scanner implementa a interface do GORM para recuperar dados do banco
func (d *DateOnly) Scan(value interface{}) error {
	// Converte o valor recebido (geralmente uma string) para time.Time
	if v, ok := value.(string); ok {
		parsed, err := time.Parse(DateLayout, v)
		if err != nil {
			return err
		}
		d.Time = parsed
		return nil
	}
	return fmt.Errorf("could not scan type %T into DateOnly", value)
}
