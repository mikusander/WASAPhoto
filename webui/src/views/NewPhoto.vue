<template>
    <div>
        <header class="navbar navbar-dark sticky-top bg-dark" style="height: 80px;">
            <div class="container-fluid d-flex justify-content-between align-items-center">
                <div class="align-items-center text-center flex-grow-1">
                    <a @click="homePage" style="color: orangered; font-size: 70px; font-weight: bold; cursor: pointer;">WASA
                        Photo</a>
                </div>
                <div class="d-flex align-items-center">
                    <a @click="myAccount" class="text-white" style="margin-left: 20px;">
                        <img src="../images/utente.png" alt="Immagine Utente"
                            style="width: 70px; height: 70px; border-radius: 50%; cursor: pointer;">
                    </a>
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
            <div v-if="imagePreview" class="image-preview">
                <img class="img" :src="imagePreview" alt="Anteprima immagine" />
            </div>
            <div class="d-flex align-items-center" style="margin-top: 30px; height: 80px;">
                <form @submit.prevent="handleSubmit" class="d-flex flex-column align-items-start ml-3"
                    enctype="multipart/form-data">
                    <input type="file" ref="imageInput" class="form-control-file mb-3" accept="image/*"
                        @change="previewImage" required>
                    <textarea v-model="description" class="form-control form-control-lg form-control-auto-resize mb-3"
                        placeholder="Descrizione" required></textarea>
                    <button type="submit" class="btn btn-primary btn-lg">Carica</button>
                </form>
            </div>
            <SuccessMsg v-if="successmsg" :msg="successmsg"></SuccessMsg>
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
            successmsg: null,
            loading: false,
            some_data: null,
            description: "",
            imagePreview: null, // Aggiungi questa proprietà per l'anteprima dell'immagine
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
        async myAccount() {
            this.$router.push({ path: '/users/' + this.user.Username + '/MyAccount' });
        },
        async homePage() {
            this.$router.push({ path: '/users/' + this.user.Username + '/stream' });
        },
        sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
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
                        this.errormsg = null; // Resetta il messaggio di errore
                        this.successmsg = "foto caricata con successo";
                        await this.sleep(1000);
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
        },
        previewImage(event) {
            const file = event.target.files[0];
            if (file) {
                const reader = new FileReader();
                reader.onload = e => {
                    this.imagePreview = e.target.result; // Salva l'URL dell'immagine in base64
                };
                reader.readAsDataURL(file); // Legge il file come Data URL
            }
        },
        async getProfile() {
            try {
                const response = await this.$axios.get("/users/" + this.user.Username + "/profile",
                    { headers: { Authorization: `Bearer ${this.user.ID}` } }
                );
                this.profile = response.data;
                this.errormsg = null; // Resetta il messaggio di errore
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
    height: 80px;
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

.form-control-auto-resize {
    resize: none;
    overflow: hidden;
    min-height: 100px;
    width: 1000px;
    /* Modifica secondo necessità */
}

.image-preview {
    text-align: center;
}

.image-preview img {
    background-color: white;
    border-radius: 15px;
    padding: 30px;
    /* Aumenta il padding per rendere lo sfondo più grosso */
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    display: inline-block;
    border: 2px solid #ccc;
    /* Aggiungi un bordo per un effetto più definito */
    width: 300px;
    /* Imposta una larghezza fissa */
    height: auto;
    /* Imposta un'altezza fissa */
    overflow: hidden;
    /* Nasconde l'eccesso dell'immagine se esce dal riquadro */
}

.img {
    width: 100%;
    /* Rende l'immagine larga quanto il contenitore */
    height: 100%;
    /* Rende l'immagine alta quanto il contenitore */
    object-fit: cover;
    /* Ritaglia l'immagine per riempire il contenitore mantenendo il centro */
    border-radius: 10px;
}</style>

  