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
  getLocation
} from '@/api/gpsQuery' //  此处请自行替换地址

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
      timerId: null
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
        }
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
          offset: new window.BMap.Size(20, -10)
        })
        this.currentLabel.setStyle({
          color: '#fff',
          backgroundColor: '#333333',
          border: '0',
          fontSize: '14px',
          width: '100px',
          height: '20px',
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
    },
    setZoom(centerPoint) {
      if (this.currentBound.containsPoint(centerPoint) === false) {
        this.map.centerAndZoom(centerPoint, 18)
        this.currentBound = this.map.getBounds()
      }
    },
    async addMarker() {
      const ret = await getLocation({ gpsSn: this.currentSelectGPSSN })
      if (ret.code === 0) {
        const BMap = window.BMap
        var centerPoint = new BMap.Point(ret.data.location.lng, ret.data.location.lat)
        console.log('aaaaaaaaaaaaaaaaaaaaaaaaa : ' + ret.data.location.accStatus)
        // var point = new BMap.Point(ret.data.location.lng, ret.data.location.lat)
        var marker = null
        var dir = Number(ret.data.location.dir)
        if (dir === 0 || dir === 360) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_0 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_0 })
          }
        } else if (dir > 0 && dir < 90) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_1 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_1 })
          }
        } else if (dir === 90) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_2 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_2 })
          }
        } else if (dir > 90 && dir < 180) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_3 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_3 })
          }
        } else if (dir === 180) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_4 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_4 })
          }
        } else if (dir > 180 && dir < 270) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_5 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_5 })
          }
        } else if (dir === 270) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_6 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_6 })
          }
        } else if (dir > 270 && dir < 360) {
          if (ret.data.location.accStatus === 0) {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_static_7 })
          } else {
            marker = new window.BMap.Marker(centerPoint, { icon: this.carIcon_7 })
          }
        }
        this.currentLabel.setContent('速度:' + ret.data.location.speed + ' KM/H')
        marker.setLabel(this.currentLabel)
        this.setZoom(centerPoint)
        if (this.preMarker != null) {
          this.map.removeOverlay(this.preMarker)
        }
        this.map.addOverlay(marker)
        // this.map.panTo(centerPoint)
        this.preMarker = marker
        if (this.timerId != null) {
          clearTimeout(this.timerId)
        }
        this.timerId = setTimeout(this.addMarker, 5000)
      }
    },
    async showRealTimePath() {
      const ret = await getLocation({ gpsSn: this.currentSelectGPSSN })
      if (ret.code === 0) {
        if (this.timerId != null) {
          clearTimeout(this.timerId)
        }
        const BMap = window.BMap
        this.map = new BMap.Map('container')
        this.map.enableScrollWheelZoom(true) // 开启鼠标滚轮缩放
        this.map.addControl(new BMap.NavigationControl()) // 缩放按钮
        var centerPoint = new BMap.Point(ret.data.location.lng, ret.data.location.lat)
        this.map.centerAndZoom(centerPoint, 18)
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
