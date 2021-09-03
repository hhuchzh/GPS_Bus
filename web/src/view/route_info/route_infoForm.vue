<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="创建人ID:">
        <el-input v-model.number="formData.createUserId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="修改人ID:">
        <el-input v-model.number="formData.updateUserId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="路线名:">
        <el-input v-model="formData.routeName" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="起始地点:">
        <el-input v-model="formData.startAddress" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="结束地点:">
        <el-input v-model="formData.endAddress" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="大约时间:">
        <el-date-picker v-model="formData.aboutTime" type="date" placeholder="选择日期" clearable></el-date-picker>
      </el-form-item>
      <el-form-item label="大约距离:">
        <el-input v-model.number="formData.aboutDistance" clearable placeholder="请输入" />
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
  createRouteInfo,
  updateRouteInfo,
  findRouteInfo
} from '@/api/route_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'RouteInfo',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        createUserId: 0,
        updateUserId: 0,
        routeName: '',
        startAddress: '',
        endAddress: '',
        aboutTime: new Date(),
        aboutDistance: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findRouteInfo({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rerouteInfo
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
          res = await createRouteInfo(this.formData)
          break
        case 'update':
          res = await updateRouteInfo(this.formData)
          break
        default:
          res = await createRouteInfo(this.formData)
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
