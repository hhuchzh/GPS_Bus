import service from '@/utils/request'

export const getLocation = (params) => {
  return service({
    url: '/gps/location',
    method: 'get',
    params
  })
}

export const track = (params) => {
  return service({
    url: '/gps/track',
    method: 'get',
    params
  })
}

export const locationlist = (params) => {
  return service({
    url: '/gps/locationlist',
    method: 'get',
    params
  })
}

