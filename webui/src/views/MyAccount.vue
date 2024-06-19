<template>
    <div>
        <header class="navbar navbar-dark sticky-top bg-dark" style="height: 100px;">
            <div class="container-fluid d-flex justify-content-between align-items-center">
                <div>
                    <button @click="newPhoto" class="btn btn-light rounded-pill"
                        style="margin-left: 20px; font-size: 20px; font-weight: bold;">+ Post</button>
                </div>
                <div class="mx-auto text-center">
                    <a @click="homePage" style="color: orangered; font-size: 70px; font-weight: bold; cursor: pointer;">WASA
                        Photo</a>
                </div>
                <div class="d-flex align-items-center ml-auto">
                    <a @click="doLogout"
                        style="text-decoration: none; color: white; font-size: 30px; font-weight: bold; margin-right: 20px; cursor: pointer;">log
                        out</a>
                </div>
            </div>
        </header>
        <div class="user-info">
            <div class="num-info">
                <span style="cursor: default;">{{ user.Username }}</span>
                <button @click="toggleChangeUsername()" class="btn btn-light rounded-pill" style="margin-left: 20px; font-size: 15px; font-weight: bold;">
                    {{ this.changeUsername ? 'Nascondi' : 'Cambia username' }}
                </button>
                <form v-if="this.changeUsername" @submit.prevent="usernameChange">
                    <input v-model="newUsername" type="text" placeholder="Inserisci username" />
                    <button class="buttone" type="submit">Cambia</button>
                </form>
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
            <div class="photo-gallery" v-if="profile.ListPhoto.length > 0">
                <div class="gallery">
                    <div v-for="photo in profile.ListPhoto" :key="photo.ID" class="photo">
                        <div class="photo-wrapper">
                            <img :src="'data:image/jpeg;base64,' + photo.URL" :alt="photo.Text">
                            <p style="color: black; font-weight: bold;">Date: {{ photo.Date }}</p>
                            <p style="color: black; font-weight: bold;">Descrizione: {{ photo.Text }}</p>
                            <p style="color: black; font-weight: bold;">Like: {{ photo.likeCounter }}</p>
                            <p style="color: black; font-weight: bold;">Comment: {{ photo.commentCounter }}</p>
                            <button @click="toggleLike(photo)" class="btn mt-2" :class="{ 'liked': photo.liked }"
                                style="background-color: transparent; border: none;">
                                <img v-if="photo.liked" src="../images/heartBlack.png" alt="Liked"
                                    style="height: 30px; width: 30px;">
                                <img v-else src="../images/heartWhite.png" alt="Not Liked"
                                    style="height: 30px; width: 30px;">
                            </button>
                            <button @click="photo.showCommentArea = !photo.showCommentArea" class="btn mt-2"
                                style="background-color: transparent; border: none;">
                                <img src="../images/comment.png" style="width: 30px; height: 30px;">
                            </button>
                            <button @click="deletePhoto(photo.ID)" class="btn mt-2" style="background-color: orangered;">
                                <img src="../images/bucket.png" style="width: 20px; height: 20px;">
                            </button>
                            <div>
                                <button @click="toggleComments(photo)" class="btn mt-2" style="background: none; border: none; padding: 0;">
                                    <p class="clickable-text">{{ photo.showComments ? 'Nascondi Commenti' : 'Mostra Commenti' }}</p>
                                </button>
                                <div v-if="photo.showComments" class="comments">
                                    <div v-for="comment in photo.listComment" :key="comment.ID" style="display: flex; align-items: center;">
                                        <div class="comment-content">
                                            <p><strong>{{ comment.UserUsername }}:</strong> {{ comment.Text }}</p>
                                        </div>
                                        <button v-if="comment.User_id == this.user.ID" class="btn delete-button" @click="deleteComment(photo.ID, comment.ID)">
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
            changeUsername: false,
            newUsername: "",
            user: {
                ID: localStorage.getItem("token"),
                Username: localStorage.getItem("username"),
            },
            profile: {
                NumFollow: 0,
                NumFollowing: 0,
                NumPhoto: 0,
                ListPhoto: []
            }
        };
    },
    methods: {
        async usernameChange() {
            if (this.newUsername == "") {
                this.errormsg = "il campo username è vuoto"
            } else {
                try {
                let response = await this.$axios.put("/users/" + this.user.Username + "/username", 
                    { username: this.newUsername },
                    { headers: { Authorization: `Bearer ${this.user.ID}` } }
                    )
                this.user = response.data
                localStorage.removeItem("username")
                localStorage.setItem("username", this.user.Username)
                this.getProfile();
                this.changeUsername = false;
                this.newUsername = "";
                }
                catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again";
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "Server error, please try again later";
                } else {
                    this.errormsg = e.toString();
                }
                }
            }
            },
        async toggleComments(photo) {
            photo.showComments = !photo.showComments;
        },
        async toggleChangeUsername(photo) {
            this.changeUsername = !this.changeUsername;
        },
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
                if (this.profile.ListPhoto != null) {
                    for (let i = 0; i < this.profile.ListPhoto.length; i++) {
                        try {
                            const isLike = await this.$axios.get(`/users/${this.user.Username}/photo/${this.profile.ListPhoto[i].ID}/like/${this.user.Username}`, {
                                headers: { Authorization: `Bearer ${this.user.ID}` }
                            });
                            let like = isLike.data;
                            if (like.User_id == 0) {
                                this.profile.ListPhoto[i].liked = false;
                            }
                            else {
                                this.profile.ListPhoto[i].liked = true;
                            }
                            this.errormsg = null;
                            // Aggiungi una proprietà per gestire la visualizzazione dei commenti
                            this.profile.ListPhoto[i].showComments = false;
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
                    this.profile.ListPhoto = [];
                }

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
        async deleteComment(photoID, commentID) {
            try {
                await this.$axios.delete(`/users/${this.user.Username}/photo/${photoID}/comment/${commentID}`, {
                    headers: { Authorization: `Bearer ${this.user.ID}` }
                });
                this.errormsg = null; // Resetta il messaggio di errore
                this.getProfile();
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
                this.getProfile();
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
    },
    mounted() {
        this.getProfile();
    }
};
</script>

<style scoped>

.buttone {
  padding: 10px 20px;
  font-size: 16px;
  background-color: orangered;
  color: black;
  border: none;
  cursor: pointer;
  width: 100%;
  /* Imposta la larghezza del bottone */
  max-width: fit-content;
  /* Larghezza massima del bottone */
  border-radius: 10px;
  /* Arrotonda i bordi del bottone */
}

input[type="text"] {
  padding: 10px;
  font-size: 16px;
  margin-bottom: 10px;
  /* Aggiungi margine inferiore per spaziatura */
  width: 100%;
  /* Imposta la larghezza dell'input */
  max-width: 300px;
  /* Larghezza massima dell'input */
  border-radius: 10px;
  /* Arrotonda i bordi dell'input */
}

form {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  /* Dispone gli elementi del form in colonna */
  align-items: center;
  /* Centra gli elementi del form */
}

.clickable-text {
    font-weight: bold;
    cursor: pointer;
    transition: color 0.3s;
}

.clickable-text:hover {
    color: rgba(0, 0, 0, 0.5); /* Schiarisce il testo al passaggio del mouse */
}

.comment-wrapper {
    display: flex; /* Utilizza Flexbox */
    align-items: center; /* Allinea verticalmente al centro */
}

.comment-content {
    flex: 1; /* Fai espandere il contenuto del commento per riempire lo spazio disponibile */
    margin-right: 10px; /* Aggiungi uno spazio a destra tra il contenuto del commento e il pulsante */
}

.delete-button {
    background-color: transparent; /* Imposta lo sfondo trasparente */
    border: none; /* Rimuovi il bordo */
    padding: 0; /* Rimuovi il padding */
    cursor: pointer; /* Mostra il cursore come puntatore */
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
}</style>
