<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!-- GPS名字&序列号-->
        <el-form-item label="站点名">
          <el-input v-model="searchInfo.locationName" placeholder="站点名" />
        </el-form-item>
        <!-- query&add button -->
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询站点</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('add')">新增站点</el-button>
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
    <el-table ref="locationInfoList" :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="100" />
      <el-table-column label="ID" prop="ID" width="100" sortable="custom" align="center" />
      <el-table-column label="站点名" prop="locationName" width="300" sortable="custom" align="center" />
      <el-table-column label="经度" prop="longtitude" width="300" sortable="custom" align="center" />
      <el-table-column label="纬度" prop="latitude" width="300" sortable="custom" align="center" />
      <!-- edit&del button -->
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <template #default="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateLocationCommon(scope.row)">编辑站点</el-button>
          <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteRow(scope.row)">删除站点</el-button>
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
      <el-form ref="locationInfoForm" :model="formData" label-position="right" label-width="100px" :rules="rules" :inline="true">
        <el-form-item label="站点名:" prop="locationName">
          <el-input v-model="formData.locationName" clearable placeholder="请输入站点名" autocomplete="off" />
        </el-form-item>
        <el-form-item label="经度:" prop="longtitude">
          <el-input v-model="formData.longtitude" clearable placeholder="请输入经度" autocomplete="off" />
        </el-form-item>
        <el-form-item label="纬度:" prop="latitude">
          <el-input v-model="formData.latitude" clearable placeholder="请输入纬度" autocomplete="off" />
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
  createLocationCommon,
  deleteLocationCommon,
  deleteLocationCommonByIds,
  updateLocationCommon,
  findLocationCommon,
  getLocationCommonList
} from '@/api/location_common'
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  name: 'LocationCommon',
  mixins: [infoList],
  data() {
    return {
      listApi: getLocationCommonList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        locationName: '',
        longtitude: '',
        latitude: ''
      },
      rules: {
        locationName: [{ required: true, message: '请输入站点名', trigger: 'blur' }],
        longtitude: [{ required: true, message: '请输入经度', trigger: 'blur' }],
        latitude: [{ required: true, message: '请输入纬度', trigger: 'blur' }]
      },
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
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
        this.deleteLocationCommon(row)
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
      const res = await deleteLocationCommonByIds({ ids })
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
      this.$refs.locationInfoForm.resetFields()
      this.formData = {
        locationName: '',
        longtitude: '',
        latitude: '',
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
    async updateLocationCommon(row) {
      const res = await findLocationCommon({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.relocationCommon
        this.openDialog('edit')
      }
    },
    async deleteLocationCommon(row) {
      const res = await deleteLocationCommon({ ID: row.ID })
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
      this.$refs.locationInfoForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              console.log(this.formData)
              res = await createLocationCommon(this.formData)
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
              res = await updateLocationCommon(this.formData)
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
              res = await createLocationCommon(this.formData)
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
