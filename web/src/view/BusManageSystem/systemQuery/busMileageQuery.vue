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
          <el-date-picker v-model="currentSelectedDate" placeholder="请选择日期" type="date" :disabled-date="disabledDate" value-format="YYYY-MM-DD" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <!--<el-button size="mini" type="primary" icon="el-icon-download" @click="downLoadFile">导出上月里程</el-button>-->
        </el-form-item>
      </el-form>
    </div>
    <!--result of table -->
    <el-table ref="busInfoList" :data="listdata.tableData" border stripe @sort-change="sortChange">
      <el-table-column label="车牌号" prop="busPlate" min-width="100" sortable="custom" align="center" />
      <el-table-column label="路线名称" prop="routeName" min-width="200" sortable="custom" align="center" />
      <el-table-column label="班次ID" prop="ID" min-width="100" sortable="custom" align="center" />
      <el-table-column label="发车时间" prop="classesTime" min-width="100" sortable="custom" align="center" />
      <el-table-column label="里程数(KM)" prop="distance" min-width="100" sortable="custom" align="center" />
      <!--edit&del button
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <template #default="scope">
          <el-button size="mini" type="primary" icon="el-icon-search" class="table-button" @click="searchCheckIn(scope.row.ID)">打卡查询</el-button>
        </template>
      </el-table-column>-->
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
  getMilesInfoList
} from '@/api/miles_info'
import {
  findRouteInfo
} from '@/api/route_info'
// import infoList from '@/mixins/infoList'
// import { toSQLLine } from '@/utils/stringFun'
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
          var res = await getMilesInfoList({ classesId: parseInt(data.data.list[idx].ID), calcDate: this.currentSelectedDate })
          if (res.code === 0) {
            data.data.list[idx]['distance'] = this.getDistance(res.data.list[0].distance)
          }
        }
        this.listdata.page = data.data.page
        this.listdata.pageSize = data.data.pageSize
        this.listdata.total = data.data.total
        this.listdata.tableData = data.data.list
      }
    },
    handleSizeChangeList(val) {
      this.listdata.pageSize = val
      this.getListData()
    },
    handleCurrentChangeList(val) {
      this.listdata.page = val
      this.getListData()
    },
    getDistance(data) {
      var tmp = Math.floor(data * 100) / 100000
      if (tmp > 0) {
        return tmp.toPrecision(4)
      } else {
        return tmp
      }
    },
    disabledDate(date) {
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
    initParameter() {
      this.listdata.page = 1
      this.listdata.pageSize = 10
      this.listdata.total = 10
      this.listdata.tableData = []
    },
    // 条件搜索前端看此方法
    onSubmit() {
      // TBD
      // this.listdata.page = 1
      this.initParameter()
      this.getListData()
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
