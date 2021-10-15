<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="mini" type="primary" @click="toBusRouteAdmin">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table ref="busClassList" :data="listdata.tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="50" />
      <el-table-column label="站点ID" prop="ID" min-width="100" sortable="custom" align="center" />
      <el-table-column label="站点名称" prop="locationName" min-width="150" sortable="custom" align="center" />
      <el-table-column label="站点时间" prop="arrivalTime" min-width="150" sortable="custom" align="center" />
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <template #default="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateArrivalInfo(scope.row)">修改时间</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="listdata.page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="listdata.total"
      @current-change="handleCurrentChange1"
      @size-change="handleSizeChange1"
    />
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="修改时间">
      <el-form ref="busArrivalForm" :inline="true" :model="formData" label-position="right" label-width="100px">
        <el-form-item min-width="150" label="站点时间:" prop="arrivalTime">
          <el-time-picker v-model="formData.arrivalTime" placeholder="站点时间" clearable autocomplete="off" value-format="HH:mm:ss" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createArrivalInfo,
  updateArrivalInfo,
  findArrivalInfo,
  getArrivalInfoList
} from '@/api/arrival_info' //  此处请自行替换地址
import {
  getLocationInfoList
} from '@/api/location_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'ArrivalInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getArrivalInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      locationTable: [],
      formData: {
        locationId: -1,
        classesId: parseInt(this.$route.query.classId),
        arrivalTime: ''
      },
      arrivalFormData: {
        locationId: -1,
        classesId: parseInt(this.$route.query.classId),
        arrivalTime: ''
      },
      listdata: {
        page: 1,
        total: 10,
        pageSize: 10,
        tableData: [],
        queryInfo: {
          routeId: parseInt(this.$route.query.routeId),
          arrivalTime: ''
        },
      }
    }
  },
  async created() {
    // await this.getTableData()
    this.getArrivalInfo()
  },
  methods: {
    async getArrivalInfo(page = this.listdata.page, pageSize = this.listdata.pageSize) {
      var data = await getLocationInfoList({ page, pageSize, ...this.listdata.queryInfo })
      if (data.code === 0) {
        // this.listdata.tableData = data.data.list
        for (var i = 0; i < data.data.list.length; i++) {
          this.arrivalFormData.locationId = data.data.list[i].ID
          const res = await findArrivalInfo(this.arrivalFormData)
          if (res.code === 0 && res.data.rearrivalInfo !== null) {
            data.data.list[i]['arrivalTime'] = res.data.rearrivalInfo.arrivalTime
          } else {
            data.data.list[i]['arrivalTime'] = '-'
          }
        }
        this.listdata.total = data.data.total
        this.listdata.page = data.data.page
        this.listdata.pageSize = data.data.pageSize
        this.listdata.tableData = data.data.list
      }
    },
    toBusRouteAdmin() {
      this.$router.push({ path: './busRouteAdmin' })
    },
    // 条件搜索前端看此方法
    handleSizeChange1(val) {
      this.listdata.pageSize = val
      this.getArrivalInfo()
    },
    handleCurrentChange1(val) {
      this.listdata.page = val
      this.getArrivalInfo()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async updateArrivalInfo(row) {
      this.formData.locationId = row.ID
      const res = await findArrivalInfo(this.formData)
      if (res.code === 0 && res.data.rearrivalInfo !== null) {
        this.formData = res.data.rearrivalInfo
        this.type = 'edit'
      } else {
        this.formData.arrivalTime = ''
        this.type = 'create'
      }
      this.dialogFormVisible = true
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        locationId: -1,
        classesId: parseInt(this.$route.query.classId),
        arrivalTime: ''
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createArrivalInfo(this.formData)
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '创建成功'
            })
            this.closeDialog()
            this.getArrivalInfo()
          }
          break
        case 'edit':
          res = await updateArrivalInfo(this.formData)
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '更改成功'
            })
            this.closeDialog()
            this.getArrivalInfo()
          }
          break
        default:
          res = await createArrivalInfo(this.formData)
          break
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
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
