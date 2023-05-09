import axios from 'axios'

export async function getConfigs() {
  return await axios.get('/dhcpd', getTokenConfig())
}

export async function updateConfig(cfg) {
  return await axios.put('/dhcpd', cfg, getTokenConfig())
}

export async function getConfigFor(mac) {
  return await axios.get('/dhcpd/' + mac, getTokenConfig())
}

export async function deleteConfigFor(mac) {
  return await axios.delete('/dhcpd/' + mac, getTokenConfig())
}

export async function resetPayments() {
  return await axios.post('/dhcpd/resetPayment', getTokenConfig())
}

export async function login(credentials) {
  return await axios.post('/dhcpd/login', credentials)
}

function getTokenConfig() {
  return {
    headers: {
      Authorization: 'Bearer ' + localStorage.getItem('jwt')
    }
  }
}
