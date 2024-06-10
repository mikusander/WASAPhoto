<template>
    <div>
      <header class="navbar navbar-dark sticky-top bg-dark" style="height: 100px;">
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
      <div class="user-info">
        <div class="num-info">
            <span style="cursor: default;">{{ user.Username }}</span>
        </div>
        <div class="num-info">
            <span style="cursor: default;">{{ profile.NumPhoto }}</span>
            <span style="cursor: default;">Numero di foto</span>
        </div>
        <div class="num-info">
            <span style="cursor: default;">{{ profile.NumFollow }}</span>
            <span style="cursor: default;">Seguiti</span>
        </div>
        <div class="num-info">
            <span style="cursor: default;">{{ profile.NumFollowing }}</span>
            <span style="cursor: default;">Seguaci</span>
        </div>
        </div>
      <div class="container">
        <div class="photo-gallery">
            <div class="gallery">
                <div v-for="photo in profile.ListPhoto" :key="photo.ID" class="photo">
                    <div class="photo-wrapper">
                        <img :src="'data:image/jpeg;base64,' + photo.URL" :alt="photo.Text">
                        <p style="color: black; font-weight: bold;">Descrizione: {{ photo.Text }}</p>
                        <button @click="deletePhoto(photo.ID)" class="btn btn-danger mt-2">Elimina</button>
                    </div>
                </div>
            </div>
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
        user: {
            ID: localStorage.getItem("token"),
            Username: localStorage.getItem("username"),
        },
        profile: {
            NumFollow: 0,
            NumFollowing: 0,
            NumPhoto: 0,
            ListPhoto: [
            {
                ID: 0,
                Date: "",
                Text: "",
                URL: "",
                LikeCounter: 0,
                CommentCounter: 0,
                UserID: 0,
            }
            ],
            UserOwner: {
            ID: 0,
            Username: "",
            },
        }
        };
    },
    methods: {
        async doLogout() {
        localStorage.removeItem("token");
        localStorage.removeItem("username");
        this.$router.push({ path: '/' });
        },
        async homePage() {
        this.$router.push({ path: '/users/' + this.user.Username + '/stream' });
        },
        async newPhoto() {
        this.$router.push({ path: '/users/' + this.user.Username + '/newPhoto' });
        },
        async getProfile() {
        try {
            const response = await this.$axios.get("/users/" + this.user.Username + "/profile",
            { headers: { Authorization: `Bearer ${this.user.ID}` } }
            );
            this.profile = response.data;
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
        },
        async deletePhoto(photoID) {
        try {
            await this.$axios.delete(`/users/${this.user.Username}/photo/${photoID}`, {
            headers: { Authorization: `Bearer ${this.user.ID}` }
            });
            this.profile.ListPhoto = this.profile.ListPhoto.filter(photo => photo.ID !== photoID);
            this.errormsg = null; // Resetta il messaggio di errore
            this.profile.NumPhoto--;
        } catch (e) {
            if (e.response && e.response.status === 400) {
            this.errormsg = "Errore nel modulo, controlla tutti i campi e riprova";
            } else if (e.response && e.response.status === 500) {
            this.errormsg = "Errore del server, riprova più tardi";
            } else {
            this.errormsg = e.toString();
            }
        }
        },
    },
    mounted() {
        this.getProfile();
    }
    };
</script>

  
 <style scoped>
.user-info {
  display: flex;
  align-items: center;
  justify-content: space-around;
  text-align: center;
  width: 100%;
  height: 100px;
  margin-top: 60px;
}

.num-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 200px; /* Modifica la larghezza in base alle tue esigenze */
}

.num-info span {
  color: white;
  font-size: 24px;
  font-weight: bold;
}
.container {
  color: white;
  min-height: calc(100vh - 220px);
  padding: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.photo-gallery {
  text-align: center;
  color: white;
}

.gallery {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.photo {
  margin: 20px;
  text-align: center;
}

.photo-wrapper {
  background-color: white;
  border-radius: 15px;
  padding: 30px; /* Aumenta il padding per rendere lo sfondo più grosso */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  display: inline-block;
  border: 2px solid #ccc; /* Aggiungi un bordo per un effetto più definito */
  width: 500px; /* Imposta una larghezza fissa */
  height: 500px; /* Imposta un'altezza fissa */
  overflow: hidden; /* Nasconde l'eccesso dell'immagine se esce dal riquadro */
}

.photo img {
  width: 100%; /* Rende l'immagine larga quanto il contenitore */
  height: 100%; /* Rende l'immagine alta quanto il contenitore */
  object-fit: cover; /* Ritaglia l'immagine per riempire il contenitore mantenendo il centro */
  border-radius: 10px;
}

.photo p {
  margin-top: 10px;
  font-size: 14px;
  color: black;
}

button.btn-danger {
  font-size: 16px;
  font-weight: bold;
  color: white;
  border: none;
  border-radius: 5px;
  padding: 10px 20px;
  background-color: red;
  cursor: pointer;
}
</style>

  