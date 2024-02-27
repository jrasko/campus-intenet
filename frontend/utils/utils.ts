export function formatMac(mac: string): string {
    const macLen = 17
    let str = mac.toUpperCase()
    if (str.length < macLen) {
        str = str.replace(/([0-9A-F]{2}$)/g, '$1:')
    }
    return str
}

export function MemberCompare(field: keyof MemberConfig) {
    return function (a: MemberConfig, b: MemberConfig) {
        let aVal = a[field]
        let bVal = b[field]
        if (typeof aVal === 'string') {
            return aVal.localeCompare(<string>bVal)
        }
        return aVal === bVal ? 0 : aVal ? -1 : 1
    }
}
export function toInputMember(i: MemberConfig):MemberInput{
    return {
        id: i.id,
        comment: i.comment,
        dhcpConfig: i.dhcpConfig,
        email: i.email,
        firstname: i.firstname,
        hasPaid: i.hasPaid,
        lastname: i.lastname,
        phone: i.phone,
        roomNr: i.room.roomNr
    }
}
