const Home = { template: '<div>Home Page</div>' }
const About = { template: '<div>About Page</div>' }

const routes = [
    {path: '/', component:Home},
    {path: '/about', component:About},
]

const router = new VueRouter({
    mode: 'history',
    routes,
})

var app = new Vue( {
    router,
}).$mount('#app')

