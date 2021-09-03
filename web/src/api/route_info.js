import service from '@/utils/request'

// @Tags RouteInfo
// @Summary 创建RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RouteInfo true "创建RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /routeInfo/createRouteInfo [post]
export const createRouteInfo = (data) => {
  return service({
    url: '/routeInfo/createRouteInfo',
    method: 'post',
    data
  })
}

// @Tags RouteInfo
// @Summary 删除RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RouteInfo true "删除RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /routeInfo/deleteRouteInfo [delete]
export const deleteRouteInfo = (data) => {
  return service({
    url: '/routeInfo/deleteRouteInfo',
    method: 'delete',
    data
  })
}

// @Tags RouteInfo
// @Summary 删除RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /routeInfo/deleteRouteInfo [delete]
export const deleteRouteInfoByIds = (data) => {
  return service({
    url: '/routeInfo/deleteRouteInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags RouteInfo
// @Summary 更新RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RouteInfo true "更新RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /routeInfo/updateRouteInfo [put]
export const updateRouteInfo = (data) => {
  return service({
    url: '/routeInfo/updateRouteInfo',
    method: 'put',
    data
  })
}

// @Tags RouteInfo
// @Summary 用id查询RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RouteInfo true "用id查询RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /routeInfo/findRouteInfo [get]
export const findRouteInfo = (params) => {
  return service({
    url: '/routeInfo/findRouteInfo',
    method: 'get',
    params
  })
}

// @Tags RouteInfo
// @Summary 分页获取RouteInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RouteInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /routeInfo/getRouteInfoList [get]
export const getRouteInfoList = (params) => {
  return service({
    url: '/routeInfo/getRouteInfoList',
    method: 'get',
    params
  })
}
