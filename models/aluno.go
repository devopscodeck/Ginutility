package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model        //cria no banco alguns campos necessarios como id timestamp etc..
	Nome       string `json:"nome"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
}
