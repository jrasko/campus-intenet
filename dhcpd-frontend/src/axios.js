import axios from 'axios'

const basePath = 'http://localhost:8000'

export async function getConfigs() {
  let resp = await axios.get(basePath + '/dhcpd').catch((err) => {
    console.log(err)
    alert('Fehler: ' + err)
  })
  return resp.data
}

export async function updateConfig(cfg) {
  return await axios.put(basePath + '/dhcpd', cfg)
}

export async function getConfigFor(mac) {
  return await axios.get(basePath + '/dhcpd/' + mac)
}

export async function deleteConfigFor(mac) {
  return await axios.delete(basePath + '/dhcpd/' + mac)
}

export async function resetPayments() {
  return await axios.post(basePath + '/dhcpd/resetPayment')
}
