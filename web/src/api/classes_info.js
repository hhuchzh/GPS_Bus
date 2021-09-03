import service from '@/utils/request'

// @Tags ClassesInfo
// @Summary 创建ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassesInfo true "创建ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classesInfo/createClassesInfo [post]
export const createClassesInfo = (data) => {
  return service({
    url: '/classesInfo/createClassesInfo',
    method: 'post',
    data
  })
}

// @Tags ClassesInfo
// @Summary 删除ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassesInfo true "删除ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classesInfo/deleteClassesInfo [delete]
export const deleteClassesInfo = (data) => {
  return service({
    url: '/classesInfo/deleteClassesInfo',
    method: 'delete',
    data
  })
}

// @Tags ClassesInfo
// @Summary 删除ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classesInfo/deleteClassesInfo [delete]
export const deleteClassesInfoByIds = (data) => {
  return service({
    url: '/classesInfo/deleteClassesInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags ClassesInfo
// @Summary 更新ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ClassesInfo true "更新ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /classesInfo/updateClassesInfo [put]
export const updateClassesInfo = (data) => {
  return service({
    url: '/classesInfo/updateClassesInfo',
    method: 'put',
    data
  })
}

// @Tags ClassesInfo
// @Summary 用id查询ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ClassesInfo true "用id查询ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /classesInfo/findClassesInfo [get]
export const findClassesInfo = (params) => {
  return service({
    url: '/classesInfo/findClassesInfo',
    method: 'get',
    params
  })
}

// @Tags ClassesInfo
// @Summary 分页获取ClassesInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ClassesInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classesInfo/getClassesInfoList [get]
export const getClassesInfoList = (params) => {
  return service({
    url: '/classesInfo/getClassesInfoList',
    method: 'get',
    params
  })
}
