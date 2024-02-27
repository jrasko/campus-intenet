export default defineNuxtRouteMiddleware((to, from) => {
    if (!isAuthenticated()) {
        return navigateTo('/login')
    }
})

function isAuthenticated(): boolean {
    return localStorage.getItem('jwt') != null
}

