<script>
export default {
  data() {
    return {
      errormsg: null,
      loading: false,
      some_data: null,
      username: '',
      user: {
        ID: 0,
        Username: "",
      },
    };
  },
  methods: {
    async doLogin(){
      if(this.username == ""){
        this.errormsg = "il campo username è vuoto"
      }
      else{
        try{
          let response = await this.$axios.post("/session", { username: this.username })
          this.user = response.data
          localStorage.setItem("token",this.user.ID)
          localStorage.setItem("username",this.user.Username)
          this.$router.push({path: '/users/'+this.user.Username+'/stream'})
		
        }
        catch(e){
			if (e.response && e.response.status === 400) {
				this.errormsg = "Form error, please check all fields and try again";
			}else if(e.response && e.response.status === 500){
				this.errormsg = "Server error, please try again later";
			}else{
				this.errormsg = e.toString();
			}
        }		
      }
    },
  },
};
</script>

<template>
  <header class="navbar navbar-dark sticky-top bg-dark" style="height: 80px;">
		<div class="d-flex justify-content-center w-100">
			<a href="#/" style="text-decoration: none; color: orangered; font-size: 70px; font-weight: bold;">WASA Photo</a>
		</div>
	</header>
	<div class="container">
	  <img src="../images/utente.png" alt="Mia Immagine">
	  <!-- Form per inserire testo e inviarlo -->
	  <form @submit.prevent="handleSubmit">
		<input v-model="username" type="text" placeholder="Inserisci username" />
		<button type="submit" @click="doLogin">Invia</button>
	  </form>
  
	  <!-- Messaggio di errore -->
	  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
  </template>

<style scoped>
.container {
  color: white; /* Imposta il colore del testo a bianco per la leggibilità */
  min-height: calc(100vh - 80px);
  padding: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center; /* Centra verticalmente */
  align-items: center; /* Centra orizzontalmente */
}

form {
  margin-bottom: 20px;
  display: flex;
  flex-direction: column; /* Dispone gli elementi del form in colonna */
  align-items: center; /* Centra gli elementi del form */
}

input[type="text"] {
  padding: 10px;
  font-size: 16px;
  margin-bottom: 10px; /* Aggiungi margine inferiore per spaziatura */
  width: 100%; /* Imposta la larghezza dell'input */
  max-width: 300px; /* Larghezza massima dell'input */
  border-radius: 10px; /* Arrotonda i bordi dell'input */
}

button {
  padding: 10px 20px;
  font-size: 16px;
  background-color: white;
  color: black;
  border: none;
  cursor: pointer;
  width: 100%; /* Imposta la larghezza del bottone */
  max-width: 300px; /* Larghezza massima del bottone */
  border-radius: 10px; /* Arrotonda i bordi del bottone */
}

button:hover {
  background-color: #ddd;
}

img {
  margin-bottom: 20px; /* Aggiungi margine inferiore per spaziatura */
}
</style>

  