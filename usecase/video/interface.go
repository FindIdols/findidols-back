package video

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Video, error)
	GetVideos(id entity.ID) ([]*entity.Video, error)
}

//Writer idol writer
type Writer interface {
	// UploadVideo(e *entity.Video) (entity.ID, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	Get(
		idolID entity.ID,
	) (*entity.Video, error)
	GetVideos(id entity.ID) ([]*entity.Video, error)
}
