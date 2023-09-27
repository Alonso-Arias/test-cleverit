package dao

import (
	"context"

	"github.com/Alonso-Arias/test-cleverit/db/base"
	"github.com/Alonso-Arias/test-cleverit/db/model"
	"github.com/Alonso-Arias/test-cleverit/log"
	"gorm.io/gorm"
)

var loggerf = log.LoggerJSON().WithField("package", "dao")

// TaskDAO - Task dao interface
type TaskDAO interface {
	FindAll(ctx context.Context) (model.Task, error)
	Get(ctx context.Context, sku string) (model.Task, error)
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, task model.Task) error
	Save(ctx context.Context, task model.Task) error
}

// TaskDAOImpl - Task dao implementation
type TaskDAOImpl struct {
}

// NewTaskDAO - gets an TaskDAOImpl instance
func NewTaskDAO() *TaskDAOImpl {
	return &TaskDAOImpl{}
}

// FindAll -
func (pd *TaskDAOImpl) FindAll(ctx context.Context) ([]model.Task, error) {

	log := loggerf.WithField("struct", "TaskDAOImpl").WithField("function", "FindAll")

	db := base.GetDB()

	tasks := []model.Task{}
	err := db.Find(&tasks).Error

	if err != nil {
		log.WithError(err).Error("get Tasks fails")
		return []model.Task{}, err
	} else if tasks == nil {
		return []model.Task{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", tasks)

	return tasks, nil

}

// FindAll -
func (pd *TaskDAOImpl) Get(ctx context.Context, id int32) (model.Task, error) {

	log := loggerf.WithField("struct", "TaskDAOImpl").WithField("function", "Get")

	db := base.GetDB()

	task := model.Task{}
	err := db.Where("ID = ?", id).FirstOrInit(&task).Error

	if err != nil {
		log.WithError(err).Error("get Tasks fails")
		return model.Task{}, err
	} else if task.Id == 0 {
		return model.Task{}, gorm.ErrRecordNotFound
	}

	log.Debugf("%v", task)

	return task, nil

}

// FindAll -
func (pd *TaskDAOImpl) Delete(ctx context.Context, id int32) error {

	log := loggerf.WithField("struct", "TaskDAOImpl").WithField("function", "Delete")

	db := base.GetDB()

	// inits tx
	err := db.Transaction(func(tx *gorm.DB) error {

		task := model.Task{}

		err := db.Where("ID = ?", id).Delete(&task).Error
		if err != nil {
			log.WithError(err).Error("problems with deleting Task")
			return err
		}

		return nil
	})

	if err != nil {
		log.WithError(err).Error("fails to delete order")
		return err
	}

	log.Infof("DEBUG : Deleted Sucessfull\n")

	return nil

}

func (pd *TaskDAOImpl) Update(ctx context.Context, task model.Task) error {

	log := loggerf.WithField("struct", "TaskDAOImpl").WithField("function", "Update")

	db := base.GetDB()

	tx := db.Model(&task).
		Where("ID = ?", task.Id).
		Updates(map[string]interface{}{
			"title":       gorm.Expr("IF(? = '', title, ?)", task.Title, task.Title),
			"description": gorm.Expr("IF(? = '', description, ?)", task.Description, task.Description),
			"due_date":    gorm.Expr("IF(? = '', due_date, ?)", task.DueDate, task.DueDate),
			"state":       gorm.Expr("IF(? = '', state, ?)", task.State, task.State),
		})

	if tx.Error != nil {
		log.Debugf("%v", tx.Error)
		return tx.Error
	}

	return nil
}

func (pd *TaskDAOImpl) Save(ctx context.Context, task model.Task) error {

	log := loggerf.WithField("struct", "TaskDAOImpl").WithField("function", "Save")

	db := base.GetDB()

	err := db.Create(&task)

	if err.Error != nil {
		log.Debugf("%v", err.Error)
		return err.Error
	}

	log.Infof("Save Task Sucessfull\n")

	return nil

}
