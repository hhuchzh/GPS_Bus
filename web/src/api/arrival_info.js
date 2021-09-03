import service from '@/utils/request'

// @Tags ArrivalInfo
// @Summary 创建ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArrivalInfo true "创建ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /arrivalInfo/createArrivalInfo [post]
export const createArrivalInfo = (data) => {
  return service({
    url: '/arrivalInfo/createArrivalInfo',
    method: 'post',
    data
  })
}

// @Tags ArrivalInfo
// @Summary 删除ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArrivalInfo true "删除ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /arrivalInfo/deleteArrivalInfo [delete]
export const deleteArrivalInfo = (data) => {
  return service({
    url: '/arrivalInfo/deleteArrivalInfo',
    method: 'delete',
    data
  })
}

// @Tags ArrivalInfo
// @Summary 删除ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /arrivalInfo/deleteArrivalInfo [delete]
export const deleteArrivalInfoByIds = (data) => {
  return service({
    url: '/arrivalInfo/deleteArrivalInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags ArrivalInfo
// @Summary 更新ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArrivalInfo true "更新ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /arrivalInfo/updateArrivalInfo [put]
export const updateArrivalInfo = (data) => {
  return service({
    url: '/arrivalInfo/updateArrivalInfo',
    method: 'put',
    data
  })
}

// @Tags ArrivalInfo
// @Summary 用id查询ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ArrivalInfo true "用id查询ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /arrivalInfo/findArrivalInfo [get]
export const findArrivalInfo = (params) => {
  return service({
    url: '/arrivalInfo/findArrivalInfo',
    method: 'get',
    params
  })
}

// @Tags ArrivalInfo
// @Summary 分页获取ArrivalInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ArrivalInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /arrivalInfo/getArrivalInfoList [get]
export const getArrivalInfoList = (params) => {
  return service({
    url: '/arrivalInfo/getArrivalInfoList',
    method: 'get',
    params
  })
}
