export function formatMac(mac) {
  const macLen = 17
  if (mac < macLen) {
    mac = this.person.mac.replace(/:/g, '').replace(/(.{2})/g, '$1:')
  }
  return mac
}
