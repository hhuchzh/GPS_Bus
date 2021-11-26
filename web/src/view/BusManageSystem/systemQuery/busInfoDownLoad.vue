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
          <el-date-picker v-model="currentSelectedStartDate" placeholder="请选择日期" value-format="YYYY-MM-DD" type="date" :disabled-date="disabledStartDate" @change="changeEndFormatDate" />
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
  exportExcelMileage,
  exportExcelCheckIn
} from '@/api/excel'
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
      currentSelectedEndDate: this.getNowFormatDate(),
      currentSelectedStartDate: this.getStartFormatDate(),
    }
  },
  async created() {
    this.getBusInfos()
  },
  methods: {
    disabledStartDate(date) {
      // var endDate = new Date(this.currentSelectedEndDate)
      var endDate = Date.now() - 24 * 60 * 60 * 1000
      return (date.getTime() < endDate - 24 * 60 * 60 * 1000 * 30) || (date.getTime() > endDate)
    },
    disabledEndDate(date) {
      // alert('date')
      var endDate = Date.now() - 24 * 60 * 60 * 1000
      return (date.getTime() > Date.now() - 24 * 60 * 60 * 1000) || (date.getTime() < endDate - 24 * 60 * 60 * 1000 * 30)
    },
    async getBusInfos() {
      const res = await getBusInfoList()
      if (res.code === 0) {
        this.busInfoList = res.data.list
        console.log(this.busInfoList)
        if (this.busInfoList && this.busInfoList.length > 0) {
          this.currentSelectedPlate = this.busInfoList[0].busPlate
          this.currentSelectedID = this.busInfoList[0].ID
        } else {
          this.currentSelectedPlate = ''
          this.currentSelectedID = -1
        }
      }
    },
    changeOption(event) {
      this.currentSelectedID = event
      this.getBusPlate()
    },
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order === 'descending'
      }
      this.getTableData()
    },
    async getBusPlate() {
      const res = await findBusInfo({ ID: this.currentSelectedID })
      if (res.code === 0) {
        if (res.data.rebusInfo) {
          this.currentSelectedPlate = res.data.rebusInfo.busPlate
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
    changeEndFormatDate() {
      if (new Date(this.currentSelectedStartDate) > new Date(this.currentSelectedEndDate)) {
        this.currentSelectedEndDate = this.getNowFormatDate()
      }
    },
    changeStartFormatDate() {
      /* var date = this.currentSelectedEndDate
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
      this.currentSelectedStartDate = year + seperator1 + month + seperator1 + strDate*/
      if (new Date(this.currentSelectedStartDate) > new Date(this.currentSelectedEndDate)) {
        this.currentSelectedStartDate = this.getStartFormatDate()
      }
      console.log(this.currentSelectedStartDate)
    },
    downLoadFile(type) {
      var fileName
      if (type === 'checkin') {
        fileName = this.currentSelectedPlate + '_checkin_from_' + this.currentSelectedStartDate + '_to_' + this.currentSelectedEndDate + '.xlsx'
        exportExcelCheckIn({ plate: this.currentSelectedPlate, beginTime: this.currentSelectedStartDate, endTime: this.currentSelectedEndDate }, fileName)
      } else if (type === 'mileage') {
        fileName = this.currentSelectedPlate + '_miles_from_' + this.currentSelectedStartDate + '_to_' + this.currentSelectedEndDate + '.xlsx'
        exportExcelMileage({ plate: this.currentSelectedPlate, beginTime: this.currentSelectedStartDate, endTime: this.currentSelectedEndDate }, fileName)
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
