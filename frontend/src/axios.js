import axios from 'axios'

export async function getConfigs(search) {
  let config = getTokenConfig()
  config.params = {
    search: search
  }
  return await axios.get('/dhcp', config)
}

export async function createConfig(cfg) {
  return await axios.post('/dhcp', cfg, getTokenConfig())
}

export async function updateConfig(cfg) {
  return await axios.put('/dhcp/' + cfg.id, cfg, getTokenConfig())
}

export async function getConfigFor(id) {
  return await axios.get('/dhcp/' + id, getTokenConfig())
}

export async function deleteConfigFor(id) {
  return await axios.delete('/dhcp/' + id, getTokenConfig())
}

export async function resetPayments() {
  return await axios.post('/dhcp/resetPayment', {}, getTokenConfig())
}

export async function login(credentials) {
  return await axios.post('/dhcp/login', credentials)
}

export async function updateDhcp() {
  return await axios.post('/dhcp/write', {}, getTokenConfig())
}

export async function getShameList() {
  return await axios.get('/dhcp/shame', getTokenConfig())
}

function getTokenConfig() {
  return {
    headers: {
      Authorization: 'Bearer ' + localStorage.getItem('jwt')
    }
  }
}
