<template>
  <div>
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
      <el-form-item>
        <el-button size="mini" type="primary" @click="save">保存</el-button>
        <el-button size="mini" type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {
  createBusInfo,
  updateBusInfo,
  findBusInfo
} from '@/api/bus_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'BusInfo',
  mixins: [infoList],
  data() {
    return {
      type: '',
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
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findBusInfo({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rebusInfo
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
  },
  methods: {
    async save() {
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
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
