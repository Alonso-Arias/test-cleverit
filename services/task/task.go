package task

import (
	"context"
	"time"

	"github.com/Alonso-Arias/test-cleverit/db/dao"
	md "github.com/Alonso-Arias/test-cleverit/db/model"
	errs "github.com/Alonso-Arias/test-cleverit/errors"
	"github.com/Alonso-Arias/test-cleverit/log"
	"github.com/Alonso-Arias/test-cleverit/services/enums"
	"github.com/Alonso-Arias/test-cleverit/services/model"
	"gopkg.in/dealancer/validate.v2"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "services")

var format = "2006-01-02T15:04:05"

// TaskService contiene los métodos relacionados con las tareas.
type TaskService struct{}

// FindAllTasksResponse es la respuesta para FindAllTasks.
type FindAllTasksResponse struct {
	Tasks []model.Task `json:"tasks"`
}

// FindAllTasks recupera todas las tareas.
func (ts TaskService) FindAllTasks(ctx context.Context) (FindAllTasksResponse, error) {
	log := loggerf.WithField("service", "TaskService").WithField("func", "FindAllTasks")

	taskDAO := dao.NewTaskDAO()

	tasks, err := taskDAO.FindAll(ctx)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting tasks")
		return FindAllTasksResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return FindAllTasksResponse{}, errs.TasksNotFound
	}

	results := []model.Task{}

	for _, v := range tasks {
		// Convierte la fecha de time.Time a una cadena en el formato especificado.
		dueDateStr := v.DueDate.Format(format)
		task := model.Task{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			DueDate:     dueDateStr,
			State:       v.State,
		}
		results = append(results, task)
	}

	return FindAllTasksResponse{Tasks: results}, nil
}

// GetTaskRequest es la solicitud para GetTask.
type GetTaskRequest struct {
	Id int32 `json:"id"`
}

// GetTaskResponse es la respuesta para GetTask.
type GetTaskResponse struct {
	Task model.Task `json:"task"`
}

// GetTask obtiene una tarea por su ID.
func (ts TaskService) GetTask(ctx context.Context, in GetTaskRequest) (GetTaskResponse, error) {
	log := loggerf.WithField("service", "TaskService").WithField("func", "GetTask")

	if in.Id == 0 {
		return GetTaskResponse{}, errs.BadRequest
	}

	taskDAO := dao.NewTaskDAO()

	v, err := taskDAO.Get(ctx, in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting task")
		return GetTaskResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return GetTaskResponse{}, errs.TasksNotFound
	}

	// Convierte la fecha de time.Time a una cadena en el formato especificado.
	dueDateStr := v.DueDate.Format(format)

	task := model.Task{
		Id:          v.Id,
		Title:       v.Title,
		Description: v.Description,
		DueDate:     dueDateStr,
		State:       v.State,
	}

	return GetTaskResponse{Task: task}, nil
}

// DeleteTaskRequest es la solicitud para DeleteTask.
type DeleteTaskRequest struct {
	Id int32 `json:"id"`
}

// DeleteTaskResponse es la respuesta para DeleteTask.
type DeleteTaskResponse struct{}

// DeleteTask elimina una tarea por su ID.
func (ts TaskService) DeleteTask(ctx context.Context, in DeleteTaskRequest) (DeleteTaskResponse, error) {
	log := loggerf.WithField("service", "TaskService").WithField("func", "DeleteTask")

	if in.Id == 0 {
		return DeleteTaskResponse{}, errs.BadRequest
	}

	taskDAO := dao.NewTaskDAO()

	_, err := taskDAO.Get(ctx, in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting task")
		return DeleteTaskResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return DeleteTaskResponse{}, errs.TasksNotFound
	}

	err = taskDAO.Delete(ctx, in.Id)
	if err != nil {
		return DeleteTaskResponse{}, err
	}

	return DeleteTaskResponse{}, nil
}

// UpdateTaskRequest es la solicitud para UpdateTask.
type UpdateTaskRequest struct {
	Task model.Task `json:"task"`
}

// UpdateTaskResponse es la respuesta para UpdateTask.
type UpdateTaskResponse struct{}

// UpdateTask actualiza una tarea.
func (ts TaskService) UpdateTask(ctx context.Context, in UpdateTaskRequest) (UpdateTaskResponse, error) {
	log := loggerf.WithField("service", "TaskService").WithField("func", "UpdateTask")

	// Valida la solicitud de entrada
	if err := validate.Validate(in); err != nil {
		log.WithError(err).Error("validation problems")
		return UpdateTaskResponse{}, errs.BadRequest
	}

	taskDAO := dao.NewTaskDAO()

	_, err := taskDAO.Get(ctx, in.Task.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("problems with getting task")
		return UpdateTaskResponse{}, err
	} else if err == gorm.ErrRecordNotFound {
		return UpdateTaskResponse{}, errs.TasksNotFound
	}

	dateFormatted, err := time.Parse(format, in.Task.DueDate)
	if err != nil {
		log.WithError(err).Error("binding error")
		return UpdateTaskResponse{}, errs.BadRequest
	}

	err = taskDAO.Update(ctx, md.Task(md.Task{
		Id:          in.Task.Id,
		Title:       in.Task.Title,
		Description: in.Task.Description,
		DueDate:     dateFormatted,
		State:       in.Task.State,
	}))
	if err != nil {
		return UpdateTaskResponse{}, err
	}

	return UpdateTaskResponse{}, nil
}

// SaveTaskRequest es la solicitud para SaveTask.
type SaveTaskRequest struct {
	Task model.Task `json:"task"`
}

// SaveTaskResponse es la respuesta para SaveTask.
type SaveTaskResponse struct{}

// SaveTask guarda una nueva tarea.
func (ts TaskService) SaveTask(ctx context.Context, in SaveTaskRequest) (SaveTaskResponse, error) {
	log := loggerf.WithField("service", "TaskService").WithField("func", "SaveTask")

	err := stateValidate(ctx, in)
	if err != nil {
		return SaveTaskResponse{}, err
	}
	// Valida la solicitud de entrada
	if err := validate.Validate(in); err != nil {
		log.WithError(err).Error("validation problems")
		return SaveTaskResponse{}, errs.BadRequest
	}

	taskDAO := dao.NewTaskDAO()

	dateFormatted, err := time.Parse(format, in.Task.DueDate)
	if err != nil {
		log.WithError(err).Error("binding error")
		return SaveTaskResponse{}, errs.BadRequest
	}

	err = taskDAO.Save(ctx, md.Task(md.Task{
		Title:       in.Task.Title,
		Description: in.Task.Description,
		DueDate:     dateFormatted,
		State:       in.Task.State,
	}))
	if err != nil {
		return SaveTaskResponse{}, err
	}

	return SaveTaskResponse{}, nil
}

func stateValidate(ctx context.Context, in SaveTaskRequest) error {

	var flag bool
	taskStatuses := []string{
		enums.PendingTaskStatus,
		enums.InProgressTaskStatus,
		enums.CompletedTaskStatus,
	}

	for _, v := range taskStatuses {
		if in.Task.State != v {
			flag = true
		}
	}

	if flag {
		return errs.TaskStateInvalid
	}

	return nil

}
