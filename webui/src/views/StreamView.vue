<template>
  <div>
    <header class="navbar navbar-dark sticky-top bg-dark" style="height: 80px;">
      <div class="container-fluid d-flex justify-content-between align-items-center">
        <div class="align-items-center text-center flex-grow-1">
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
      <form @submit.prevent="searchUser">
        <input v-model="searchUsername" type="text" placeholder="Inserisci un utente" />
        <button type="submit">Cerca</button>
      </form>
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
      profile: {
        NumFollow: 0,
        NumFollowing: 0,
        NumPhoto: 0,
        ListPhoto: []
      },
      searchUsername: "",
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
    async searchUser() {
      if (this.searchUsername == "") {
        this.errormsg = "il campo username è vuoto"
      } else if (this.searchUsername == this.user.Username) {
        this.$router.push({ path: '/users/' + this.user.Username + '/MyAccount' })
      } else {
        try {
          let risposta = await this.$axios.get("/users/" + this.searchUsername + "/id");
          let id = risposta.data;
          let response = await this.$axios.get("/users/" + id.Username + "/profile",
            { headers: { Authorization: `Bearer ${id.ID}` } }
          );
          let utente = response.data;
          this.$router.push({ path: '/users/' + utente.UserOwner.Username + '/profile' });
        }
        catch (e) {
          if (e.response && e.response.status === 400) {
            this.errormsg = "Form error, please check all fields and try again";
          } else if (e.response && e.response.status === 404) {
            this.errormsg = "Utente non trovato";
          } else if (e.response && e.response.status === 500) {
            this.errormsg = "Server error, please try again later";
          } else {
            this.errormsg = e.toString();
          }
        }
      }
    },
  },
};
</script>
<style scoped>
.container {
  color: white;
  min-height: calc(100vh - 80px);
  padding: 50px;
}

form {
  margin-bottom: 20px;
  display: flex;
  flex-direction: column;
  /* Dispone gli elementi del form in colonna */
  align-items: center;
  /* Centra gli elementi del form */
}

input[type="text"] {
  padding: 10px;
  font-size: 16px;
  margin-bottom: 10px;
  /* Aggiungi margine inferiore per spaziatura */
  width: 100%;
  /* Imposta la larghezza dell'input */
  /* Larghezza massima dell'input */
  border-radius: 10px;
  /* Arrotonda i bordi dell'input */
}

button {
  padding: 10px 20px;
  font-size: 16px;
  background-color: orangered;
  color: black;
  font-weight: bold;
  border: none;
  cursor: pointer;
  width: 100%;
  /* Imposta la larghezza del bottone */
  max-width: 300px;
  /* Larghezza massima del bottone */
  border-radius: 10px;
  /* Arrotonda i bordi del bottone */
}

button:hover {
  background-color: #ddd;
}
</style>
