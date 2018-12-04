package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Task 任务表 model 定义
type Task struct {
	gorm.Model
	User    User    `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID  int     `gorm:"column:userId;not null"`
	Type    string  `sql:"type:ENUM('TICKER', 'OTHER')"`
	Status  string  `sql:"type:ENUM('ENABLE', 'DISABLE')"`
	Rules    string  `gorm:"column:rules;type:varchar(200);not null"`
}

// Insert 新增任务
func (task *Task) Insert() (taskID uint, err error) {

	result := DB.Create(&task)
	taskID = task.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询任务信息
func (task *Task) FindOne(condition map[string]interface{}) (*Task, error) {
	var taskInfo Task
	result := DB.Select("id, name, email, avatar, password").Where(condition).First(&taskInfo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if taskInfo.ID > 0 {
		return &taskInfo, nil
	}
	return nil, nil
}

// UpdateOne 修改任务
func (task *Task) UpdateOne(taskID uint, data map[string]interface{}) (*Task, error) {
	err := DB.Model(&Task{}).Where("id = ?", taskID).Updates(data).Error
	if err != nil {
		return nil, err
	}
	var updTask Task
	err = DB.Select([]string{"id", "userId", "type", "status", "rules"}).First(&updTask, taskID).Error
	if err != nil {
		return nil, err
	}
	return &updTask, nil
}

// DeleteOne 删除任务
func (task *Task) DeleteOne(taskID uint) error {
	if err := DB.Select([]string{"id"}).First(&task, taskID).Error; err != nil {
		return err
	}
	if err := DB.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

// Search 分页数据查询
func (task *Task) Search(query interface{}, page int, pageSize int) ([]*Task, error) {
	var tasks []*Task
	err := DB.Offset(pageSize * (page - 1)).Limit(pageSize).Where(query).Find(&tasks).Error
	return tasks, errors.WithStack(err)
}

// Count 分页总数查询
func (task *Task) Count(query interface{}) (int, error) {
	var count int
	err := DB.Model(&Task{}).Where(query).Count(&count).Error
	return count, errors.WithStack(err)
}