package metadata

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"

	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/das/pkg/resp"
)

const (
	dbNameStruct  = "DbName"
	ownerIdStruct = "OwnerId"
	envIdStruct   = "EnvId"
)

// @Tags database
// @Summary get all databases
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "owner_group": "2,3,4", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [get]
func GetDB(c *gin.Context) {
	// init service
	s := metadata.NewDbServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetDbAll, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataGetDbAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetDbAll)
}

// @Tags database
// @Summary get database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "owner_group": "2,3,4", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/:id [get]
func GetDBByID(c *gin.Context) {
	// get param
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
		return
	}
	// init service
	s := metadata.NewDbServiceWithDefault()
	// get entity
	err := s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataGetDbByID, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, id, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataGetDbByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataGetDbByID, id)
}

// @Tags database
// @Summary add a new database
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "owner_group": "2,3,4", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db [post]
func AddDB(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.DbInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, ok := fields[dbNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbNameStruct)
		return
	}
	// init service
	s := metadata.NewDbServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataAddDb, fmt.Sprintf("%s and %s and %s", dbNameStruct, ownerIdStruct, envIdStruct), err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, fmt.Sprintf("%s and %s and %s", dbNameStruct, ownerIdStruct, envIdStruct), err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataAddDb, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.InfoMetadataAddDb, fmt.Sprintf("%s and %s and %s", dbNameStruct, ownerIdStruct, envIdStruct))
}

// @Tags database
// @Summary update database by id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": [{"id": 1, "db_name": "db1", "cluster_id": 1, "cluster_type": 1, "owner_id": 1, "owner_group": "2,3,4", "env_id": 1, "del_flag": 0, "create_time": "2021-01-22T09:59:21.379851+08:00", "last_update_time": "2021-01-22T09:59:21.379851+08:00"}]}"
// @Router /api/v1/metadata/db/:id [post]
func UpdateDBByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	id := c.Param(idJSON)
	if id == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, idJSON)
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.DbInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	_, dbNameExists := fields[dbNameStruct]
	_, ownerIdExists := fields[ownerIdStruct]
	_, envIdExists := fields[envIdStruct]
	_, delFlagExists := fields[delFlagStruct]
	if !dbNameExists && !ownerIdExists && !envIdExists && !delFlagExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s", dbNameStruct, ownerIdStruct, envIdStruct, delFlagStruct))
		return
	}
	// init service
	s := metadata.NewDbServiceWithDefault()
	// update entity
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMetadataUpdateDb, id, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalService, id, err.Error())
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(message.DebugMetadataUpdateDb, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, message.DebugMetadataUpdateDb, id)
}
