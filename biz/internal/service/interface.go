package project

import (
	"context"
	"github.com/li1553770945/personal-project-service/biz/internal/repo"
	"github.com/li1553770945/personal-project-service/kitex_gen/project"
)

type ProjectService struct {
	Repo repo.IRepository
}

type IProjectService interface {
	GetProjects(ctx context.Context, req *project.ProjectsReq) (*project.ProjectsResp, error)
	GetProjectNum(ctx context.Context) (*project.ProjectNumResp, error)
}

func NewProjectService(repo repo.IRepository) IProjectService {
	return &ProjectService{
		Repo: repo,
	}
}
