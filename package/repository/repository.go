package repository

type Save interface {

}

type Repository struct {
	Save
}

func NewRepository() *Repository {
	return &Repository{}
}

	