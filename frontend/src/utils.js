export function formatMac(mac) {
  const macLen = 17
  let str = mac.toUpperCase()
  if (str.length < macLen) {
    str = str.replace(/([0-9A-F]{2}$)/g, '$1:')
  }
  return str
}

export const noAuthPages = ['login']

export function isLoggedIn() {
  return localStorage.getItem('jwt') != null
}
