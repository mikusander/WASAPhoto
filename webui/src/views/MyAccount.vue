<template>
    <div>
      <header class="navbar navbar-dark sticky-top bg-dark" style="height: 80px;">
        <div class="container-fluid d-flex justify-content-between align-items-center">
          <div>
            <button @click="newPhoto" class="btn btn-light rounded-pill" style="margin-left: 20px; font-size: 20px; font-weight: bold;">+ Post</button>
          </div>
          <div class="mx-auto text-center">
            <a @click="homePage" style="color: orangered; font-size: 70px; font-weight: bold; cursor: pointer;">WASA Photo</a>
          </div>
          <div class="d-flex align-items-center ml-auto">
            <a @click="doLogout" style="text-decoration: none; color: white; font-size: 30px; font-weight: bold; margin-right: 20px; cursor: pointer;">log out</a>
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
			this.$router.push({path: '/'})
		},
        async homePage() {
			this.$router.push({path: '/users/'+this.user.Username+'/stream'})
		},
        async newPhoto() {
			this.$router.push({path: '/users/'+this.user.Username+'/newPhoto'})
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
  }
  </style>
  