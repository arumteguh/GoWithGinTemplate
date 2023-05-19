package albums

type AlbumRepository interface {
	GetAlbums() ([]album, error)
}

type DefaultAlbumRepository struct {
	db string
}

func NewDefaultAlbumRepository(db string) AlbumRepository {
	return &DefaultAlbumRepository{
		db: db,
	}
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (ar *DefaultAlbumRepository) GetAlbums() ([]album, error) {
	return albums, nil
}
