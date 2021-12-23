<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!-- GPS名字&序列号-->
        <el-form-item label="GPS名字">
          <el-input v-model="searchInfo.gpsName" placeholder="GPS名字" />
        </el-form-item>
        <el-form-item label="GPS序列号">
          <el-input v-model="searchInfo.gpsSn" placeholder="GPS序列号" />
        </el-form-item>
        <el-form-item label="车辆ID">
          <el-input v-model="searchInfo.busId" placeholder="车辆ID" />
        </el-form-item>
        <!-- query&add button -->
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询GPS</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('add')">新增GPS</el-button>
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
    <el-table ref="GPSInfoList" :data="tableData" border stripe @sort-change="sortChange" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="100" />
      <el-table-column label="GPSID" prop="ID" width="200" sortable="custom" align="center" />
      <el-table-column label="GPS序列号" prop="gpsSn" width="300" sortable="custom" align="center" />
      <el-table-column label="GPS名字" prop="gpsName" width="300" sortable="custom" align="center" />
      <el-table-column label="车辆ID" prop="busId" width="300" sortable="custom" align="center">
        <template #default="scope">
          <span>{{ setBusID(scope.row.busId) }}</span>
        </template>
      </el-table-column>
      <!-- edit&del button -->
      <el-table-column fixed="right" label="操作" min-width="200" align="center">
        <template #default="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateGpsInfo(scope.row)">编辑GPS</el-button>
          <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteRow(scope.row)">删除GPS</el-button>
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
      <el-form ref="gpsInfoForm" :model="formData" label-position="right" label-width="100px" :rules="rules" :inline="true">
        <el-form-item label="GPS序列号:" prop="gpsSn">
          <el-input v-model="formData.gpsSn" clearable placeholder="GPS序列号" autocomplete="off" />
        </el-form-item>
        <el-form-item label="GPS名字:" prop="gpsName">
          <el-input v-model="formData.gpsName" clearable placeholder="GPS名字" autocomplete="off" />
        </el-form-item>
        <!--<el-form-item label="车辆ID:" prop="busPlate">
          <el-input v-model.number="formData.busId" clearable placeholder="车辆ID" autocomplete="off" />
        </el-form-item>-->
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
  createGpsInfo,
  deleteGpsInfo,
  deleteGpsInfoByIds,
  updateGpsInfo,
  findGpsInfo,
  getGpsInfoList
} from '@/api/gps_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  name: 'GpsInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getGpsInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        gpsSn: '',
        gpsName: '',
        busId: -1,
      },
      rules: {
        gpsSn: [{ required: true, message: '请输入GPS序列号', trigger: 'blur' }],
        gpsName: [{ required: true, message: '请输入GPS名字', trigger: 'blur' }],
      },
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
    setBusID(busId) {
      if (busId === -1) {
        return '-'
      }
      return busId
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      /* if (this.formData.busId === 0) {
        this.formData.busId = -1
      }*/
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
        this.deleteGpsInfo(row)
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
      const res = await deleteGpsInfoByIds({ ids })
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
      this.$refs.gpsInfoForm.resetFields()
      this.formData = {
        gpsSn: '',
        gpsName: '',
        busId: -1,
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
    async updateGpsInfo(row) {
      const res = await findGpsInfo({ ID: row.ID })
      if (res.code === 0) {
        this.formData = res.data.regpsInfo
        this.openDialog('edit')
      }
    },
    async deleteGpsInfo(row) {
      const res = await deleteGpsInfo({ ID: row.ID })
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
      this.$refs.gpsInfoForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'add':
              res = await createGpsInfo(this.formData)
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
              res = await updateGpsInfo(this.formData)
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
              res = await createGpsInfo(this.formData)
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
