/*
* gin-vue-admin web框架组
*
* */
// 加载网站配置文件夹
import { register } from './global'

export const run = function(app) {
  register(app)
  console.log(`
     欢迎使用 研创园车辆管理系统 
     项目合作微信：hhuchzh
  `)
}

