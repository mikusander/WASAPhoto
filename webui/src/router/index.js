import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'
import MyAccount from '../views/MyAccount.vue'
import NewPhoto from '../views/NewPhoto.vue'
import Profile from '../views/Profile.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/users/:username/stream', component: StreamView},
		{path: '/users/:username/MyAccount', component: MyAccount},
		{path: '/users/:username/newPhoto', component: NewPhoto},
		{path: '/users/:username/profile', component: Profile}
	]
})

export default router
