package main

import (
	"context"
	"github.com/li1553770945/personal-projects-service/biz/infra/container"
	projects "github.com/li1553770945/personal-projects-service/kitex_gen/projects"
)

// ProjectsServiceImpl implements the last service interface defined in the IDL.
type ProjectsServiceImpl struct{}

// GetProjects implements the ProjectsServiceImpl interface.
func (s *ProjectsServiceImpl) GetProjects(ctx context.Context, req *projects.ProjectsReq) (resp *projects.ProjectsResp, err error) {
	// TODO: Your code here...
	App := container.GetGlobalContainer()
	resp, err = App.ProjectService.GetProjects(ctx, req)
	return
}

// GetProjectNum implements the ProjectsServiceImpl interface.
func (s *ProjectsServiceImpl) GetProjectNum(ctx context.Context) (resp *projects.ProjectNumResp, err error) {
	// TODO: Your code here...
	return
}
