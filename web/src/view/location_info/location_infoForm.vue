<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="创建人ID:">
        <el-input v-model.number="formData.createUserId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="修改人ID:">
        <el-input v-model.number="formData.updateUserId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="站点名:">
        <el-input v-model="formData.locationName" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="routeId字段:">
        <el-input v-model.number="formData.routeId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="经度:">
        <el-input v-model="formData.longtitude" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="纬度:">
        <el-input v-model="formData.latitude" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="站点在路线中的位置，第几站:">
        <el-input v-model.number="formData.orderNo" clearable placeholder="请输入" />
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
  createLocationInfo,
  updateLocationInfo,
  findLocationInfo
} from '@/api/location_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'LocationInfo',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        createUserId: 0,
        updateUserId: 0,
        locationName: '',
        routeId: 0,
        longtitude: '',
        latitude: '',
        orderNo: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findLocationInfo({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.relocationInfo
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
          res = await createLocationInfo(this.formData)
          break
        case 'update':
          res = await updateLocationInfo(this.formData)
          break
        default:
          res = await createLocationInfo(this.formData)
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
