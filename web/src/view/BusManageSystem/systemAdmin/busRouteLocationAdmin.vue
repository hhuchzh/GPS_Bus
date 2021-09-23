<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('add')">新增站点</el-button>
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
    <el-table ref="busRouteLocationList" :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="50" />
      <el-table-column label="站点ID" prop="ID" min-width="100" sortable="custom" align="center" />
      <el-table-column label="站点名称" prop="locationName" min-width="200" sortable="custom" align="center" />
      <el-table-column label="所属路线" prop="routeId" min-width="100" sortable="custom" align="center" />
      <el-table-column label="经度" prop="longtitude" min-width="200" sortable="custom" align="center" />
      <el-table-column label="纬度" prop="latitude" min-width="200" sortable="custom" align="center" />
      <el-table-column label="位于站数" prop="orderNo" min-width="100" sortable="custom" align="center" />
      <!--edit&del button -->
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <template #default="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateLocationInfo(scope.row)">修改站点</el-button>
          <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteRow(scope.row)">删除站点</el-button>
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
    <!-- add&edit popup-->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form ref="busRouteLocationForm" :inline="true" :model="formData" label-position="right" label-width="100px" :rules="rules">
        <el-form-item min-width="150" label="站点名称:" prop="locationName">
          <el-input v-model="formData.locationName" clearable placeholder="站点名称" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="经度:" prop="longtitude">
          <el-input v-model="formData.longtitude" clearable placeholder="经度" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="纬度:" prop="latitude">
          <el-input v-model="formData.latitude" clearable placeholder="纬度" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="位于站数:" prop="orderNo">
          <el-input v-model.number="formData.orderNo" clearable placeholder="位于站数" autocomplete="off" />
        </el-form-item>
        <el-form-item min-width="150" label="所属路线:" prop="routeId">
          <el-input v-model.number="formData.routeId" readonly />
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
  createLocationInfo,
  deleteLocationInfo,
  deleteLocationInfoByIds,
  updateLocationInfo,
  findLocationInfo,
  getLocationInfoList
} from '@/api/location_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  name: 'LocationInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getLocationInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      dialogTitle: '新增站点',
      multipleSelection: [],
      formData: {
        locationName: '',
        routeId: parseInt(this.$route.query.ID),
        longtitude: '',
        latitude: '',
        orderNo: 1,
      },
      rules: {
        locationName: [{ required: true, message: '请输入站点名称', trigger: 'blur' }],
        longtitude: [{ required: true, message: '请输入经度', trigger: 'blur' }],
        latitude: [{ required: true, message: '请输入纬度', trigger: 'blur' }],
        orderNo: [{ required: true, message: '请输入位于站数', trigger: 'blur' }],
      },
    }
  },
  async created() {
    this.searchInfo.routeId = parseInt(this.$route.query.ID)
    await this.getTableData()
  },
  methods: {
  // 条件搜索前端看此方法
    toBusRouteAdmin() {
      this.$router.push({ path: './busRouteAdmin' })
    },
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
        this.deleteLocationInfo(row)
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
      const res = await deleteLocationInfoByIds({ ids })
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
      this.$refs.busRouteLocationForm.resetFields()
      this.formData = {
        locationName: '',
        routeId: parseInt(this.$route.query.ID),
        longtitude: '',
        latitude: '',
        orderNo: 0,
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'add':
          this.dialogTitle = '新增站点'
          break
        case 'edit':
          this.dialogTitle = '修改站点'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async updateLocationInfo(row) {
      const res = await findLocationInfo({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.relocationInfo
        this.openDialog('edit')
      }
    },
    async deleteLocationInfo(row) {
      const res = await deleteLocationInfo({ ID: row.ID })
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
      this.$refs.busRouteLocationForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              res = await createLocationInfo(this.formData)
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
              res = await updateLocationInfo(this.formData)
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
              res = await createLocationInfo(this.formData)
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
