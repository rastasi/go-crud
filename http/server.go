package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/go-crud/domain"
	"github.com/rastasi/go-crud/http/controller"
	"github.com/rastasi/go-crud/http/router"
	"github.com/rastasi/go-crud/lib/utils"
)

func StartHttpServer(domain domain.Domain) {
	r := mux.NewRouter()
	r.StrictSlash(false)
	utils.AddRouter(r, "/albums", *router.AlbumRouter{
		AlbumController: controller.AlbumController{
			AlbumService: domain.AlbumService,
		},
	}.Init())
	utils.AddRouter(r, "/artists", *router.ArtistRouter{
		ArtistController: controller.ArtistController{
			ArtistService: domain.ArtistService,
		},
	}.Init())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
