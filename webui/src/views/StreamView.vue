<template>
  <div>
    <header class="navbar navbar-dark sticky-top bg-dark" style="height: 100px;">
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
      <div class="photo-gallery" v-if="(myStream.ListPhoto.length > 0)">
        <div class="gallery">
          <div v-for="photo in myStream.ListPhoto" :key="photo.ID" class="photo">
            <div class="photo-wrapper">
              <img :src="'data:image/jpeg;base64,' + photo.URL" :alt="photo.Text">
              <p style="color: black; font-weight: bold;">Date: {{ photo.Date }}</p>
              <p style="color: black; font-weight: bold;">caricato da: {{ photo.UserUsername }}</p>
              <p style="color: black; font-weight: bold;">Descrizione: {{ photo.Text }}</p>
              <p style="color: black; font-weight: bold;">Like: {{ photo.likeCounter }}</p>
              <p style="color: black; font-weight: bold;">Comment: {{ photo.commentCounter }}</p>
              <div class="button-group">
                <button @click="toggleLike(photo)" class="btn like-button" :class="{ 'liked': photo.liked }">
                  <img v-if="photo.liked" src="../images/heartBlack.png" alt="Liked" class="button-image">
                  <img v-else src="../images/heartWhite.png" alt="Not Liked" class="button-image">
                </button>
                <button @click="photo.showCommentArea = !photo.showCommentArea" class="btn comment-button">
                  <img src="../images/comment.png" alt="Comment" class="button-image">
                </button>
              </div>

              <div>
                <button @click="toggleComments(photo)" class="btn mt-2"
                  style="background: none; border: none; padding: 0;">
                  <p class="clickable-text">{{ photo.showComments ? 'Nascondi Commenti' : 'Mostra Commenti' }}</p>
                </button>
                <div v-if="photo.showComments" class="comments">
                  <div v-for="comment in photo.listComment" :key="comment.ID" style="display: flex; align-items: center;">
                    <div class="comment-content">
                      <p><strong>{{ comment.UserUsername }}:</strong> {{ comment.Text }}</p>
                    </div>
                    <button v-if="comment.User_id == this.user.ID" class="btn delete-button"
                      @click="deleteComment(photo.ID, comment.ID)">
                      <img src="../images/bucket.png" alt="Delete" class="delete-icon">
                    </button>
                  </div>

                </div>
              </div>

              <div v-if="photo.showCommentArea" class="comment-area">
                <textarea v-model="photo.newComment" rows="3" class="form-control mt-2"
                  placeholder="Scrivi un commento..."></textarea>
                <button @click="addComment(photo)" class="btn mt-2"
                  style="background-color: green; color: white;">Invia</button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-else style="text-align: center; color: white;">
        <p>Nessuna foto disponibile</p>
      </div>
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
      myStream: {
        ListPhoto: [],
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
    async toggleComments(photo) {
      photo.showComments = !photo.showComments;
    },
    async deleteComment(photoID, commentID) {
      try {
        await this.$axios.delete(`/users/${this.user.Username}/photo/${photoID}/comment/${commentID}`, {
          headers: { Authorization: `Bearer ${this.user.ID}` }
        });
        this.errormsg = null; // Resetta il messaggio di errore
        this.stream();
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
    async toggleLike(photo) {
      photo.liked = !photo.liked; // Cambia lo stato del like
      photo.likeCounter += photo.liked ? 1 : -1; // Aggiorna il contatore dei like
      if (photo.liked) {
        try {
          let response = await this.$axios.put('/users/' + this.user.Username + '/photo/' + photo.ID + '/like/' + this.user.Username, {}, {
            headers: {
              Authorization: "Bearer " + this.user.ID
            }
          });
        } catch (e) {
          photo.liked = false; // Ripristina lo stato precedente in caso di errore
          photo.likeCounter -= 1; // Ripristina il contatore dei like in caso di errore
          if (e.response && e.response.status === 400) {
            this.errormsg = "Errore nel modulo, controlla tutti i campi e riprova";
          } else if (e.response && e.response.status === 500) {
            this.errormsg = "Errore del server, riprova più tardi";
          } else {
            this.errormsg = e.toString();
          }
        }
      } else {
        try {
          await this.$axios.delete(`/users/${this.user.Username}/photo/${photo.ID}/like/${this.user.Username}`, {
            headers: { Authorization: `Bearer ${this.user.ID}` }
          });
        } catch (e) {
          photo.liked = true; // Ripristina lo stato precedente in caso di errore
          photo.likeCounter += 1; // Ripristina il contatore dei like in caso di errore
          if (e.response && e.response.status === 400) {
            this.errormsg = "Errore nel modulo, controlla tutti i campi e riprova";
          } else if (e.response && e.response.status === 500) {
            this.errormsg = "Errore del server, riprova più tardi";
          } else {
            this.errormsg = e.toString();
          }
        }
      }
    },
    async addComment(photo) {
      if (!photo.newComment) {
        this.errormsg = "Il commento non può essere vuoto";
        return;
      }
      try {
        photo.commentCounter += 1;
        await this.$axios.post(`/users/${this.user.Username}/photo/${photo.ID}/comment`, {
          Text: photo.newComment
        }, {
          headers: { Authorization: `Bearer ${this.user.ID}` }
        });
        photo.newComment = ""; // Pulisci l'area di testo del commento
        photo.showCommentArea = false; // Nascondi l'area di testo dopo aver inviato il commento
        this.errormsg = null; // Resetta il messaggio di errore
        this.stream();
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
    async stream() {
      try {
        let response = await this.$axios.get("/users/" + this.user.Username + "/stream", {
          headers: { Authorization: `Bearer ${this.user.ID}` }
        });
        this.myStream = response.data;
        if (this.myStream.ListPhoto != null) {
          for (let i = 0; i < this.myStream.ListPhoto.length; i++) {
            try {
              const isLike = await this.$axios.get(`/users/${this.user.Username}/photo/${this.myStream.ListPhoto[i].ID}/like/${this.user.Username}`, {
                headers: { Authorization: `Bearer ${this.user.ID}` }
              });
              let like = isLike.data;
              if (like.User_id == 0) {
                this.myStream.ListPhoto[i].liked = false;
              }
              else {
                this.myStream.ListPhoto[i].liked = true;
              }
              this.errormsg = null;
              // Aggiungi una proprietà per gestire la visualizzazione dei commenti
              this.myStream.ListPhoto[i].showComments = false;
              // Aggiungi una proprietà per memorizzare i commenti relativi ad ogni foto
            } catch (e) {
              if (e.response && e.response.status === 400) {
                this.errormsg = "Errore nel modulo, controlla tutti i campi e riprova";
              } else if (e.response && e.response.status === 500) {
                this.errormsg = "Errore del server, riprova più tardi";
              } else {
                this.errormsg = e.toString();
              }
            }
          }
        }
        else {
          this.myStream.ListPhoto = [];
        }
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
    async searchUser() {
      if (this.searchUsername == "") {
        this.errormsg = "il campo username è vuoto"
      } else if (this.searchUsername == this.user.Username) {
        this.$router.push({ path: '/users/' + this.user.Username + '/MyAccount' })
      } else {
        try {
          let risposta = await this.$axios.get("/users/" + this.searchUsername + "/id");
          let id = risposta.data;
          try {
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
  mounted() {
    this.stream();
  }
};
</script>
<style scoped>
.d-flex.align-items-center {
  margin-right: 20px;
  /* Aumenta il margine destro per più spazio */
}

.button-group {
  display: flex;
  gap: 20px; /* Spazio ridotto tra i bottoni */
  justify-content: center; /* Centra i bottoni orizzontalmente */
  align-items: center; /* Allinea i bottoni verticalmente al centro */
}


.like-button, .comment-button {
  background-color: transparent;
  border: none;
  padding: 0; /* Rimuove il padding */
  width: 30px; /* Imposta la larghezza del bottone */
  height: 30px; /* Imposta l'altezza del bottone */
  display: flex;
  justify-content: center; /* Centra l'immagine orizzontalmente */
  align-items: center; /* Centra l'immagine verticalmente */
}

.button-image {
  width: 30px; /* Imposta la larghezza dell'immagine */
  height: 30px; /* Imposta l'altezza dell'immagine */
}

.buttons-wrapper {
  display: flex;
  align-items: center;
}

.btn {
    margin-right: 10px;
    /* Aumenta il margine destro per più spazio */
    border-radius: 50px;
    /* Rendi i bordi completamente tondeggianti */
}

.comment-wrapper {
  display: flex;
  /* Utilizza Flexbox */
  align-items: center;
  /* Allinea verticalmente al centro */
}

.comment-content {
  flex: 1;
  /* Fai espandere il contenuto del commento per riempire lo spazio disponibile */
  margin-right: 10px;
  /* Aggiungi uno spazio a destra tra il contenuto del commento e il pulsante */
}

.delete-button {
  background-color: transparent;
  /* Imposta lo sfondo trasparente */
  border: none;
  /* Rimuovi il bordo */
  padding: 0;
  /* Rimuovi il padding */
  cursor: pointer;
  /* Mostra il cursore come puntatore */
  height: 20px;
  width: 20px;
}

.delete-icon {
  width: 20px;
  height: 20px;
}


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
  width: 200px;
  /* Modifica la larghezza in base alle tue esigenze */
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
  padding: 30px;
  /* Aumenta il padding per rendere lo sfondo più grosso */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  display: inline-block;
  border: 2px solid #ccc;
  /* Aggiungi un bordo per un effetto più definito */
  width: 500px;
  /* Imposta una larghezza fissa */
  height: auto;
  /* Imposta un'altezza fissa */
  overflow: hidden;
  /* Nasconde l'eccesso dell'immagine se esce dal riquadro */
}

.photo img {
  width: 100%;
  /* Rende l'immagine larga quanto il contenitore */
  height: 100%;
  /* Rende l'immagine alta quanto il contenitore */
  object-fit: cover;
  /* Ritaglia l'immagine per riempire il contenitore mantenendo il centro */
  border-radius: 10px;
}

.photo p {
  margin-top: 10px;
  font-size: 14px;
  color: black;
}

.comment-area {
  margin-top: 10px;
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
  width: 1500px;
  font-size: 16px;
  margin-bottom: 10px;
  border-radius: 10px;
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

.clickable-text {
  font-weight: bold;
  cursor: pointer;
  transition: color 0.3s;
}

.clickable-text:hover {
  color: rgba(0, 0, 0, 0.5);
  /* Schiarisce il testo al passaggio del mouse */
}
</style>
