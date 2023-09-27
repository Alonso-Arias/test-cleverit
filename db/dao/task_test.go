package dao

import (
	"context"
	"testing"
	"time"

	"github.com/apex/log"

	"github.com/Alonso-Arias/test-cleverit/db/model"
	"github.com/stretchr/testify/assert"
)

var TaskDao = NewTaskDAO()

func TestSave_OK(t *testing.T) {

	format := "2006-01-02T15:04:05"

	dateNow := time.Now()

	dateFormated, err := time.Parse(format, dateNow.Format("2006-01-02T15:04:05"))
	if err != nil {
		log.WithError(err).Error("Binding error")
		assert.FailNowf(t, "fails", "fails to gets exam: %v", err)
	}

	err = TaskDao.Save(context.TODO(), model.Task{Id: 999, Title: "Test", Description: "Test", DueDate: dateFormated, State: ""})

	if err != nil {
		assert.FailNowf(t, "fails", "fails to update Task: %v", err)
	}

}

func TestFindAll_OK(t *testing.T) {

	result, err := TaskDao.FindAll(context.TODO())

	if err != nil {
		assert.FailNowf(t, "fails", "fails to gets Tasks: %v", err)
	}

	t.Logf("Result : %v", result)

}

func TestGetBySku_OK(t *testing.T) {

	result, err := TaskDao.Get(context.TODO(), 999)

	if err != nil {
		assert.FailNowf(t, "fails", "fails to gets Tasks: %v", err)
	}

	t.Logf("Result : %v", result)

}

func TestUpdate_OK(t *testing.T) {

	err := TaskDao.Update(context.TODO(), model.Task{Id: 999, Title: "Testing2"})

	if err != nil {
		assert.FailNowf(t, "fails", "fails to update Task: %v", err)
	}

}
