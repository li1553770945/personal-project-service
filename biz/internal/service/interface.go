package project

import (
	"context"
	"github.com/li1553770945/personal-projects-service/biz/internal/repo"
	"github.com/li1553770945/personal-projects-service/kitex_gen/projects"
)

type ProjectService struct {
	Repo repo.IRepository
}

type IProjectService interface {
	GetProjects(ctx context.Context, req *projects.ProjectsReq) (*projects.ProjectsResp, error)
}

func NewProjectService(repo repo.IRepository) IProjectService {
	return &ProjectService{
		Repo: repo,
	}
}
