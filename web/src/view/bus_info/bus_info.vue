<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="车主电话">
          <el-input v-model="searchInfo.busUserPhone" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="车主姓名">
          <el-input v-model="searchInfo.busUserName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="车牌">
          <el-input v-model="searchInfo.busPlate" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="公司id">
          <el-input v-model="searchInfo.companyId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
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
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="日期" width="180">
        <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
      </el-table-column>
      <el-table-column label="车主电话" prop="busUserPhone" width="120" />
      <el-table-column label="车主姓名" prop="busUserName" width="120" />
      <el-table-column label="车牌" prop="busPlate" width="120" />
      <el-table-column label="座位数" prop="seatNumber" width="120" />
      <el-table-column label="创建用户id" prop="createUserId" width="120" />
      <el-table-column label="更新用户id" prop="updateUserId" width="120" />
      <el-table-column label="公司id" prop="companyId" width="120" />
      <el-table-column label="按钮组">
        <template #default="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateBusInfo(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="车主电话:">
          <el-input v-model="formData.busUserPhone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="车主姓名:">
          <el-input v-model="formData.busUserName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="车牌:">
          <el-input v-model="formData.busPlate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="座位数:">
          <el-input v-model.number="formData.seatNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建用户id:">
          <el-input v-model.number="formData.createUserId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="更新用户id:">
          <el-input v-model.number="formData.updateUserId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公司id:">
          <el-input v-model.number="formData.companyId" clearable placeholder="请输入" />
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
  createBusInfo,
  deleteBusInfo,
  deleteBusInfoByIds,
  updateBusInfo,
  findBusInfo,
  getBusInfoList
} from '@/api/bus_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'BusInfo',
  mixins: [infoList],
  data() {
    return {
      listApi: getBusInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        busUserPhone: '',
        busUserName: '',
        busPlate: '',
        seatNumber: 0,
        createUserId: 0,
        updateUserId: 0,
        companyId: 0,
      }
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
    async updateBusInfo(row) {
      const res = await findBusInfo({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rebusInfo
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        busUserPhone: '',
        busUserName: '',
        busPlate: '',
        seatNumber: 0,
        createUserId: 0,
        updateUserId: 0,
        companyId: 0,
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
      switch (this.type) {
        case 'create':
          res = await createBusInfo(this.formData)
          break
        case 'update':
          res = await updateBusInfo(this.formData)
          break
        default:
          res = await createBusInfo(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
