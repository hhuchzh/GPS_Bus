<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="车牌号:" prop="busId">
          <el-select v-model="currentSelectedPlate" placeholder="请选择车牌号" @change="changeOption($event)">
            <el-option
              v-for="item in busInfoList"
              :key="item.busPlate"
              :label="item.busPlate"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="日期:" prop="time">
          <el-date-picker v-model="currentSelectedDate" placeholder="请选择日期" clearable autocomplete="off" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit('morning')">早高峰查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit('afternoon')">晚高峰查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div id="container" style="width: 1600px; height: 750px;" />
  </div>
</template>

<script>
// import BMapGL from 'BMapGL'
import {
  findBusInfo,
  getBusInfoList
} from '@/api/bus_info' //  此处请自行替换地址
import {
  track
} from '@/api/gpsQuery' //  此处请自行替换地址
export default {
  data() {
    return {
      map: {},
      busInfoList: [],
      currentSelectedID: -1,
      currentSelectedPlate: '',
      currentSelectGPSSN: '',
      currentSelectedDate: this.getNowFormatDate(),
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
      currentLabel: null
    }
  },
  async created() {
    this.getBusInfos()
  },
  mounted() {
    this.loadJScript()
  },
  methods: {
    async getBusInfos() {
      const res = await getBusInfoList()
      if (res.code === 0) {
        this.busInfoList = res.data.list
        // console.log(this.busInfoList.length)
        if (this.busInfoList && this.busInfoList.length > 0) {
          this.currentSelectedPlate = this.busInfoList[0].busPlate
          this.currentSelectedID = this.busInfoList[0].ID
          this.getBusInfo()
        } else {
          this.currentSelectedPlate = ''
          this.currentSelectedID = -1
        }
      }
    },
    changeOption(event) {
      this.currentSelectedID = event
      this.getBusInfo()
    },
    async getBusInfo() {
      const res = await findBusInfo({ ID: this.currentSelectedID })
      // console.log(res)
      // console.log(res.data)
      if (res.code === 0) {
        if (res.data.rebusInfo && res.data.rebusInfo.gpsInfos[0]) {
          this.currentSelectGPSSN = res.data.rebusInfo.gpsInfos[0].gpsSn
          // console.log('getBusInfo : ' + this.currentSelectGPSSN)
        }
      }
    },
    getNowFormatDate() {
      var date = new Date()
      var seperator1 = '-'
      var year = date.getFullYear()
      var month = date.getMonth() + 1
      var strDate = date.getDate()
      if (month >= 1 && month <= 9) {
        month = '0' + month
      }
      if (strDate >= 0 && strDate <= 9) {
        strDate = '0' + strDate
      }
      var currentdate = year + seperator1 + month + seperator1 + strDate
      return currentdate
    },
    loadJScript() {
      var script = document.createElement('script')
      script.type = 'text/javascript'
      script.src = '//api.map.baidu.com/getscript?v=2.0&ak=qG1aM2UK80lGHVHKmHczRcmzBBOd8blK'
      script.onload = () => {
        console.log('script load success')
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
      this.map.addOverlay(mkr) // 标点
      this.preMarker = mkr
      this.points.push(this.arrPoints[this.pointIdx])
      this.centerPoint = this.arrPoints[this.pointIdx]
      this.setZoom()
      this.driveline(this.points)
      this.pointIdx++
      setTimeout(this.dynamicLine, 500)
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
      this.carIcon_0 = new BMap.Icon('car_0.png', new BMap.Size(48, 48))
      this.carIcon_1 = new BMap.Icon('car_1.png', new BMap.Size(48, 48))
      this.carIcon_2 = new BMap.Icon('car_2.png', new BMap.Size(48, 48))
      this.carIcon_3 = new BMap.Icon('car_3.png', new BMap.Size(48, 48))
      this.carIcon_4 = new BMap.Icon('car_4.png', new BMap.Size(48, 48))
      this.carIcon_5 = new BMap.Icon('car_5.png', new BMap.Size(48, 48))
      this.carIcon_6 = new BMap.Icon('car_6.png', new BMap.Size(48, 48))
      this.carIcon_7 = new BMap.Icon('car_7.png', new BMap.Size(48, 48))

      this.currentLabel = new window.BMap.Label('速度:' + this.pointList[this.pointIdx].speed + ' KM/H', {
        offset: new window.BMap.Size(20, -10)
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
    },
    // 条件搜索前端看此方法
    async onSubmit(time) {
      this.initParameters()
      var beginTime = ''
      var endTime = ''
      if (time === 'morning') {
        beginTime = this.currentSelectedDate + ' 07:00:00'
        endTime = this.currentSelectedDate + ' 11:00:00'
      } else if (time === 'afternoon') {
        beginTime = this.currentSelectedDate + ' 15:00:00'
        endTime = this.currentSelectedDate + ' 18:00:00'
      }
      // console.log('beginTime : ' + beginTime + '/ endTime : ' + endTime)
      const ret = await track({ gpsSn: this.currentSelectGPSSN, beginTime: beginTime, endTime: endTime })
      if (ret.code === 0 && ret.data.track.length > 0) {
        // console.log(ret.data.track)
        this.showHistoryPath(ret.data.track)
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
