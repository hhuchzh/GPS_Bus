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
          <!--<el-date-picker v-model="currentSelectedDate" placeholder="请选择日期" clearable autocomplete="off" value-format="YYYY-MM-DD" />-->
          <el-date-picker v-model="currentSelectedDate" placeholder="请选择日期" value-format="YYYY-MM-DD" type="date" :disabled-date="disabledDate" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <!--<el-button size="mini" type="primary" icon="el-icon-download" @click="downLoadFile">导出上月打卡</el-button>-->
        </el-form-item>
      </el-form>
    </div>
    <!--result of table -->
    <el-table ref="busInfoList" :data="listdata.tableData" border stripe @sort-change="sortChange">
      <el-table-column label="车牌号" prop="busPlate" min-width="100" sortable="custom" align="center" />
      <el-table-column label="路线名称" prop="routeName" min-width="200" sortable="custom" align="center" />
      <el-table-column label="班次ID" prop="ID" min-width="100" sortable="custom" align="center" />
      <el-table-column label="发车时间" prop="classesTime" min-width="100" sortable="custom" align="center" />
      <!--edit&del button -->
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <template #default="scope">
          <el-button size="mini" type="primary" icon="el-icon-search" class="table-button" @click="searchCheckIn(scope.row.ID)">打卡查询</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--pages -->
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="listdata.page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="listdata.total"
      @current-change="handleCurrentChangeList"
      @size-change="handleSizeChangeList"
    />
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="打卡详情">
      <el-table ref="locationList" :data="locationListdata.tableData" border stripe>
        <el-table-column label="站点名称" prop="arrivalInfo.Location.locationName" min-width="100" align="center" />
        <el-table-column label="站点时间" prop="arrivalInfo.arrivalTime" min-width="200" align="center" />
        <el-table-column label="打卡状态" prop="checkinTime" min-width="100" align="center">
          <template #default="scope">
            <span v-if="scope.row.checkinTime==='00:00:00'" style="color: red">异常打卡</span>
            <span v-else style="color: black">正常打卡</span>
          </template>
        </el-table-column>
      </el-table>
      <!--pages -->
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="locationListdata.page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="locationListdata.total"
        @current-change="handleCurrentChangeLocation"
        @size-change="handleSizeChangeLocation"
      />
      <template #footer>
        <div class="dialog-footer" />
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  findBusInfo,
  getBusInfoList
} from '@/api/bus_info'
import {
  getClassesInfoList
} from '@/api/classes_info'
import {
  findRouteInfo
} from '@/api/route_info'
import {
  getCheckinInfoList,
  createCheckinInfo
} from '@/api/checkin_info'
// import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  data() {
    return {
      busInfoList: [],
      currentSelectedID: -1,
      currentSelectedPlate: '',
      currentSelectGPSSN: '',
      currentSelectedDate: this.getNowFormatDate(),
      listdata: {
        page: 1,
        total: 10,
        pageSize: 10,
        tableData: [],
        queryInfo: {
          busId: parseInt(this.currentSelectedID)
        }
      },
      dialogFormVisible: false,
      locationListdata: {
        page: 1,
        total: 10,
        pageSize: 10,
        tableData: [],
        queryInfo: {
          classesId: parseInt(this.currentClassId),
          checkinDate: this.currentSelectedDate
        }
      },
      currentClassId: -1,
      textColor: 'black'
    }
  },
  async created() {
    this.getBusInfos()
  },
  methods: {
    async getListData(page = this.listdata.page, pageSize = this.listdata.pageSize) {
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
    },
    disabledDate(date) {
      // alert('date')
      return date.getTime() > Date.now() - 24 * 60 * 60 * 1000
    },
    handleSizeChangeList(val) {
      this.listdata.pageSize = val
      this.getListData()
    },
    handleCurrentChangeList(val) {
      this.listdata.page = val
      this.getListData()
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
    // 条件搜索前端看此方法
    onSubmit() {
      // TBD
      this.listdata.page = 1
      this.getListData()
    },
    searchCheckIn(classID) {
      this.dialogFormVisible = true
      this.currentClassId = classID
      this.getCheckInList()
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.locationListdata.page = 1
    },
    handleSizeChangeLocation(val) {
      this.locationListdata.pageSize = val
      this.getCheckInList()
    },
    handleCurrentChangeLocation(val) {
      this.locationListdata.page = val
      this.getCheckInList()
    },
    async getCheckInList(page = this.locationListdata.page, pageSize = this.locationListdata.pageSize) {
      const res = await getCheckinInfoList({ page, pageSize, classesId: parseInt(this.currentClassId), checkinDate: this.currentSelectedDate })
      if (res.code === 0) {
        console.log(res)
        this.locationListdata.page = res.data.page
        this.locationListdata.pageSize = res.data.pageSize
        this.locationListdata.total = res.data.total
        this.locationListdata.tableData = res.data.list
      }
    },
    createCheckinInfoTest() {
      createCheckinInfo({ classesId: 41, locationId: 63, checkinTime: '2021-10-15 07:00:00', checkinDate: '2021-10-15', arrivalId: 38 })
      createCheckinInfo({ classesId: 41, locationId: 64, checkinTime: '2021-10-15 07:10:00', checkinDate: '2021-10-15', arrivalId: 39 })
      createCheckinInfo({ classesId: 41, locationId: 65, checkinTime: '2021-10-15 07:20:00', checkinDate: '2021-10-15', arrivalId: 40 })
      createCheckinInfo({ classesId: 41, locationId: 66, checkinTime: '2021-10-15 07:30:00', checkinDate: '2021-10-15', arrivalId: 41 })
      createCheckinInfo({ classesId: 41, locationId: 67, checkinTime: '2021-10-15 07:40:00', checkinDate: '2021-10-15', arrivalId: 42 })
      createCheckinInfo({ classesId: 41, locationId: 68, checkinTime: '', checkinDate: '2021-10-15', arrivalId: 43 })
      createCheckinInfo({ classesId: 41, locationId: 69, checkinTime: '2021-10-15 07:00:00', checkinDate: '2021-10-15', arrivalId: 44 })
      createCheckinInfo({ classesId: 41, locationId: 70, checkinTime: '2021-10-15 07:00:00', checkinDate: '2021-10-15', arrivalId: 45 })
      createCheckinInfo({ classesId: 41, locationId: 71, checkinTime: '', checkinDate: '2021-10-15', arrivalId: 46 })
      createCheckinInfo({ classesId: 41, locationId: 72, checkinTime: '2021-10-15 07:00:00', checkinDate: '2021-10-15', arrivalId: 47 })
      createCheckinInfo({ classesId: 41, locationId: 73, checkinTime: '', checkinDate: '2021-10-15', arrivalId: 48 })
    }
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
