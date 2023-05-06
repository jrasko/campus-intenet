import axios from "axios";

const basePath = 'http://localhost:8000'

export async function getConfigs() {
    let resp = await axios
        .get(basePath + '/dhcpd')
        .catch(err => {
            console.log(err)
            alert("Fehler: " + err)
        })
    return resp.data
}

export async function updateConfig(cfg) {
    return await axios.put(basePath + '/dhcpd', cfg);
}