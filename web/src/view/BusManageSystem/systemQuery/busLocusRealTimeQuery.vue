<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="车牌号:" prop="busId">
          <el-select v-model="currentSelectedPlate" placeholder="请选择" @change="changeOption($event)">
            <el-option
              v-for="item in busInfoList"
              :key="item.busPlate"
              :label="item.busPlate"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">实时位置查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div id="container" style="width: 1600px; height: 750px;" />
  </div>
</template>

<script>
// const BMap = {}
import {
  findBusInfo,
  getBusInfoList
} from '@/api/bus_info' //  此处请自行替换地址
import {
  getLocation,
  locationlist
} from '@/api/gpsQuery' //  此处请自行替换地址
import {
  getNotAvailableGpsInfoList
} from '@/api/gps_info'

export default {
  data() {
    return {
      map: {},
      busInfoList: [],
      currentSelectedID: -1,
      currentSelectedPlate: '',
      currentSelectGPSSN: '',
      preMarker: null,
      carIcon_0: null,
      carIcon_1: null,
      carIcon_2: null,
      carIcon_3: null,
      carIcon_4: null,
      carIcon_5: null,
      carIcon_6: null,
      carIcon_7: null,
      carIcon_static_0: null,
      carIcon_static_1: null,
      carIcon_static_2: null,
      carIcon_static_3: null,
      carIcon_static_4: null,
      carIcon_static_5: null,
      carIcon_static_6: null,
      carIcon_static_7: null,
      currentBound: null,
      currentLabel: null,
      timerId: null,
      gpsInfoList: [],
      viewSize: 18,
      queryType: 'single',
      currentGpsList: []
    }
  },
  async created() {
    this.getBusInfos()
  },
  mounted() {
    this.loadJScript()
  },
  methods: {
    async getBusGPSList(type) {
      this.gpsInfoList = []
      var res
      if (type === 'all') {
        res = await getNotAvailableGpsInfoList()
        if (res.code === 0) {
          if (res.data.list.length > 0) {
            this.gpsInfoList = res.data.list
          }
        }
      } else if (type === 'single') {
        res = await findBusInfo({ ID: this.currentSelectedID })
        if (res.code === 0) {
          if (res.data.rebusInfo && res.data.rebusInfo.gpsInfos[0]) {
            this.gpsInfoList.push(res.data.rebusInfo.gpsInfos[0])
          }
        }
      }
      this.type = type
      for (var i = 0; i < this.gpsInfoList.length; i++) {
        this.gpsInfoList[i].label = new window.BMap.Label('速度: 0 KM/H', {
          offset: new window.BMap.Size(35, -75)
        })
        this.gpsInfoList[i].label.setStyle({
          color: '#fff',
          backgroundColor: '#333333',
          border: '0',
          fontSize: '14px',
          width: '200px',
          height: '100px',
          opacity: '0.8',
          verticalAlign: 'center',
          textAlign: 'left',
          borderRadius: '2px',
          whiteSpace: 'normal',
          wordWrap: 'break-word',
          padding: '5px',
        })
        this.gpsInfoList[i].preMarker = null
        this.gpsInfoList[i].busPlate = ''
        var tempBus = await findBusInfo({ ID: this.gpsInfoList[i].busId })
        // console.log(tempBus)
        if (tempBus.code === 0) {
          if (tempBus.data.rebusInfo && tempBus.data.rebusInfo.busPlate) {
            this.gpsInfoList[i].busPlate = tempBus.data.rebusInfo.busPlate
          }
        }
      }
      // console.log(this.gpsInfoList)
    },
    async getBusInfos() {
      const res = await getBusInfoList()
      if (res.code === 0) {
        this.busInfoList = res.data.list
        // console.log(this.busInfoList.length)
        if (this.busInfoList && this.busInfoList.length > 0) {
          this.currentSelectedPlate = this.busInfoList[0].busPlate
          this.currentSelectedID = this.busInfoList[0].ID
          this.getBusGPSList('single')
        } else {
          this.currentSelectedPlate = 'ALL'
          this.currentSelectedID = -1
        }
        this.busInfoList.push({ busPlate: 'ALL', ID: -1 })
      }
    },
    changeOption(event) {
      this.currentSelectedID = event
      if (this.currentSelectedID !== -1) {
        this.getBusGPSList('single')
      } else if (this.currentSelectedID === -1) {
        this.getBusGPSList('all')
      }
    },
    loadJScript() {
      var script = document.createElement('script')
      script.type = 'text/javascript'
      script.src = '//api.map.baidu.com/getscript?v=2.0&ak=qG1aM2UK80lGHVHKmHczRcmzBBOd8blK'
      script.onload = () => {
        // this.initMap()
        var BMap = window.BMap
        this.carIcon_0 = new BMap.Icon('./images/car_0.png', new BMap.Size(48, 48))
        this.carIcon_1 = new BMap.Icon('./images/car_1.png', new BMap.Size(48, 48))
        this.carIcon_2 = new BMap.Icon('./images/car_2.png', new BMap.Size(48, 48))
        this.carIcon_3 = new BMap.Icon('./images/car_3.png', new BMap.Size(48, 48))
        this.carIcon_4 = new BMap.Icon('./images/car_4.png', new BMap.Size(48, 48))
        this.carIcon_5 = new BMap.Icon('./images/car_5.png', new BMap.Size(48, 48))
        this.carIcon_6 = new BMap.Icon('./images/car_6.png', new BMap.Size(48, 48))
        this.carIcon_7 = new BMap.Icon('./images/car_7.png', new BMap.Size(48, 48))
        this.carIcon_static_0 = new BMap.Icon('./images/car_static_0.png', new BMap.Size(48, 48))
        this.carIcon_static_1 = new BMap.Icon('./images/car_static_1.png', new BMap.Size(48, 48))
        this.carIcon_static_2 = new BMap.Icon('./images/car_static_2.png', new BMap.Size(48, 48))
        this.carIcon_static_3 = new BMap.Icon('./images/car_static_3.png', new BMap.Size(48, 48))
        this.carIcon_static_4 = new BMap.Icon('./images/car_static_4.png', new BMap.Size(48, 48))
        this.carIcon_static_5 = new BMap.Icon('./images/car_static_5.png', new BMap.Size(48, 48))
        this.carIcon_static_6 = new BMap.Icon('./images/car_static_6.png', new BMap.Size(48, 48))
        this.carIcon_static_7 = new BMap.Icon('./images/car_static_7.png', new BMap.Size(48, 48))
        this.currentLabel = new window.BMap.Label('速度: 0 KM/H', {
          offset: new window.BMap.Size(35, -75)
        })
        this.currentLabel.setStyle({
          color: '#fff',
          backgroundColor: '#333333',
          border: '0',
          fontSize: '14px',
          width: '200px',
          height: '180px',
          opacity: '0.8',
          verticalAlign: 'center',
          textAlign: 'left',
          borderRadius: '2px',
          whiteSpace: 'normal',
          wordWrap: 'break-word',
          padding: '5px',
        })
      }
      document.body.appendChild(script)
    },
    setZoom(centerPoint) {
      if (this.currentBound.containsPoint(centerPoint) === false) {
        this.map.centerAndZoom(centerPoint, this.viewSize)
        this.currentBound = this.map.getBounds()
      }
    },
    createMarker(gpsInfo, idx) {
      const BMap = window.BMap
      var centerPoint = new BMap.Point(gpsInfo.lng, gpsInfo.lat)
      console.log('aaaaaaaaaaaaaaaaaaaaaaaaa : ' + gpsInfo.accStatus)
      var marker = null
      var dir = Number(gpsInfo.dir)
      var accStatus = '关'
      if (dir === 0 || dir === 360) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_0 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_0 })
          accStatus = '开'
        }
      } else if (dir > 0 && dir < 90) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_1 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_1 })
          accStatus = '开'
        }
      } else if (dir === 90) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_2 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_2 })
          accStatus = '开'
        }
      } else if (dir > 90 && dir < 180) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_3 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_3 })
          accStatus = '开'
        }
      } else if (dir === 180) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_4 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_4 })
          accStatus = '开'
        }
      } else if (dir > 180 && dir < 270) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_5 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_5 })
          accStatus = '开'
        }
      } else if (dir === 270) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_6 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_6 })
          accStatus = '开'
        }
      } else if (dir > 270 && dir < 360) {
        if (gpsInfo.accStatus === 0) {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_7 })
          accStatus = '关'
        } else {
          marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_7 })
          accStatus = '开'
        }
      }
      var str = '速度: ' + gpsInfo.speed + ' KM/H'
      str += '<br>'
      str += 'IMEI : ' + gpsInfo.gpsSn
      str += '<br>'
      str += '车牌号 : ' + this.gpsInfoList[idx].busPlate
      str += '<br>'
      str += 'ACC : ' + accStatus
      str += '<br>'
      str += '定位时间 : ' + this.getNowDate()
      str += '<br>'

      this.gpsInfoList[idx].label.setContent(str)
      marker.setLabel(this.gpsInfoList[idx].label)
      if ((this.type === 'all' && gpsInfo.speed !== 0) || this.type === 'single') {
        this.setZoom(centerPoint)
      }
      if (this.gpsInfoList[idx].preMarker != null) {
        this.map.removeOverlay(this.gpsInfoList[idx].preMarker)
      }
      this.map.addOverlay(marker)
      this.gpsInfoList[idx].preMarker = marker
    },
    getBusInfoFromList(gspSn) {
      for (var i = 0; i < this.currentGpsList.length; i++) {
        if (this.currentGpsList[i].gpsSn === gspSn) {
          return this.currentGpsList[i]
        }
      }
      return null
    },
    async addMarker() {
      var res
      if (this.type === 'all') {
        res = await locationlist()
        if (res.code === 0) {
          if (res.data.location.length > 0) {
            this.currentGpsList = res.data.location
          }
        }
      }
      for (var i = 0; i < this.gpsInfoList.length; i++) {
        // console.log(this.gpsInfoList[i].gpsSn)
        if (this.type === 'single') {
          const ret = await getLocation({ gpsSn: this.gpsInfoList[i].gpsSn })
          if (ret.code === 0) {
            this.createMarker(ret.data.location, i)
          }
        } else if (this.type === 'all') {
          var temp = this.getBusInfoFromList(this.gpsInfoList[i].gpsSn)
          if (temp !== null) {
            this.createMarker(temp, i)
          }
        }
        /* if (ret.code === 0) {
          this.createMarker(ret.data.location, i)
           const BMap = window.BMap
          var centerPoint = new BMap.Point(ret.data.location.lng, ret.data.location.lat)
          console.log('aaaaaaaaaaaaaaaaaaaaaaaaa : ' + ret.data.location.accStatus)
          var marker = null
          var dir = Number(ret.data.location.dir)
          var accStatus = '关'
          if (dir === 0 || dir === 360) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_0 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_0 })
              accStatus = '开'
            }
          } else if (dir > 0 && dir < 90) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_1 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_1 })
              accStatus = '开'
            }
          } else if (dir === 90) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_2 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_2 })
              accStatus = '开'
            }
          } else if (dir > 90 && dir < 180) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_3 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_3 })
              accStatus = '开'
            }
          } else if (dir === 180) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_4 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_4 })
              accStatus = '开'
            }
          } else if (dir > 180 && dir < 270) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_5 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_5 })
              accStatus = '开'
            }
          } else if (dir === 270) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_6 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_6 })
              accStatus = '开'
            }
          } else if (dir > 270 && dir < 360) {
            if (ret.data.location.accStatus === 0) {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_7 })
              accStatus = '关'
            } else {
              marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_7 })
              accStatus = '开'
            }
          }
          var str = '速度: ' + ret.data.location.speed + ' KM/H'
          str += '<br>'
          str += 'IMEI : ' + this.gpsInfoList[i].gpsSn
          str += '<br>'
          str += 'ACC : ' + accStatus
          str += '<br>'
          str += '定位时间 : ' + this.getNowDate()
          str += '<br>'

          this.gpsInfoList[i].label.setContent(str)
          marker.setLabel(this.gpsInfoList[i].label)
          if ((this.type === 'all' && ret.data.location.speed !== 0) || this.type === 'single') {
            this.setZoom(centerPoint)
          }
          if (this.gpsInfoList[i].preMarker != null) {
            this.map.removeOverlay(this.gpsInfoList[i].preMarker)
          }
          this.map.addOverlay(marker)
          this.gpsInfoList[i].preMarker = marker
        }*/
      }
      if (this.timerId != null) {
        clearTimeout(this.timerId)
      }
      this.timerId = setTimeout(this.addMarker, 5000)
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
    async showRealTimePath() {
      console.log(this.gpsInfoList[0].gpsSn)
      const ret = await getLocation({ gpsSn: this.gpsInfoList[0].gpsSn })
      if (ret.code === 0) {
        if (this.timerId != null) {
          clearTimeout(this.timerId)
        }
        if (this.type === 'all') {
          this.viewSize = 13
        } else {
          this.viewSize = 18
        }
        const BMap = window.BMap
        this.map = new BMap.Map('container')
        this.map.enableScrollWheelZoom(true) // 开启鼠标滚轮缩放
        this.map.addControl(new BMap.NavigationControl()) // 缩放按钮
        var centerPoint = new BMap.Point(ret.data.location.lng, ret.data.location.lat)
        // console.log(this.viewSize)
        this.map.centerAndZoom(centerPoint, this.viewSize)
        this.currentBound = this.map.getBounds()
        this.addMarker()
      }
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.showRealTimePath()
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
