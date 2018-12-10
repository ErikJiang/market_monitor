package service

import (
	"encoding/json"
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
func (ts *TaskService) QueryByID() (task *models.Task, err error) {
	taskModel := &models.Task{}
	condition := map[string]interface{}{
		"id": ts.TaskID,
	}
	task, err = taskModel.FindOne(condition)
	return
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

type TaskRuleParam struct {
	Operator string
	WarnPrice float64
}

func (ts *TaskService) QueryByPage(condition interface{}, page, pageSize int) ([]map[string]interface{}, int, error) {
	taskModel := &models.Task{}
	taskList, err := taskModel.Search(condition, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resList := make([]map[string]interface{}, len(taskList))
	for i, v := range taskList {
		log.Debug().Msgf("rule: %v, type: %T", v.Rules, v.Rules)
		rule := TaskRuleParam{}
		err := json.Unmarshal([]byte(v.Rules), &rule)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, 0, err
		}

		resList[i] = map[string]interface{}{
			"taskId": v.ID,
			"userId": v.User.ID,
			"email": v.User.Email,
			"taskType": v.Type,
			"status": v.Status,
			"operator": rule.Operator,
			"warnPrice": rule.WarnPrice,
		}
	}

	count, err := taskModel.Count(condition)
	if err != nil {
		return nil, 0, err
	}
	return resList, count, nil
}
