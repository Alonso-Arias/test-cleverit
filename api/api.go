package main

import (
	"context"
	"net/http"
	"strconv"

	errs "github.com/Alonso-Arias/test-cleverit/errors"
	"github.com/Alonso-Arias/test-cleverit/log"
	"github.com/Alonso-Arias/test-cleverit/services/task"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var loggerf = log.LoggerJSON().WithField("package", "main")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.POST("/api/v1/task", taskPost)
	e.GET("/api/v1/task/findAll", findAllTasksGet)
	e.GET("/api/v1/task/:id", taskGet)
	e.PUT("/api/v1/task", taskPut)
	e.DELETE("/api/v1/task/:id", taskDelete)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))

}

// PermissionValidator - filters users and validates if they have permissions for execute the API.
// func PermissionValidator(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		au, _ := security.AuthenticatedUserFromClaims(c)
// 		if security.IsAuthorized(au, c.Request().Method, c.Request().RequestURI) {
// 			return next(c)
// 		}
// 		return c.String(http.StatusUnauthorized, "Without privileges for this function")
// 	}
// }

// find all tasks
// @Summary Find all tasks
// @tags tasks
// @Description obtiene todos los task
// @ID findAlltasksGet
// @Accept  json
// @Produce  json
// @Success 200  {object} task.FindAllTasksResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /tasks/findAll [get]
func findAllTasksGet(c echo.Context) error {

	res, err := task.TaskService{}.FindAllTasks(context.TODO())
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// get task
// @Summary get task by id
// @tags task
// @Description obtiene task por id
// @ID taskGet
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200  {object} task.GetTaskResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /task/{id} [get]
func taskGet(c echo.Context) error {

	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := task.GetTaskRequest{
		Id: int32(idInt),
	}

	res, err := task.TaskService{}.GetTask(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// delete task
// @Summary delete task by id
// @tags task
// @Description elimina un task
// @ID taskDelete
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200  {object} task.DeleteTaskResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /task/{id} [delete]
func taskDelete(c echo.Context) error {

	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := task.DeleteTaskRequest{
		Id: int32(idInt),
	}

	res, err := task.TaskService{}.DeleteTask(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// update task
// @Summary update task by id
// @tags task
// @Description actualiza un task
// @ID taskPut
// @Accept  json
// @Produce  json
// @Param UpdatetaskRequest body task.UpdateTaskRequest true "task"
// @Success 200  {object} task.UpdateTaskResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /task [put]
func taskPut(c echo.Context) error {

	log := loggerf.WithField("func", "taskPut")

	req := task.UpdateTaskRequest{}

	if err := c.Bind(req); err != nil {
		log.WithError(err).Error("Binding error")
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := task.TaskService{}.UpdateTask(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// save task
// @Summary save task
// @tags task
// @Description guarda un task
// @ID taskPost
// @Accept  json
// @Produce  json
// @Param SavetaskRequest body task.SaveTaskRequest true "task"
// @Success 200  {object} task.SaveTaskResponse
// @Failure 404 {object}  errors.CustomError
// @Failure 500 {object}  errors.CustomError
// @Router /task [post]
func taskPost(c echo.Context) error {

	log := loggerf.WithField("func", "taskPost")

	req := task.SaveTaskRequest{}

	if err := c.Bind(&req); err != nil {
		log.WithError(err).Error("Binding error")
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := task.TaskService{}.SaveTask(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
