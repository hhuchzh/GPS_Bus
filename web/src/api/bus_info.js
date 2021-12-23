import service from '@/utils/request'

// @Tags BusInfo
// @Summary 创建BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusInfo true "创建BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busInfo/createBusInfo [post]
export const createBusInfo = (data) => {
  return service({
    url: '/busInfo/createBusInfo',
    method: 'post',
    data
  })
}

// @Tags BusInfo
// @Summary 删除BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusInfo true "删除BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busInfo/deleteBusInfo [delete]
export const deleteBusInfo = (data) => {
  return service({
    url: '/busInfo/deleteBusInfo',
    method: 'delete',
    data
  })
}

// @Tags BusInfo
// @Summary 删除BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busInfo/deleteBusInfo [delete]
export const deleteBusInfoByIds = (data) => {
  return service({
    url: '/busInfo/deleteBusInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags BusInfo
// @Summary 更新BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusInfo true "更新BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busInfo/updateBusInfo [put]
export const updateBusInfo = (data) => {
  return service({
    url: '/busInfo/updateBusInfo',
    method: 'put',
    data
  })
}

// @Tags BusInfo
// @Summary 用id查询BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BusInfo true "用id查询BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busInfo/findBusInfo [get]
export const findBusInfo = (params) => {
  return service({
    url: '/busInfo/findBusInfo',
    method: 'get',
    params
  })
}

// @Tags BusInfo
// @Summary 分页获取BusInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BusInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busInfo/getBusInfoList [get]
export const getBusInfoList = (params) => {
  return service({
    url: '/busInfo/getBusInfoList',
    method: 'get',
    params
  })
}
