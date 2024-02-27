interface Credentials {
    username: string,
    password: string
}

interface LoginResponse {
    token: string,
    role: string,
    username: string
}

interface ManageFilters {
    search: string | null,
    payment: boolean | null
    disabled: boolean | null
}

type Block = 1 | 2 | 3 | 4 | 5

interface RoomFilters {
    occupied: boolean | null,
    block: Block[]
}

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
        mac: string,
        ip: string,
        disabled: boolean,
        manufacturer?: string
    }
    email: string,
    phone: string
    comment: string,
    hasPaid: boolean,
    createdAt?: string
    updatedAt?: string
    lastEditor?: string
}

interface ReducedPerson {
    firstname: string,
    lastname: string
}

interface MemberInput {
    id: number,
    firstname: string,
    lastname: string,
    roomNr: string,
    dhcpConfig: {
        mac: string,
        ip: string,
        disabled: boolean,
    }
    email: string,
    phone: string
    hasPaid: boolean,
    comment: string,
}

type ColumnFormat = 'text' | 'date'

interface ColumnOptions {
    key: string,
    header: string,
    field: string,
    kind: ColumnFormat
}

type Column =
    'firstname' |
    'lastname' |
    'mac' |
    'roomNr' |
    'wg' |
    'email' |
    'phone' |
    'ip' |
    'manufacturer' |
    'comment' |
    'createdAt' |
    'updatedAt' |
    'lastEditor'

type Columns = Record<Column, ColumnOptions>

interface ManageFilterList {
    payment: { header: string, value: boolean | null }[]
    disabled: { header: string, value: boolean | null }[]
}
interface RoomFilterList {
    occupied: { header: string, value: boolean | null }[]
    block: { header: string, value: Block | null }[]
}

interface Room {
    roomNr: string
    wg: string
    block: string
    member?: MemberConfig
}
