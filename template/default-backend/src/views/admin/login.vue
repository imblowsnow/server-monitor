<script>
import NavHeader from "@/components/nav-header.vue";
import {login} from "@/api/backend/user.js";
import {useUserStore} from "@/stores/user.js";
export default {
  name: "login",
  components: {NavHeader},
  data() {
    return {
      userStore: useUserStore(),
      loginForm: {
        username: "",
        password: "",
      },
    };
  },
  methods: {
    login() {
      login(this.loginForm).then((token) => {
        this.userStore.setToken(token);
        this.$router.push("/");
      }).catch((err) => {
        this.$message.error(err);
      });
    },
  },
}
</script>

<template>
  <div>
    <nav-header :login="true"></nav-header>
    <main>
      <div class="form-container">
        <div class="form">
          <el-form ref="loginForm" :model="loginForm">
            <h1 class="h3 mb-3 fw-normal"></h1>
            <div class="form-floating">
              <input
                  id="floatingInput"
                  type="text"
                  class="form-control"
                  placeholder="Username"
                  autocomplete="username"
                  required=""
                  v-model="loginForm.username"
              /><label for="floatingInput"
            >用户名</label
            >
            </div>
            <div class="form-floating mt-3">
              <input

                  id="floatingPassword"
                  type="password"
                  class="form-control"
                  placeholder="Password"
                  autocomplete="current-password"
                  required=""
                  v-model="loginForm.password"
              /><label for="floatingPassword"
            >密码</label
            >
            </div>
            <div class="form-check mb-3 mt-3 d-flex justify-content-center pe-4">
              <div class="form-check">
                <input
                    id="remember"
                    type="checkbox"
                    class="form-check-input"
                    value="remember-me"
                /><label

                  class="form-check-label"
                  for="remember"
              >记住我</label
              >
              </div>
            </div>
            <button class="w-100 btn btn-primary"  type="button" @click="login">
              登录
            </button
            ><!---->
          </el-form>
        </div>
      </div>
    </main>
  </div>
</template>

<style scoped>
.form-container {
  display: flex;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
}

.form {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
  text-align: center;
}
</style>
