export const useAuth = async () => {
    let token = await nextTick(() => {
        return localStorage.getItem('jwt')
    })
    return {
        "Authorization": 'Bearer ' + token
    }
}