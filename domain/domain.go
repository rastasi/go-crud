package domain

import (
	"github.com/rastasi/go-crud/domain/repository"
	"github.com/rastasi/go-crud/domain/service"
	"github.com/rastasi/go-crud/lib/database"
)

type Domain struct {
	AlbumService  service.AlbumService
	ArtistService service.ArtistService
}

func NewDomain() Domain {
	DB := database.Init()
	return Domain{
		AlbumService: service.AlbumService{
			AlbumRepository: repository.AlbumRepository{
				DB: DB,
			},
		},
		ArtistService: service.ArtistService{
			ArtistRepository: repository.ArtistRepository{
				DB: DB,
			},
		},
	}
}
