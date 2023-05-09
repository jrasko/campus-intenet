import axios from 'axios'

export async function getConfigs() {
    let resp = await axios.get('/dhcpd').catch((err) => {
        console.log(err)
        alert('Fehler: ' + err)
    })
    return resp.data
}

export async function updateConfig(cfg) {
    return await axios.put('/dhcpd', cfg)
}

export async function getConfigFor(mac) {
    return await axios.get('/dhcpd/' + mac)
}

export async function deleteConfigFor(mac) {
    return await axios.delete('/dhcpd/' + mac)
}

export async function resetPayments() {
    return await axios.post('/dhcpd/resetPayment')
}

export async function login(credentials) {
    return await axios.post('/dhcpd/login', credentials)
}
