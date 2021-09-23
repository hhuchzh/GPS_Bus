<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('add')">添加班次</el-button>
          <el-button size="mini" type="primary" @click="toBusRouteAdmin">返回</el-button>
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
              <el-button icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <!--result of table -->
    <el-table ref="busClassList" :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="50" />
      <el-table-column label="路线ID" prop="routeId" min-width="100" sortable="custom" align="center" />
      <el-table-column label="班次ID" prop="ID" min-width="100" sortable="custom" align="center" />
      <el-table-column label="路线名称" prop="routeName" min-width="200" sortable="custom" align="center">
        <template #default="scope">
          <div>
            {{ scope.row.routeName = routename }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="发车时间" prop="classesTime" min-width="100" sortable="custom" align="center">
        <template #default="scope">
          <div>
            {{ methodFiletr(scope.row.classesTime) }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="车牌号" prop="busPlate" min-width="150" align="center">
        <template #default="scope">
          <div>
            {{ getBusPlate(scope.row.busInfo) }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="备注" prop="remark" min-width="150" align="center" />
      <!--edit&del button -->
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <template #default="scope">
          <el-button size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateClassesInfo(scope.row)">修改班次</el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" @click="deleteRow(scope.row)">删除班次</el-button>
          <el-button size="mini" type="success" icon="el-icon-edit" @click="toBusLocationArrivalAdmin(scope.row.ID)">站点到达</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form ref="busClassForm" :inline="true" :model="formData" label-position="right" label-width="100px" :rules="rules">
        <el-form-item min-width="150" label="发车时间:" prop="classesTime">
          <el-time-picker v-model="formData.classesTime" placeholder="发车时间" clearable autocomplete="off" value-format="HH:mm:ss" />
        </el-form-item>
        <el-form-item min-width="150" label="备注:" prop="remark">
          <el-input v-model="formData.remark" placeholder="请输入" clearable autocomplete="off" />
        </el-form-item>
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
  createClassesInfo,
  deleteClassesInfo,
  deleteClassesInfoByIds,
  updateClassesInfo,
  findClassesInfo,
  getClassesInfoList
} from '@/api/classes_info' //  此处请自行替换地址
import {
  getBusInfoList
} from '@/api/bus_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  name: 'ClassesInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getClassesInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      dialogTitle: '添加班次',
      multipleSelection: [],
      routename: '',
      busInfoList: [],
      currentSelectedID: -1,
      currentSelectedPlate: '',
      formData: {
        routeId: parseInt(this.$route.query.ID),
        classesTime: '',
        busId: -1,
        remark: ''
      },
      rules: {
        classesTime: [{ required: true, message: '请输入发车时间', trigger: 'blur' }],
        // busId: [{ required: true, message: '请选择车牌号', trigger: 'blur' }],
      },
    }
  },
  async created() {
    this.searchInfo.routeId = parseInt(this.$route.query.ID)
    this.routename = this.$route.query.routeName
    await this.getTableData()
    this.getBusInfos()
    // alert(parseInt(this.$route.query.ID))
  },
  methods: {
    toBusRouteAdmin() {
      this.$router.push({ path: './busRouteAdmin' })
    },
    toBusLocationArrivalAdmin(id) {
      this.$router.push({ path: './busLocationArrivalAdmin', query: { classId: id, routeId: parseInt(this.$route.query.ID) }})
    },
    methodFiletr(value) {
      return value.toString()
    },
    getBusPlate(value) {
      if (value != null && value.busPlate != null) {
        return value.busPlate
      } else {
        return '-'
      }
    },
    async getBusInfos() {
      const res = await getBusInfoList()
      if (res.code === 0) {
        this.busInfoList = res.data.list
        console.log(this.busInfoList.length)
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
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order === 'descending'
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteClassesInfo(row)
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
      const res = await deleteClassesInfoByIds({ ids })
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
      this.$refs.busClassForm.resetFields()
      this.formData = {
        routeId: parseInt(this.$route.query.ID),
        classesTime: '',
        busId: -1,
        remark: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'add':
          this.dialogTitle = '添加班次'
          if (this.busInfoList && this.busInfoList.length > 0) {
            this.currentSelectedPlate = this.busInfoList[0].busPlate
            this.currentSelectedID = this.busInfoList[0].ID
          } else {
            this.currentSelectedPlate = ''
            this.currentSelectedID = -1
          }
          break
        case 'edit':
          this.dialogTitle = '修改班次'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async updateClassesInfo(row) {
      const res = await findClassesInfo({ ID: row.ID })
      this.currentSelectedPlate = row.busInfo.busPlate
      // alert(this.currentSelectedPlate)
      this.type = 'edit'
      if (res.code === 0) {
        // res.data.reclassesInfo.busInfo = null
        this.formData = res.data.reclassesInfo
        this.openDialog('edit')
      }
    },
    async deleteClassesInfo(row) {
      const res = await deleteClassesInfo({ ID: row.ID })
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
      this.$refs.busClassForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              this.formData.busId = this.currentSelectedID
              res = await createClassesInfo(this.formData)
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
              this.formData.busId = this.currentSelectedID
              this.formData.busInfo = null
              res = await updateClassesInfo(this.formData)
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
              res = await createClassesInfo(this.formData)
              break
          }
        }
      })
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
