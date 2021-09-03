import service from '@/utils/request'

// @Tags CheckinInfo
// @Summary 创建CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckinInfo true "创建CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkinInfo/createCheckinInfo [post]
export const createCheckinInfo = (data) => {
  return service({
    url: '/checkinInfo/createCheckinInfo',
    method: 'post',
    data
  })
}

// @Tags CheckinInfo
// @Summary 删除CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckinInfo true "删除CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkinInfo/deleteCheckinInfo [delete]
export const deleteCheckinInfo = (data) => {
  return service({
    url: '/checkinInfo/deleteCheckinInfo',
    method: 'delete',
    data
  })
}

// @Tags CheckinInfo
// @Summary 删除CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkinInfo/deleteCheckinInfo [delete]
export const deleteCheckinInfoByIds = (data) => {
  return service({
    url: '/checkinInfo/deleteCheckinInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags CheckinInfo
// @Summary 更新CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckinInfo true "更新CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkinInfo/updateCheckinInfo [put]
export const updateCheckinInfo = (data) => {
  return service({
    url: '/checkinInfo/updateCheckinInfo',
    method: 'put',
    data
  })
}

// @Tags CheckinInfo
// @Summary 用id查询CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CheckinInfo true "用id查询CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkinInfo/findCheckinInfo [get]
export const findCheckinInfo = (params) => {
  return service({
    url: '/checkinInfo/findCheckinInfo',
    method: 'get',
    params
  })
}

// @Tags CheckinInfo
// @Summary 分页获取CheckinInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CheckinInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkinInfo/getCheckinInfoList [get]
export const getCheckinInfoList = (params) => {
  return service({
    url: '/checkinInfo/getCheckinInfoList',
    method: 'get',
    params
  })
}
