
export function formatMac(str: string): string {
    str = str.toUpperCase()
    str = str.replace(/[^0-9A-F]/g, '')

    let out = ""
    for (let i = 1; i <= Math.min(str.length, 12); i++) {
        out += str[i-1]
        if (i%2 == 0 && i < str.length){
            out += ":"
        }
    }
    return out
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

export function toInputMember(i: MemberConfig): InputMember {
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

export function jsonTransform<T>(r: string): T | null {
    if (r == '') {
        return null
    }
    return JSON.parse(r)
}

export function authHeader() {
    return {
        "Authorization": 'Bearer ' + localStorage.getItem('jwt')
    }
}

export function getBaseURL() {
    const config = useRuntimeConfig()
    return config.public.baseURL
}
