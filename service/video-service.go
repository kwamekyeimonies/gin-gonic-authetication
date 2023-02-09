package service

import "github.com/kwamekyeimonies/gin-gonic-framework/entity"

type Video_Service interface{
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type video_service struct{
	videos []entity.Video
}

func New() Video_Service{
	return &video_service{}
}

func (service *video_service) Save(video entity.Video) entity.Video{
	
	service.videos = append(service.videos,video )

	return video
}

func (service * video_service) FindAll() []entity.Video{

	return service.videos

}