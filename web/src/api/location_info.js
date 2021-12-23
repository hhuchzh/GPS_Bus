import service from '@/utils/request'

// @Tags LocationInfo
// @Summary 创建LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocationInfo true "创建LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationInfo/createLocationInfo [post]
export const createLocationInfo = (data) => {
  return service({
    url: '/locationInfo/createLocationInfo',
    method: 'post',
    data
  })
}

// @Tags LocationInfo
// @Summary 删除LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocationInfo true "删除LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /locationInfo/deleteLocationInfo [delete]
export const deleteLocationInfo = (data) => {
  return service({
    url: '/locationInfo/deleteLocationInfo',
    method: 'delete',
    data
  })
}

// @Tags LocationInfo
// @Summary 删除LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /locationInfo/deleteLocationInfo [delete]
export const deleteLocationInfoByIds = (data) => {
  return service({
    url: '/locationInfo/deleteLocationInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags LocationInfo
// @Summary 更新LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocationInfo true "更新LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /locationInfo/updateLocationInfo [put]
export const updateLocationInfo = (data) => {
  return service({
    url: '/locationInfo/updateLocationInfo',
    method: 'put',
    data
  })
}

// @Tags LocationInfo
// @Summary 用id查询LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LocationInfo true "用id查询LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /locationInfo/findLocationInfo [get]
export const findLocationInfo = (params) => {
  return service({
    url: '/locationInfo/findLocationInfo',
    method: 'get',
    params
  })
}

// @Tags LocationInfo
// @Summary 分页获取LocationInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LocationInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationInfo/getLocationInfoList [get]
export const getLocationInfoList = (params) => {
  return service({
    url: '/locationInfo/getLocationInfoList',
    method: 'get',
    params
  })
}
