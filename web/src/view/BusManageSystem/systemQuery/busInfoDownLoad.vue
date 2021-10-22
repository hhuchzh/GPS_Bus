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
        <el-form-item label="开始时间:" prop="startTime">
          <el-date-picker v-model="currentSelectedStartDate" placeholder="请选择日期" value-format="YYYY-MM-DD" type="date" :disabled-date="disabledStartDate" />
        </el-form-item>
        <el-form-item label="结束时间:" prop="endTime">
          <el-date-picker v-model="currentSelectedEndDate" placeholder="请选择日期" value-format="YYYY-MM-DD" type="date" :disabled-date="disabledEndDate" @change="changeStartFormatDate" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-download" @click="downLoadFile('checkin')">导出打卡数据</el-button>
          <el-button size="mini" type="primary" icon="el-icon-download" @click="downLoadFile('mileage')">导出里程数据</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  findBusInfo,
  getBusInfoList
} from '@/api/bus_info'
// import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  data() {
    return {
      busInfoList: [],
      currentSelectedID: -1,
      currentSelectedPlate: '',
      currentSelectGPSSN: '',
      currentSelectedEndDate: this.getNowFormatDate(),
      currentSelectedStartDate: this.getStartFormatDate(),
    }
  },
  async created() {
    this.getBusInfos()
  },
  methods: {
    /* async getListData(page = this.listdata.page, pageSize = this.listdata.pageSize) {
      var data = await getClassesInfoList({ page, pageSize, busId: parseInt(this.currentSelectedID) })
      if (data.code === 0) {
        for (var idx = 0; idx < data.data.list.length; idx++) {
          data.data.list[idx]['busPlate'] = this.currentSelectedPlate
          var routeId = parseInt(data.data.list[idx].routeId)
          var ret = await findRouteInfo({ ID: routeId })
          if (ret.code === 0) {
            data.data.list[idx]['routeName'] = ret.data.rerouteInfo.routeName
          }
        }
        this.listdata.page = data.data.page
        this.listdata.pageSize = data.data.pageSize
        this.listdata.total = data.data.total
        this.listdata.tableData = data.data.list
      }
    },*/
    disabledStartDate(date) {
      var endDate = new Date(this.currentSelectedEndDate)
      return (date.getTime() < endDate - 24 * 60 * 60 * 1000 * 31) || (date.getTime() > endDate)
    },
    disabledEndDate(date) {
      // alert('date')
      return date.getTime() > Date.now() - 24 * 60 * 60 * 1000
    },
    async getBusInfos() {
      const res = await getBusInfoList()
      if (res.code === 0) {
        this.busInfoList = res.data.list
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
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order === 'descending'
      }
      this.getTableData()
    },
    async getBusInfo() {
      const res = await findBusInfo({ ID: this.currentSelectedID })
      if (res.code === 0) {
        if (res.data.rebusInfo && res.data.rebusInfo.gpsInfos[0]) {
          this.currentSelectGPSSN = res.data.rebusInfo.gpsInfos[0].gpsSn
        }
      }
    },
    getNowFormatDate() {
      var date = new Date()
      date = date - 1000 * 60 * 60 * 24
      date = new Date(date)
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
    getStartFormatDate() {
      var date = new Date()
      date = date - 1000 * 60 * 60 * 24 * 30
      date = new Date(date)
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
    changeStartFormatDate() {
      var date = this.currentSelectedEndDate
      date = new Date(date)
      date = date - 1000 * 60 * 60 * 24 * 30
      date = new Date(date)
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
      this.currentSelectedStartDate = year + seperator1 + month + seperator1 + strDate
      console.log(this.currentSelectedStartDate)
    },
    downLoadFile(type) {
      // TBD
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
