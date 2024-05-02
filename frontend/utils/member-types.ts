interface MemberConfig {
    [index: string]: any,

    id: number,
    firstname: string,
    lastname: string,
    room: {
        roomNr: string,
        wg: string
    }
    dhcpConfig: {
        id: number,
        mac: string,
        ip: string,
        disabled: boolean,
        manufacturer?: string
    }
    email: string,
    phone: string,
    movedIn: string,
    comment: string,
    hasPaid: boolean,
    createdAt?: string
    updatedAt?: string
    lastEditor?: string
}

interface InputMember {
    id: number,
    firstname: string,
    lastname: string,
    roomNr: string,
    dhcpConfig: {
        id: number,
        mac: string,
        ip: string,
        disabled: boolean,
    }
    email: string,
    phone: string
    hasPaid: boolean,
    comment: string,
    movedIn: string
}

interface ReducedPerson {
    firstname: string,
    lastname: string
}