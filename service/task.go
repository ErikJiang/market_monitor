package service

import (
	"github.com/JiangInk/market_monitor/extend/code"
	"github.com/JiangInk/market_monitor/models"
	"github.com/rs/zerolog/log"
)

// UserService 用户服务层逻辑
type TaskService struct{
	TaskID  uint
	UserID  int
	Type    string
	Rules   string
}

// QueryByID 通过任务ID查询任务信息
func (ts *TaskService) QueryByID() (*models.Task, error) {
	taskModel := &models.Task{}
	condition := map[string]interface{}{
		"id": ts.TaskID,
	}
	return taskModel.FindOne(condition)
}

// StoreUser 添加用户
func (ts *TaskService) StoreTask() (taskID uint, err error) {
	log.Info().Msg("enter StoreTask service")

	task := &models.Task{
		UserID: ts.UserID,
		Type: ts.Type,
		Rules: ts.Rules,
		Status: "ENABLE",
	}
	taskID, err = task.Insert()
	return
}

// 更新任务信息
func (ts *TaskService) UpdateInfo(taskID uint) (*models.Task, *code.Code) {
	taskModel := &models.Task{}
	updateTask, err := taskModel.UpdateOne(taskID, map[string]interface{}{
		"type": ts.Type,
		"rules": ts.Rules,
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, code.ServiceInsideError
	}
	return updateTask, nil
}

func (ts *TaskService) RemoveTask(taskID uint) error {
	taskModel := &models.Task {}
	err := taskModel.DeleteOne(taskID)
	if err != nil {
		return err
	}
	return nil
}
