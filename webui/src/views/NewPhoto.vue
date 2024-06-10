<template>
    <div>
      <header class="navbar navbar-dark sticky-top bg-dark" style="height: 80px;">
        <div class="container-fluid d-flex justify-content-between align-items-center">
          <div class="align-items-center text-center flex-grow-1">
            <a @click="homePage" style="color: orangered; font-size: 70px; font-weight: bold; cursor: pointer;">WASA Photo</a>
          </div>
          <div class="d-flex align-items-center">
            <a @click="myAccount" class="text-white" style="margin-left: 20px;">
              <img src="../images/utente.png" alt="Immagine Utente" style="width: 70px; height: 70px; border-radius: 50%; cursor: pointer;">
            </a>
          </div>
        </div>
      </header>
      <div class="container">
        <div class="d-flex align-items-center" style="height: 80px;">
          <form @submit.prevent="handleSubmit" class="d-flex flex-column align-items-start ml-3" enctype="multipart/form-data">
            <input type="file" ref="imageInput" class="form-control-file mb-3" accept="image/*" required>
            <textarea v-model="description" class="form-control form-control-lg form-control-auto-resize mb-3" placeholder="Descrizione" required></textarea>
            <button type="submit" class="btn btn-primary btn-lg">Carica</button>
          </form>
        </div>
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
        description: "",
        user: {
          ID: localStorage.getItem("token"),
          Username: localStorage.getItem("username"),
        },
        photo: {
          ID: 0,
          Date: "",
          Text: "",
          URL: "",
          LikeCounter: 0,
          CommentCounter: 0,
          UserID: 0,
        }
      };
    },
    methods: {
      async myAccount() {
        this.$router.push({ path: '/users/' + this.user.Username + '/MyAccount' });
      },
  
      async homePage() {
        this.$router.push({ path: '/users/' + this.user.Username + '/stream' });
      },
  
      async handleSubmit() {
        const fileInput = this.$refs.imageInput;
        const description = this.description;
        const file = fileInput.files[0];
  
        if (file) {
          const reader = new FileReader();
          reader.onloadend = async () => {
            const base64String = reader.result.replace("data:", "").replace(/^.+,/, "");
  
            try {
                const response = await this.$axios.post("/users/" + this.user.Username + "/photo", 
                    { URL: base64String, Text: description }, 
                    { headers: { Authorization: `Bearer ${this.user.ID}` } }
                );
              this.photo = response.data;
              console.log("URL: ", base64String);
              console.log("Text: ", description);
              this.errormsg = null; // Resetta il messaggio di errore
              this.$router.push({ path: '/users/' + this.user.Username + '/MyAccount' });
            } catch (e) {
              if (e.response && e.response.status === 400) {
                this.errormsg = "Errore nel modulo, controlla tutti i campi e riprova";
              } else if (e.response && e.response.status === 500) {
                this.errormsg = "Errore del server, riprova più tardi";
              } else {
                this.errormsg = e.toString();
              }
            }
          };
          reader.readAsDataURL(file); // Avvia la lettura del file
        }
      }
    }
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
  .form-control-auto-resize {
    resize: none;
    overflow: hidden;
    min-height: 100px;
    width: 1000px; /* Modifica secondo necessità */
  }
  </style>
  