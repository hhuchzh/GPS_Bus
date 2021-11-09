import service from '@/utils/request'

// @Tags GpsInfo
// @Summary 创建GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GpsInfo true "创建GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpsInfo/createGpsInfo [post]
export const createGpsInfo = (data) => {
  return service({
    url: '/gpsInfo/createGpsInfo',
    method: 'post',
    data
  })
}

// @Tags GpsInfo
// @Summary 删除GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GpsInfo true "删除GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gpsInfo/deleteGpsInfo [delete]
export const deleteGpsInfo = (data) => {
  return service({
    url: '/gpsInfo/deleteGpsInfo',
    method: 'delete',
    data
  })
}

// @Tags GpsInfo
// @Summary 删除GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gpsInfo/deleteGpsInfo [delete]
export const deleteGpsInfoByIds = (data) => {
  return service({
    url: '/gpsInfo/deleteGpsInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags GpsInfo
// @Summary 更新GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GpsInfo true "更新GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gpsInfo/updateGpsInfo [put]
export const updateGpsInfo = (data) => {
  return service({
    url: '/gpsInfo/updateGpsInfo',
    method: 'put',
    data
  })
}

// @Tags GpsInfo
// @Summary 用id查询GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GpsInfo true "用id查询GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gpsInfo/findGpsInfo [get]
export const findGpsInfo = (params) => {
  return service({
    url: '/gpsInfo/findGpsInfo',
    method: 'get',
    params
  })
}

// @Tags GpsInfo
// @Summary 分页获取GpsInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GpsInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpsInfo/getGpsInfoList [get]
export const getGpsInfoList = (params) => {
  return service({
    url: '/gpsInfo/getGpsInfoList',
    method: 'get',
    params
  })
}

// @Tags GpsInfo
// @Summary 分页获取GpsInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GpsInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpsInfo/getAvailableGpsInfoList [get]
export const getAvailableGpsInfoList = (params) => {
  return service({
    url: '/gpsInfo/getAvailableGpsInfoList',
    method: 'get',
    params
  })
}

export const getNotAvailableGpsInfoList = (params) => {
  return service({
    url: '/gpsInfo/getNotAvailableGpsInfoList',
    method: 'get',
    params
  })
}
