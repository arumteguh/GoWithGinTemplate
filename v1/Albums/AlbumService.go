package albums

type AlbumService interface {
	GetAlbums() ([]album, error)
}

type DefaultAlbumService struct {
	repo AlbumRepository
}

func NewDefaultAlbumService(repo AlbumRepository) AlbumService {
	return &DefaultAlbumService{
		repo: repo,
	}
}

func (as *DefaultAlbumService) GetAlbums() ([]album, error) {
	return as.repo.GetAlbums()
}
