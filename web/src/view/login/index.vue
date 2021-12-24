<template>
  <div class="loginback">
    <div class="login-bg">
      <div class="login">
        <h1 class="logo-login">
          <img src="~@/assets/title.png" width="318" height="81" alt="Tracker">
        </h1>
        <el-form
          ref="loginForm"
          :model="loginForm"
          :rules="rules"
          @keyup.enter="submitForm"
        >
          <div>
            <div class="form-group form-username">
              <el-form-item prop="username">
                <el-input v-model="loginForm.username" placeholder="请输入用户名">
                  <template #suffix>
                    <i class="el-input__icon el-icon-user" />
                  </template>
                </el-input>
              </el-form-item>
            </div>
            <div class="form-group form-password">
              <el-form-item prop="password">
                <el-input
                  v-model="loginForm.password"
                  :type="lock === 'lock' ? 'password' : 'text'"
                  placeholder="请输入密码"
                >
                  <template #suffix>
                    <i
                      :class="'el-input__icon el-icon-' + lock"
                      @click="changeLock"
                    />
                  </template>
                </el-input>
              </el-form-item>
            </div>
            <div class="form-group form-username">
              <el-form-item style="position: relative">
                <el-input
                  v-model="loginForm.captcha"
                  name="logVerify"
                  placeholder="请输入验证码"
                  style="width: 60%"
                />
                <div class="vPic">
                  <img
                    v-if="picPath"
                    :src="picPath"
                    alt="请输入验证码"
                    @click="loginVerify()"
                  >
                </div>
              </el-form-item>
            </div>
            <div class="form-group">
              <el-form-item>
                <el-button
                  type="primary"
                  style="width: 100%;"
                  @click="submitForm"
                >登 录</el-button>
              </el-form-item>
            </div>
          </div>
        </el-form>
        <span class="mappin">
          <i class="shenzhen" />
          <i class="beijing" />
          <i class="dubai" />
          <i class="canberra" />
          <i class="washington" />
          <i class="losangeles" />
          <i class="santiago" />
        </span>
      </div>
      <footer class="footer" data-automation="footer">
        <div class="copyright ta-c">
          © 2021 研创园车辆管理系统
        </div>
      </footer>
    </div>
  </div>
</template>
<script>
import { mapActions } from 'vuex'
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
export default {
  name: 'Login',
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5) {
        return callback(new Error('请输入正确的用户名'))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6) {
        return callback(new Error('请输入正确的密码'))
      } else {
        callback()
      }
    }
    return {
      curYear: 0,
      lock: 'lock',
      loginForm: {
        username: '',
        password: '',
        captcha: '',
        captchaId: ''
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }]
      },
      logVerify: '',
      picPath: '',
      ad_1: document.getElementById('son_1'),
      ad_2: document.getElementById('son_2'),
      ad_3: document.getElementById('son_3'),
      ad_4: document.getElementById('son_4'),
      adList: [],
      idx: 0
    }
  },
  created() {
    this.loginVerify()
    this.curYear = new Date().getFullYear()
  },
  mounted() {
    /* this.ad_1 = document.getElementById('son_1')
    this.ad_2 = document.getElementById('son_2')
    this.ad_3 = document.getElementById('son_3')
    this.ad_4 = document.getElementById('son_4')
    this.adList.push(this.ad_1)
    this.adList.push(this.ad_2)
    this.adList.push(this.ad_3)
    this.adList.push(this.ad_4)
    setInterval(this.changeAD, 2000)*/
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    async checkInit() {
      const res = await checkDB()
      if (res.code === 0) {
        if (res.data?.needInit) {
          this.$store.commit('user/NeedInit')
          this.$router.push({ name: 'Init' })
        } else {
          this.$message({
            type: 'info',
            message: '已配置数据库信息，无法初始化'
          })
        }
      }
    },
    async login() {
      return await this.LoginIn(this.loginForm)
    },
    async submitForm() {
      this.$refs.loginForm.validate(async(v) => {
        if (v) {
          const flag = await this.login()
          if (!flag) {
            this.loginVerify()
          }
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写登录信息',
            showClose: true
          })
          this.loginVerify()
          return false
        }
      })
    },
    changeLock() {
      this.lock = this.lock === 'lock' ? 'unlock' : 'lock'
    },
    loginVerify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath
        this.loginForm.captchaId = ele.data.captchaId
      })
    },
    changeAD() {
      for (var i = 0; i < 4; i++) {
        if (i === this.idx) {
          this.adList[i].style.display = 'block'
        } else {
          this.adList[i].style.display = 'none'
        }
      }
      this.idx++
      if (this.idx === 4) {
        this.idx = 0
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/style/logincss.scss";
</style>
