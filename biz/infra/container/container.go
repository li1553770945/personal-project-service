package container

import (
	"github.com/li1553770945/personal-project-service/biz/infra/config"
	"github.com/li1553770945/personal-project-service/biz/internal/service"
	"sync"
)

type Container struct {
	Config         *config.Config
	ProjectService project.IProjectService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(cfg *config.Config) {
	once.Do(func() {
		APP = GetContainer(cfg)
	})
}

func NewContainer(config *config.Config, projectService project.IProjectService,
) *Container {
	return &Container{
		Config:         config,
		ProjectService: projectService,
	}

}
