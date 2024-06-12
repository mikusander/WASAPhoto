<template>
  <div>
    <header class="navbar navbar-dark sticky-top bg-dark" style="height: 80px;">
      <div class="container-fluid d-flex justify-content-between align-items-center">
        <div class="align-items-center text-center flex-grow-1"> <!-- Modifica qui -->
          <span style="color: orangered; font-size: 70px; font-weight: bold; cursor: default;">WASA Photo</span>
        </div>
        <div class="d-flex align-items-center">
          <a @click="doLogout"
            style="text-decoration: none; color: white; font-size: 30px; font-weight: bold; margin-left: auto; cursor: pointer;">log
            out</a>
          <a @click="myAccount" class="text-white" style="margin-left: 20px;">
            <img src="../images/utente.png" alt="User Image"
              style="width: 70px; height: 70px; border-radius: 50%; cursor: pointer;">
          </a>
        </div>
      </div>
    </header>
    <div class="container">
      <!-- Messaggio di errore -->
      <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      errormsg: null,
      loading: false,
      some_data: null,
      user: {
        ID: localStorage.getItem("token"),
        Username: localStorage.getItem("username"),
      },
    };
  },
  methods: {
    async doLogout() {
      localStorage.removeItem("token")
      localStorage.removeItem("username")
      this.$router.push({ path: '/' })
    },
    async myAccount() {
      this.$router.push({ path: '/users/' + this.user.Username + '/MyAccount' })
    },
  },
};
</script>

<style scoped>
.container {
  color: white;
  min-height: calc(100vh - 80px);
  padding: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}</style>
