<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="创建人ID:">
        <el-input v-model.number="formData.createUserId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="修改人ID:">
        <el-input v-model.number="formData.updateUserId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="routeId字段:">
        <el-input v-model.number="formData.routeId" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="备注:">
        <el-input v-model="formData.remark" clearable placeholder="请输入" />
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
  createClassesInfo,
  updateClassesInfo,
  findClassesInfo
} from '@/api/classes_info' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'ClassesInfo',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        createUserId: 0,
        updateUserId: 0,
        routeId: 0,
        remark: '',
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findClassesInfo({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.reclassesInfo
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
          res = await createClassesInfo(this.formData)
          break
        case 'update':
          res = await updateClassesInfo(this.formData)
          break
        default:
          res = await createClassesInfo(this.formData)
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
