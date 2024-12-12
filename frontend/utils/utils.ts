export const roles = {
    admin: 'admin',
    editor: 'editor',
    financer: 'financer',
    viewer: 'viewer'
}

export function formatMac(str: string): string {
    str = str.toUpperCase()
    str = str.replace(/[^0-9A-F]/g, '')

    let out = ""
    for (let i = 1; i <= Math.min(str.length, 12); i++) {
        out += str[i - 1]
        if (i % 2 == 0 && i < str.length) {
            out += ":"
        }
    }
    return out
}

export function toInputMember(i: MemberConfig): InputMember {
    return {
        id: i.id,
        comment: i.comment,
        dhcpConfig: {
            id: i.dhcpConfig.id,
            ip: i.dhcpConfig.ip,
            mac: i.dhcpConfig.mac,
            disabled: i.dhcpConfig.disabled,
        },
        email: i.email,
        firstname: i.firstname,
        hasPaid: i.hasPaid,
        lastname: i.lastname,
        phone: i.phone,
        roomNr: i.room.roomNr,
        movedIn: i.movedIn,
        nationality: i.nationality,
        isFurnished: i.isFurnished
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
