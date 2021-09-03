import service from '@/utils/request'

// @Tags ShiftInfo
// @Summary 创建ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShiftInfo true "创建ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiftInfo/createShiftInfo [post]
export const createShiftInfo = (data) => {
  return service({
    url: '/shiftInfo/createShiftInfo',
    method: 'post',
    data
  })
}

// @Tags ShiftInfo
// @Summary 删除ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShiftInfo true "删除ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shiftInfo/deleteShiftInfo [delete]
export const deleteShiftInfo = (data) => {
  return service({
    url: '/shiftInfo/deleteShiftInfo',
    method: 'delete',
    data
  })
}

// @Tags ShiftInfo
// @Summary 删除ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shiftInfo/deleteShiftInfo [delete]
export const deleteShiftInfoByIds = (data) => {
  return service({
    url: '/shiftInfo/deleteShiftInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags ShiftInfo
// @Summary 更新ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShiftInfo true "更新ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shiftInfo/updateShiftInfo [put]
export const updateShiftInfo = (data) => {
  return service({
    url: '/shiftInfo/updateShiftInfo',
    method: 'put',
    data
  })
}

// @Tags ShiftInfo
// @Summary 用id查询ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShiftInfo true "用id查询ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shiftInfo/findShiftInfo [get]
export const findShiftInfo = (params) => {
  return service({
    url: '/shiftInfo/findShiftInfo',
    method: 'get',
    params
  })
}

// @Tags ShiftInfo
// @Summary 分页获取ShiftInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ShiftInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiftInfo/getShiftInfoList [get]
export const getShiftInfoList = (params) => {
  return service({
    url: '/shiftInfo/getShiftInfoList',
    method: 'get',
    params
  })
}
