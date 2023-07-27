import axios from 'axios'

export async function getConfigs(search) {
  let config = getTokenConfig();
  config.params = {
    search: search,
  }
  console.log(config)
  return await axios.get('/dhcpd', config)
}

export async function createConfig(cfg) {
  return await axios.post('/dhcpd', cfg, getTokenConfig())
}

export async function updateConfig(cfg) {
  return await axios.put('/dhcpd/' + cfg.id, cfg, getTokenConfig())
}

export async function getConfigFor(id) {
  return await axios.get('/dhcpd/' + id, getTokenConfig())
}

export async function deleteConfigFor(id) {
  return await axios.delete('/dhcpd/' + id, getTokenConfig())
}

export async function resetPayments() {
  return await axios.post('/dhcpd/resetPayment', {}, getTokenConfig())
}

export async function login(credentials) {
  return await axios.post('/dhcpd/login', credentials)
}

export async function updateDhcpd() {
  return await axios.post('/dhcpd/write', {}, getTokenConfig())
}

export async function getShameList() {
  return await axios.get('/dhcpd/shame')
}

function getTokenConfig() {
  return {
    headers: {
      Authorization: 'Bearer ' + localStorage.getItem('jwt')
    }
  }
}
