<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!-- 路线名称 -->
        <el-form-item label="路线名">
          <el-input v-model="searchInfo.routeName" placeholder="路线名" />
        </el-form-item>
        <!-- query&add button -->
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询路线</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('add')">新增路线</el-button>
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <!-- batch del button -->
            <template #reference>
              <el-button icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <!-- details -->
    <el-table ref="busRouteInfoList" :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <!--result of busRouteInfo table -->
      <el-table-column type="selection" width="50" />
      <el-table-column label="路线ID" min-width="50" prop="ID" sortable="custom" align="center" />
      <el-table-column label="路线名" prop="routeName" min-width="150" sortable="custom" align="center" />
      <el-table-column label="起始地点" prop="startAddress" min-width="150" sortable="custom" align="center" />
      <el-table-column label="结束地点" prop="endAddress" min-width="150" sortable="custom" align="center" />
      <el-table-column label="大约时间" prop="aboutTime" min-width="150" sortable="custom" align="center">
        <template #default="scope">
          <div>
            {{ methodFiletr(scope.row.aboutTime) }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="大约距离" prop="aboutDistance" min-width="150" sortable="custom" align="center" />
      <!-- edit&del button -->
      <el-table-column fixed="right" label="操作" min-width="300" align="center">
        <template #default="scope">
          <el-button size="mini" type="primary" icon="el-icon-edit" @click="updateRouteInfo(scope.row)">修改路线</el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" @click="deleteRow(scope.row)">删除路线</el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" @click="toBusRouteLocationAdmin(scope.row.ID)">编辑站点</el-button>
          <el-button size="mini" type="success" icon="el-icon-search" @click="toBusClassAdmin(scope.row.ID, scope.row.routeName)">查看班次</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- pages -->
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
    <!-- add&edit popup-->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form ref="routeForm" :inline="true" :model="formData" label-position="right" label-width="100px" :rules="rules">
        <el-form-item min-width="150" label="路线名:" prop="routeName">
          <el-input v-model="formData.routeName" clearable placeholder="路线名" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="起始地点:" prop="startAddress">
          <el-input v-model="formData.startAddress" clearable placeholder="起始地点" autocomplete="off" />
        </el-form-item>
        <br>
        <el-form-item min-width="150" label="结束地点:" prop="endAddress">
          <el-input v-model="formData.endAddress" clearable placeholder="结束地点" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="大约时间:" prop="aboutTime">
          <el-time-picker v-model="formData.aboutTime" placeholder="大约时间" clearable autocomplete="off" />
        </el-form-item>
        <br>
        <el-form-item min-width="150" label="大约距离:" prop="aboutDistance">
          <el-input v-model.number="formData.aboutDistance" clearable placeholder="大约距离" autocomplete="off" />
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
  createRouteInfo,
  deleteRouteInfo,
  deleteRouteInfoByIds,
  updateRouteInfo,
  findRouteInfo,
  getRouteInfoList
} from '@/api/route_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  name: 'RouteInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getRouteInfoList,
      dialogFormVisible: false,
      dialogTitle: '新增路线',
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        routeName: '',
        startAddress: '',
        endAddress: '',
        aboutTime: '',
        aboutDistance: '',
      },
      rules: {
        routeName: [{ required: true, message: '请输入路线名', trigger: 'blur' }],
        startAddress: [{ required: true, message: '请输入起始地点', trigger: 'blur' }],
        endAddress: [{ required: true, message: '请输入结束地点', trigger: 'blur' }],
        aboutTime: [{ required: true, message: '请输入大约时间', trigger: 'blur' }],
        aboutDistance: [{ required: true, message: '请输入大约距离', trigger: 'blur' }]
      },
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
    toBusRouteLocationAdmin(id) {
      this.$router.push({ path: './busRouteLocationAdmin', query: { ID: id }})
    },
    toBusClassAdmin(id, name) {
      this.$router.push({ path: './busClassAdmin', query: { ID: id, routeName: name }})
    },
    methodFiletr(value) {
      var startIdx = value.indexOf('T')
      var endIdx = value.indexOf('+')
      console.log('startIdx : ' + startIdx + '/ endIdx : ' + endIdx)
      if (startIdx >= 0 && endIdx >= 0) {
        return value.substring(startIdx + 1, endIdx)
      }
      return '-'
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
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
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
        this.deleteRouteInfo(row)
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
      const res = await deleteRouteInfoByIds({ ids })
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
      this.$refs.routeForm.resetFields()
      this.formData = {
        routeName: '',
        startAddress: '',
        endAddress: '',
        aboutTime: '',
        aboutDistance: '',
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'add':
          this.dialogTitle = '新增路线'
          break
        case 'edit':
          this.dialogTitle = '修改路线'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async updateRouteInfo(row) {
      const res = await findRouteInfo({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.rerouteInfo
        this.openDialog('edit')
      }
    },
    async deleteRouteInfo(row) {
      const res = await deleteRouteInfo({ ID: row.ID })
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
      this.$refs.routeForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              res = await createRouteInfo(this.formData)
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
              res = await updateRouteInfo(this.formData)
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
              res = await createRouteInfo(this.formData)
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
