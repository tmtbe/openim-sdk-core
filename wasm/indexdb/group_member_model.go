//go:build js && wasm
// +build js,wasm

package indexdb

import "context"

import (
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/utils"
)

type LocalGroupMember struct {
}

func NewLocalGroupMember() *LocalGroupMember {
	return &LocalGroupMember{}
}

func (i *LocalGroupMember) GetGroupMemberInfoByGroupIDUserID(ctx context.Context, groupID, userID string) (*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID, userID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return &temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetAllGroupMemberList(ctx context.Context) ([]model_struct.LocalGroupMember, error) {
	member, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetAllGroupMemberUserIDList(ctx context.Context) ([]model_struct.LocalGroupMember, error) {
	member, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupMemberCount(ctx context.Context, groupID string) (uint32, error) {
	count, err := Exec(groupID)
	if err != nil {
		return 0, err
	}
	if v, ok := count.(float64); ok {
		return uint32(v), nil
	}
	return 0, ErrType
}

func (i *LocalGroupMember) GetGroupSomeMemberInfo(ctx context.Context, groupID string, userIDList []string) ([]*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID, utils.StructToJsonString(userIDList))
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupAdminID(ctx context.Context, groupID string) ([]string, error) {
	IDList, err := Exec(groupID)
	if err != nil {
		return nil, err
	}
	if v, ok := IDList.(string); ok {
		var temp []string
		err := utils.JsonStringToStruct(v, &temp)
		if err != nil {
			return nil, err
		}
		return temp, nil
	}
	return nil, ErrType
}

func (i *LocalGroupMember) GetGroupMemberListByGroupID(ctx context.Context, groupID string) ([]*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupMemberListSplit(ctx context.Context, groupID string, filter int32, offset, count int) ([]*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID, filter, offset, count)
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupMemberOwnerAndAdmin(ctx context.Context, groupID string) ([]*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupMemberOwner(ctx context.Context, groupID string) (*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return &temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupMemberListSplitByJoinTimeFilter(ctx context.Context, groupID string, offset, count int, joinTimeBegin, joinTimeEnd int64, userIDList []string) ([]*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID, offset, count, joinTimeBegin, joinTimeEnd, utils.StructToJsonString(userIDList))
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupOwnerAndAdminByGroupID(ctx context.Context, groupID string) ([]*model_struct.LocalGroupMember, error) {
	member, err := Exec(groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) GetGroupMemberUIDListByGroupID(ctx context.Context, groupID string) (result []string, err error) {
	IDList, err := Exec(groupID)
	if err != nil {
		return nil, err
	}
	if v, ok := IDList.(string); ok {
		err := utils.JsonStringToStruct(v, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, ErrType
}

func (i *LocalGroupMember) InsertGroupMember(ctx context.Context, groupMember *model_struct.LocalGroupMember) error {
	_, err := Exec(utils.StructToJsonString(groupMember))
	return err
}

func (i *LocalGroupMember) BatchInsertGroupMember(ctx context.Context, groupMemberList []*model_struct.LocalGroupMember) error {
	_, err := Exec(utils.StructToJsonString(groupMemberList))
	return err
}

func (i *LocalGroupMember) DeleteGroupMember(ctx context.Context, groupID, userID string) error {
	_, err := Exec(groupID, userID)
	return err
}

func (i *LocalGroupMember) DeleteGroupAllMembers(ctx context.Context, groupID string) error {
	_, err := Exec(groupID)
	return err
}

func (i *LocalGroupMember) UpdateGroupMember(ctx context.Context, groupMember *model_struct.LocalGroupMember) error {
	_, err := Exec(utils.StructToJsonString(groupMember))
	return err
}

func (i *LocalGroupMember) UpdateGroupMemberField(ctx context.Context, groupID, userID string, args map[string]interface{}) error {
	_, err := Exec(groupID, userID, utils.StructToJsonString(args))
	return err
}

func (i *LocalGroupMember) GetGroupMemberInfoIfOwnerOrAdmin(ctx context.Context) ([]*model_struct.LocalGroupMember, error) {
	member, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalGroupMember) SearchGroupMembersDB(ctx context.Context, keyword string, groupID string, isSearchMemberNickname, isSearchUserID bool, offset, count int) (result []*model_struct.LocalGroupMember, err error) {
	member, err := Exec(keyword, groupID, isSearchMemberNickname, isSearchUserID, offset, count)
	if err != nil {
		return nil, err
	} else {
		if v, ok := member.(string); ok {
			var temp []*model_struct.LocalGroupMember
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			return temp, err
		} else {
			return nil, ErrType
		}
	}
}
