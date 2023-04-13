//go:build !js
// +build !js

package db

import "context"

import (
	"errors"
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/utils"
)

func (d *DataBase) InsertGroupRequest(ctx context.Context, groupRequest *model_struct.LocalGroupRequest) error {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	return utils.Wrap(d.conn.Create(groupRequest).Error, "InsertGroupRequest failed")
}
func (d *DataBase) DeleteGroupRequest(ctx context.Context, groupID, userID string) error {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	return utils.Wrap(d.conn.Where("group_id=? and user_id=?", groupID, userID).Delete(&model_struct.LocalGroupRequest{}).Error, "DeleteGroupRequest failed")
}
func (d *DataBase) UpdateGroupRequest(ctx context.Context, groupRequest *model_struct.LocalGroupRequest) error {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	t := d.conn.Model(groupRequest).Select("*").Updates(*groupRequest)
	if t.RowsAffected == 0 {
		return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
	}
	return utils.Wrap(t.Error, "")
}

func (d *DataBase) GetSendGroupApplication(ctx context.Context) ([]*model_struct.LocalGroupRequest, error) {
	d.groupMtx.Lock()
	defer d.groupMtx.Unlock()
	var groupRequestList []model_struct.LocalGroupRequest
	err := utils.Wrap(d.conn.Order("create_time DESC").Find(&groupRequestList).Error, "")
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var transfer []*model_struct.LocalGroupRequest
	for _, v := range groupRequestList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, nil
}
