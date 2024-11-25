import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHashHistory } from 'vue-router';
import ServerPage from './pages/ServerPage.vue';
import Home from './pages/Home.vue';
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'
import AddServer from './pages/AddServer.vue';
const routes = [
    { path: '/', component: Home },
    { name: 'server', path: '/server/:id', component: ServerPage },
    { path: '/add', component: AddServer}
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

createApp(App)
    .use(router)
    .mount('#app');
