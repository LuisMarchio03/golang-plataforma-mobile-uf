package handler

type Handler struct {
    User               *UserHandler
    Project            *ProjectHandler
    ProjectApplication *ProjectApplicationHandler
    Task               *TaskHandler
}

func NewHandler(user *UserHandler, project *ProjectHandler, projectApplication *ProjectApplicationHandler, task *TaskHandler) *Handler {
    return &Handler{
        User:               user,
        Project:            project,
        ProjectApplication: projectApplication,
        Task:               task,
    }
}