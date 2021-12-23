<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!--query fileds-->
        <el-form-item label="车主姓名">
          <el-input v-model="searchInfo.busUserName" placeholder="车主姓名" />
        </el-form-item>
        <el-form-item label="车主电话">
          <el-input v-model="searchInfo.busUserPhone" placeholder="车主电话" />
        </el-form-item>
        <el-form-item label="车牌号">
          <el-input v-model="searchInfo.busPlate" placeholder="车牌号" />
        </el-form-item>
        <!--query/add button-->
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询车辆</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('add')">新增车辆</el-button>
          <!--batch del popup-->
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <!--batch del button-->
            <template #reference>
              <el-button icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table ref="busInfoList" :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <!--result of busInfo table -->
      <el-table-column type="selection" width="50" />
      <el-table-column label="车辆ID" min-width="100" prop="ID" sortable="custom" align="center" />
      <el-table-column label="车主姓名" prop="busUserName" min-width="150" sortable="custom" align="center" />
      <el-table-column label="车主电话" prop="busUserPhone" min-width="150" sortable="custom" align="center" />
      <el-table-column label="车牌号" prop="busPlate" min-width="150" sortable="custom" align="center" />
      <el-table-column label="座位数" prop="seatNumber" min-width="100" sortable="custom" align="center" />
      <el-table-column label="GPS序列号" prop="gpsSn" min-width="100" sortable="custom" align="center">
        <template #default="scope">
          <span>{{ getGpsSn(scope.row.gpsInfos) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="GPS名字" prop="gpsName" min-width="100" sortable="custom" align="center">
        <template #default="scope">
          <span>{{ getGpsName(scope.row.gpsInfos) }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <!--edit&del button -->
        <template #default="scope">
          <el-button size="small" type="success" @click="openDialogGPS(scope.row.gpsInfos, scope.row.ID)">
            {{ setGPSBtnTitle(scope.row.gpsInfos) }}
          </el-button>
          <el-button size="small" type="primary" icon="el-icon-edit" @click="updateBusInfo(scope.row)">编辑车辆</el-button>
          <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteRow(scope.row)">删除车辆</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--pages -->
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!--add&edit popup-->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form ref="busInfoForm" :model="formData" label-position="right" label-width="100px" :rules="rules" :inline="true">
        <el-form-item min-width="150" label="车主姓名:" prop="busUserName">
          <el-input v-model="formData.busUserName" clearable placeholder="车主姓名" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="车牌号:" prop="busPlate">
          <el-input v-model="formData.busPlate" clearable placeholder="车牌号" autocomplete="off" />
        </el-form-item>
        <br>
        <el-form-item min-width="150" label="车主电话:" prop="busUserPhone">
          <el-input v-model="formData.busUserPhone" clearable placeholder="车主电话" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="座位数:" prop="seatNumber">
          <el-input v-model.number="formData.seatNumber" clearable placeholder="座位数" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <!-- add&edit gps popup-->
    <el-dialog v-model="dialogFormVisibleGPS" :before-close="closeDialogGPS" :title="dialogTitleGPS">
      <el-form ref="gpsInfoForm" :model="gpsFormData" label-position="right" label-width="200px" :inline="true">
        <template v-if="isBinding">
          <el-form-item label="GPS序列号(GPS名字)" prop="gpsSn">
            <el-select v-model="defaultSelect" placeholder="请选择" @change="changeOption($event)">
              <el-option
                v-for="item in gpsOptions"
                :key="item.gpsSn"
                :label="`${item.gpsSn}(${item.gpsName})`"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
        </template>
        <template v-else>
          <el-form-item label="GPS序列号:" prop="gpsSn">
            <el-input v-model="gpsFormData.gpsSn" :readonly="isReadOnly" />
          </el-form-item>
          <el-form-item label="GPS名字:" prop="gpsName">
            <el-input v-model="gpsFormData.gpsName" :readonly="isReadOnly" />
          </el-form-item>
          <el-form-item label="GPSID:" prop="ID">
            <el-input v-model.number="gpsFormData.ID" :readonly="isReadOnly" />
          </el-form-item>
        </template>
        <el-form-item label="车辆ID:" prop="busId">
          <el-input v-model.number="gpsFormData.busId" readonly />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialogGPS">取 消</el-button>
          <el-button type="primary" @click="enterDialogGPS">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createBusInfo,
  deleteBusInfo,
  deleteBusInfoByIds,
  updateBusInfo,
  findBusInfo,
  getBusInfoList
} from '@/api/bus_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
import {
  updateGpsInfo,
  findGpsInfo,
  getAvailableGpsInfoList
} from '@/api/gps_info' //  此处请自行替换地址
export default {
  name: 'BusInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getBusInfoList,
      dialogFormVisible: false,
      dialogTitle: '新增车辆',
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        busUserName: '',
        busUserPhone: '',
        busPlate: '',
        seatNumber: ''
      },
      rules: {
        busUserName: [{ required: true, message: '请输入车主姓名', trigger: 'blur' }],
        busUserPhone: [{ required: true, message: '请输入车主电话', trigger: 'blur' }],
        busPlate: [{ required: true, message: '请选择车牌', trigger: 'blur' }],
        seatNumber: [{ required: true, message: '请输入座位数', trigger: 'blur' }]
      },
      // gps logic
      gpsFormData: {
        ID: 0,
        gpsSn: '',
        gpsName: '',
        busId: 0,
      },
      dialogFormVisibleGPS: false,
      dialogTitleGPS: '绑定GPS',
      isReadOnly: true,
      gpsOptions: [],
      isBinding: false,
      currentSelected: -1,
      defaultSelect: '',
    }
  },
  async created() {
    await this.getTableData()
    // await this.getGPSOptions()
  },
  methods: {
    // GPS method
    async getGPSOptions() {
      const res = await getAvailableGpsInfoList()
      if (res.code === 0) {
        this.gpsOptions = res.data.list
        console.log(this.gpsOptions.length)
        if (this.gpsOptions && this.gpsOptions.length > 0) {
          this.defaultSelect = this.gpsOptions[0].gpsSn + '(' + this.gpsOptions[0].gpsName + ')'
          this.currentSelected = this.gpsOptions[0].ID
        } else {
          this.defaultSelect = ''
          this.currentSelected = -1
        }
      }
    },
    getGpsName(str) {
      if (str && str[0]) {
        return str[0].gpsName
      }
      return '-'
    },
    getGpsSn(str) {
      if (str && str[0]) {
        return str[0].gpsSn
      }
      return '-'
    },
    getGpsID(str) {
      if (str && str[0]) {
        return str[0].ID
      }
      return 0
    },
    setGPSBtnTitle(str) {
      var textTitle = '绑定GPS'
      if (str && str[0] && (str[0].gpsSn.length > 0)) {
        textTitle = '解除GPS'
      }
      return textTitle
    },
    async openDialogGPS(str, busid) {
      this.dialogTitleGPS = this.setGPSBtnTitle(str)
      if (this.dialogTitleGPS === '解除GPS') {
        this.isReadOnly = true
        const res = await findGpsInfo({ ID: this.getGpsID(str) })
        if (res.code === 0) {
          this.gpsFormData = res.data.regpsInfo
          this.isBinding = false
        }
      } else {
        this.getGPSOptions()
        this.isReadOnly = false
        this.gpsFormData.busId = busid
        this.isBinding = true
      }

      this.dialogFormVisibleGPS = true
    },
    closeDialogGPS() {
      this.dialogFormVisibleGPS = false
    },
    async enterDialogGPS() {
      let gpsres
      switch (this.dialogTitleGPS) {
        case '解除GPS':
          this.gpsFormData.busId = -1
          gpsres = await updateGpsInfo(this.gpsFormData)
          if (gpsres.code === 0) {
            this.$message({
              type: 'success',
              message: '解除GPS成功'
            })
            this.closeDialogGPS()
            this.getTableData()
          }
          break
        case '绑定GPS':
          var res = await findGpsInfo({ ID: this.currentSelected })
          if (res.code === 0) {
            var temp_busid = this.gpsFormData.busId
            this.gpsFormData = res.data.regpsInfo
            this.gpsFormData.busId = temp_busid
            gpsres = await updateGpsInfo(this.gpsFormData)
            if (gpsres.code === 0) {
              this.$message({
                type: 'success',
                message: '绑定GPS成功'
              })
              this.closeDialogGPS()
              this.getTableData()
            }
          }
          break
        default:
          gpsres = await createBusInfo(this.formData)
          break
      }
    },
    changeOption(event) {
      this.currentSelected = event
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    sortChange({ prop, order }) {
      console.log(prop, order)
      if (prop) {
        console.log('1')
        this.searchInfo.orderKey = toSQLLine(prop)
        console.log(this.searchInfo.orderKey)
        this.searchInfo.desc = order === 'descending'
      }
      this.getTableData()
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteBusInfo(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteBusInfoByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    initForm() {
      this.$refs.busInfoForm.resetFields()
      this.formData = {
        busUserPhone: '',
        busUserName: '',
        busPlate: '',
        seatNumber: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'add':
          this.dialogTitle = '新增车辆'
          break
        case 'edit':
          this.dialogTitle = '编辑车辆'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async updateBusInfo(row) {
      const res = await findBusInfo({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.rebusInfo
        this.openDialog('edit')
      }
    },
    async deleteBusInfo(row) {
      const res = await deleteBusInfo({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      this.$refs.busInfoForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              res = await createBusInfo(this.formData)
              if (res.code === 0) {
                this.$message({
                  type: 'success',
                  message: '添加成功'
                })
                this.closeDialog()
                this.getTableData()
              }
              break
            case 'edit':
              res = await updateBusInfo(this.formData)
              if (res.code === 0) {
                this.$message({
                  type: 'success',
                  message: '修改成功'
                })
                this.closeDialog()
                this.getTableData()
              }
              break
            default:
              res = await createBusInfo(this.formData)
              break
          }
        }
      })
    },
  }
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
