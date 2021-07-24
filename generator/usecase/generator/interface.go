package generator

import entity "Referral-System/generator/entity"

//Writer Generator writer
type Writer interface {
	Create(e *entity.Generator) (err error)
}

//Repository interface
type Repository interface {
	Writer
}

//UseCase interface
type UseCase interface {
	CreateGenerator(ID, Name, Email, Password string) (GeneratedLink string, err error)
}
