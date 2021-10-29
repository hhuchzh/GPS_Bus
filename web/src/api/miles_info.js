import service from '@/utils/request'

export const getMilesInfoList = (params) => {
  return service({
    url: '/milesInfo/getMilesInfoList',
    method: 'get',
    params
  })
}
