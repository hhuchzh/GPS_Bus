<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="路线:" prop="routeId">
          <el-select v-model="currnetSelectRouteName" placeholder="请选择路线" @change="changeRouteListOption($event)">
            <el-option
              v-for="item in routeList"
              :key="item.routeName"
              :label="item.routeName"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="班次:" prop="classId">
          <el-select v-model="currnetSelectClassName" placeholder="请选择班次" @change="changeClassListOption($event)">
            <el-option
              v-for="item in classList"
              :key="item.remark"
              :label="item.remark"
              :value="item.busId"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="开始时间:" prop="startTime">
          <el-date-picker v-model="currentSelectedStartDate" placeholder="请选择日期" type="datetime" :disabled-date="disabledDate" value-format="YYYY-MM-DD HH:mm:ss" />
        </el-form-item>
        <el-form-item label="结束时间:" prop="EndTime">
          <el-date-picker v-model="currentSelectedEndDate" placeholder="请选择日期" type="datetime" :disabled-date="disabledDate" value-format="YYYY-MM-DD HH:mm:ss" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit()">历史轨迹查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div id="container" style="width: 1600px; height: 750px;" />
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" width="20%" top="15%">
      <el-form>
        <p style="text-align:center">没有历史轨迹数据，请重新查询</p>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="mini" type="primary" @click="closeDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
// import BMapGL from 'BMapGL'
import {
  getRouteInfoList
} from '@/api/route_info' //  此处请自行替换地址
import {
  getClassesInfoList
} from '@/api/classes_info'
import {
  track
} from '@/api/gpsQuery' //  此处请自行替换地址
import {
  findBusInfo
} from '@/api/bus_info' //  此处请自行替换地址\
import {
  getArrivalInfoList
} from '@/api/arrival_info'
export default {
  data() {
    return {
      map: {},
      currentSelectGPSSN: '',
      currentSelectedStartDate: this.getNowDate(),
      currentSelectedEndDate: this.getNowDate(),
      preMarker: null,
      pointIdx: 0,
      pointList: [],
      arrPoints: [],
      points: [],
      centerPoint: null,
      currentBound: null,
      carIcon_0: null,
      carIcon_1: null,
      carIcon_2: null,
      carIcon_3: null,
      carIcon_4: null,
      carIcon_5: null,
      carIcon_6: null,
      carIcon_7: null,
      currentLabel: null,
      timerId: null,
      dialogFormVisible: false,
      currnetSelectRouteId: -1,
      currnetSelectRouteName: '',
      routeList: [],
      currnetSelectBusId: -1,
      currnetSelectClassName: '',
      classList: [],
      currentSelectClassId: -1,
      locationPointsList: [],
      prePoint: null
    }
  },
  async created() {
    this.getRouteInfo()
  },
  mounted() {
    this.loadJScript()
  },
  methods: {
    async getRouteInfo() {
      const res = await getRouteInfoList()
      if (res.code === 0) {
        console.log(res)
        this.routeList = res.data.list
        if (this.routeList && this.routeList.length > 0) {
          this.currnetSelectRouteName = this.routeList[0].routeName
          this.currnetSelectRouteId = this.routeList[0].ID
          this.getClassInfo(this.currnetSelectRouteId)
        } else {
          this.currnetSelectRouteName = ''
          this.currnetSelectRouteId = -1
        }
      }
    },
    async getClassInfo(routeId) {
      const res = await getClassesInfoList({ routeId: parseInt(routeId) })
      if (res.code === 0) {
        this.classList = res.data.list
        if (this.classList && this.classList.length > 0) {
          this.currnetSelectClassName = this.classList[0].remark
          this.currnetSelectBusId = this.classList[0].busId
          this.currentSelectClassId = this.classList[0].ID
          this.getBusInfo()
          // this.getArrivalInfoList()
        } else {
          this.currnetSelectClassName = ''
          this.currnetSelectBusId = -1
        }
      }
    },
    changeRouteListOption(event) {
      console.log(event)
      this.getClassInfo(event)
    },
    changeClassListOption(event) {
      // TBD
      console.log(event)
      this.currnetSelectBusId = event
      this.getBusInfo()
    },
    async getBusInfo() {
      const res = await findBusInfo({ ID: this.currnetSelectBusId })
      // console.log(res)
      // console.log(res.data)
      if (res.code === 0) {
        if (res.data.rebusInfo && res.data.rebusInfo.gpsInfos[0]) {
          this.currentSelectGPSSN = res.data.rebusInfo.gpsInfos[0].gpsSn
          console.log(this.currentSelectGPSSN)
        }
      }
    },
    disabledDate(date) {
      // alert('date')
      return date.getTime() > Date.now() - 24 * 60 * 60 * 1000
    },
    getNowDate() {
      var myDate = new Date()
      var year = myDate.getFullYear()
      var mon = myDate.getMonth() + 1
      var date = myDate.getDate()
      var hours = myDate.getHours()
      var minutes = myDate.getMinutes()
      var seconds = myDate.getSeconds()
      if (mon >= 1 && mon <= 9) {
        mon = '0' + mon
      }
      if (date >= 0 && date <= 9) {
        date = '0' + date
      }
      if (hours >= 0 && hours <= 9) {
        hours = '0' + hours
      }
      if (minutes >= 0 && minutes <= 9) {
        minutes = '0' + minutes
      }
      if (seconds >= 0 && seconds <= 9) {
        seconds = '0' + seconds
      }
      var now = year + '-' + mon + '-' + date + ' ' + hours + ':' + minutes + ':' + seconds
      return now
    },
    loadJScript() {
      var script = document.createElement('script')
      script.type = 'text/javascript'
      script.src = '//api.map.baidu.com/getscript?v=2.0&ak=qG1aM2UK80lGHVHKmHczRcmzBBOd8blK'
      script.onload = () => {
        console.log('script load success')
        var BMap = window.BMap
        this.carIcon_0 = new BMap.Icon('./images/car_0.png', new BMap.Size(48, 48))
        this.carIcon_1 = new BMap.Icon('./images/car_1.png', new BMap.Size(48, 48))
        this.carIcon_2 = new BMap.Icon('./images/car_2.png', new BMap.Size(48, 48))
        this.carIcon_3 = new BMap.Icon('./images/car_3.png', new BMap.Size(48, 48))
        this.carIcon_4 = new BMap.Icon('./images/car_4.png', new BMap.Size(48, 48))
        this.carIcon_5 = new BMap.Icon('./images/car_5.png', new BMap.Size(48, 48))
        this.carIcon_6 = new BMap.Icon('./images/car_6.png', new BMap.Size(48, 48))
        this.carIcon_7 = new BMap.Icon('./images/car_7.png', new BMap.Size(48, 48))

        this.currentLabel = new BMap.Label('速度:0 KM/H', {
          offset: new BMap.Size(20, -10)
        })
        this.currentLabel.setStyle({
          color: '#fff',
          backgroundColor: '#333333',
          border: '0',
          fontSize: '14px',
          width: '100px',
          // height: '20px',
          opacity: '0.8',
          verticalAlign: 'center',
          textAlign: 'center',
          borderRadius: '2px',
          whiteSpace: 'normal',
          wordWrap: 'break-word',
          padding: '5px',
        })
      }
      document.body.appendChild(script)

      var scriptWebGl = document.createElement('script')
      scriptWebGl.type = 'text/javascript'
      scriptWebGl.src = '//api.map.baidu.com/getscript?type=webgl&v=2.0&ak=qG1aM2UK80lGHVHKmHczRcmzBBOd8blK'
      scriptWebGl.onload = () => {
        console.log('scriptWebGl load success')
      }
      document.body.appendChild(scriptWebGl)

      var trackscript = document.createElement('script')
      trackscript.type = 'text/javascript'
      trackscript.src = '//api.map.baidu.com/library/TrackAnimation/src/TrackAnimation_min.js'
      trackscript.onload = () => {
        console.log('trackscript load success')
      }
      document.body.appendChild(trackscript)
    },
    transferPoint(trackList) {
      var preItemLng = 0.0
      var preItemLat = 0.0
      trackList.forEach((item) => {
        if (preItemLng !== item.lng && preItemLat !== item.lat) {
        // if (item.speed > 0) {
          this.arrPoints.push(new window.BMap.Point(item.lng, item.lat))
          this.pointList.push(item)
          preItemLng = item.lng
          preItemLat = item.lat
        }
      })
    },
    // 画折线
    driveline(points) {
      this.map.addOverlay(new window.BMap.Polyline(points, {
        strokeColor: '#00ff00',
        strokeWeight: 4,
        strokeOpacity: 1
      }))
    },
    driveline1(startPoint, endPoint) {
      var polyline
      var map = this.map
      var options = {
        onSearchComplete: function(results) {
          if (driving.getStatus() === 0) {
            // 获取第一条方案
            var plan = results.getPlan(0)
            // 获取方案的驾车线路
            var route = plan.getRoute(0)
            var points = route.getPath()
            polyline = new window.BMap.Polyline(points, {
              strokeColor: '#00ff00',
              strokeWeight: 4,
              strokeOpacity: 1
            })
            map.addOverlay(polyline)
          }
        }
      }
      var driving = new window.BMap.DrivingRoute(this.map, options)
      driving.search(startPoint, endPoint)
    },
    // 根据点信息实时更新地图显示范围，让轨迹完整显示。设置新的中心点和显示级别
    setZoom() {
      if (this.currentBound.containsPoint(this.centerPoint) === false) {
        this.map.centerAndZoom(this.centerPoint, 15)
        this.currentBound = this.map.getBounds()
      }
    },
    dynamicLine() {
      if (this.pointIdx === this.arrPoints.length) {
        var view = this.map.getViewport(this.points)
        var centerPoint = view.center
        this.map.centerAndZoom(centerPoint, 13)
        return
      }
      /* var label = new window.BMap.Label(this.arrPoints[this.pointIdx].lng, {
        offset: new window.BMap.Size(20, -10)
      })
      mkr.setLabel(label)*/
      if (this.preMarker != null) {
        this.map.removeOverlay(this.preMarker)
      }
      var mkr = null
      var dir = Number(this.pointList[this.pointIdx].dir)
      if (dir === 0 || dir === 360) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_0 })
      } else if (dir > 0 && dir < 90) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_1 })
      } else if (dir === 90) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_2 })
      } else if (dir > 90 && dir < 180) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_3 })
      } else if (dir === 180) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_4 })
      } else if (dir > 180 && dir < 270) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_5 })
      } else if (dir === 270) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_6 })
      } else if (dir > 270 && dir < 360) {
        mkr = new window.BMap.Marker(this.arrPoints[this.pointIdx], { icon: this.carIcon_7 })
      }
      this.currentLabel.setContent('速度:' + this.pointList[this.pointIdx].speed + ' KM/H')
      mkr.setLabel(this.currentLabel)
      this.points.push(this.arrPoints[this.pointIdx])
      this.centerPoint = this.arrPoints[this.pointIdx]
      this.setZoom()
      this.driveline(this.points)
      /* if (this.prePoint === null) {
        this.prePoint = this.arrPoints[this.pointIdx]
      }
      this.driveline(this.prePoint, this.arrPoints[this.pointIdx])
      this.prePoint = this.arrPoints[this.pointIdx]
      */
      this.map.addOverlay(mkr) // 标点
      this.preMarker = mkr
      this.pointIdx++
      if (this.timerId != null) {
        clearTimeout(this.timerId)
      }
      this.timerId = setTimeout(this.dynamicLine, 500)
    },
    showHistoryPath(trackList) {
      this.transferPoint(trackList)
      var BMap = window.BMap
      this.map = new BMap.Map('container') // 创建Map实例
      this.map.centerAndZoom(this.arrPoints[0], 15) // 初始化地图,设置中心点坐标和地图级别
      this.map.setCurrentCity('南京') // 设置地图中心显示的城市 new！
      this.map.enableScrollWheelZoom(true) // 开启鼠标滚轮缩放
      this.map.addControl(new BMap.NavigationControl()) // 缩放按钮
      this.currentBound = this.map.getBounds()
      this.getArrivalInfoList()
      this.dynamicLine()
    },
    initParameters() {
      this.pointIdx = 0
      this.arrPoints = []
      this.points = []
      this.map = {}
      this.currentBound = null
      this.preMarker = null
      this.pointList = []
      if (this.timerId != null) {
        clearTimeout(this.timerId)
      }
      var el = document.getElementById('container')
      if (el) {
        el.style.display = 'block'
      }
    },
    removeMap() {
      var el = document.getElementById('container')
      if (el) {
        el.innerHTML = ''
        el.style.display = 'none'
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
    },
    // 条件搜索前端看此方法
    async onSubmit(time) {
      this.initParameters()
      console.log('gps : ' + this.currentSelectGPSSN + '/ startTime : ' + this.currentSelectedStartDate + '/ endTime : ' + this.currentSelectedEndDate)
      const ret = await track({ gpsSn: this.currentSelectGPSSN, beginTime: this.currentSelectedStartDate, endTime: this.currentSelectedEndDate })
      if (ret.code === 0 && ret.data.track.length > 0) {
        // console.log(ret.data.track)
        this.showHistoryPath(ret.data.track)
      } else {
        console.log('no data')
        this.removeMap()
        this.dialogFormVisible = true
      }
    },
    async getArrivalInfoList() {
      var data = await getArrivalInfoList({ classesId: this.currentSelectClassId })
      console.log(data)
      if (data.code === 0) {
        this.map.clearOverlays()
        this.locationPointsList = []
        for (var i = 0; i < data.data.list.length; i++) {
          if (data.data.list[i].Location) {
            this.locationPointsList.push(new window.BMap.Point(data.data.list[i].Location.longtitude, data.data.list[i].Location.latitude))
            var locationMarker = new window.BMap.Marker(this.locationPointsList[i])
            this.map.addOverlay(locationMarker)
          }
        }
        var polyline
        var map = this.map
        var options = {
          onSearchComplete: function(results) {
            if (driving.getStatus() === 0) {
              // 获取第一条方案
              var plan = results.getPlan(0)
              // 获取方案的驾车线路
              var route = plan.getRoute(0)
              var points = route.getPath()
              polyline = new window.BMap.Polyline(points, {
                strokeColor: 'red',
                strokeWeight: 2,
                strokeOpacity: 1,
                strokeStyle: 'dashed'
              })
              map.addOverlay(polyline)
            }
          }
        }
        var driving = new window.BMap.DrivingRoute(this.map, options)
        for (var j = 0; j < this.locationPointsList.length - 1; j++) {
          driving.search(this.locationPointsList[j], this.locationPointsList[j + 1])
        }
      }
    },
  },
}
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>
