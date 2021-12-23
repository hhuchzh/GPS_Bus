import service from '@/utils/request'

// @Tags LocationCommon
// @Summary 创建LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocationCommon true "创建LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationCommon/createLocationCommon [post]
export const createLocationCommon = (data) => {
  return service({
    url: '/locationCommon/createLocationCommon',
    method: 'post',
    data
  })
}

// @Tags LocationCommon
// @Summary 删除LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocationCommon true "删除LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /locationCommon/deleteLocationCommon [delete]
export const deleteLocationCommon = (data) => {
  return service({
    url: '/locationCommon/deleteLocationCommon',
    method: 'delete',
    data
  })
}

// @Tags LocationCommon
// @Summary 删除LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /locationCommon/deleteLocationCommon [delete]
export const deleteLocationCommonByIds = (data) => {
  return service({
    url: '/locationCommon/deleteLocationCommonByIds',
    method: 'delete',
    data
  })
}

// @Tags LocationCommon
// @Summary 更新LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocationCommon true "更新LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /locationCommon/updateLocationCommon [put]
export const updateLocationCommon = (data) => {
  return service({
    url: '/locationCommon/updateLocationCommon',
    method: 'put',
    data
  })
}

// @Tags LocationCommon
// @Summary 用id查询LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LocationCommon true "用id查询LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /locationCommon/findLocationCommon [get]
export const findLocationCommon = (params) => {
  return service({
    url: '/locationCommon/findLocationCommon',
    method: 'get',
    params
  })
}

// @Tags LocationCommon
// @Summary 分页获取LocationCommon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LocationCommon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationCommon/getLocationCommonList [get]
export const getLocationCommonList = (params) => {
  return service({
    url: '/locationCommon/getLocationCommonList',
    method: 'get',
    params
  })
}
