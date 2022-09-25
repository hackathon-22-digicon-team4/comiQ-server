package repository

type Repository interface {
	UserRepository
	AuthorRepository
	BookSeriesRepository
	BookRepository
	StampRepository
}
