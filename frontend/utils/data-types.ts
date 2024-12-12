export interface Room {
  roomNr: string
  wg: string
  block: string
  member?: MemberConfig
}

export interface Server {
  id: number,
  name: string,
  mac: string,
  ip: string,
  disabled: boolean,
}

export interface MemberConfig {
    [index: string]: any

    id: number
    firstname: string
    lastname: string
    room: {
        roomNr: string
        wg: string
    }
    dhcpConfig: {
        id: number
        mac: string
        ip: string
        disabled: boolean
        manufacturer?: string
    }
    isFurnished: boolean
    email: string
    phone: string
    movedIn: string
    nationality: string
    comment: string
    hasPaid: boolean
    createdAt?: string
    updatedAt?: string
    lastEditor?: string
}

export interface InputMember {
    id: number
    firstname: string
    lastname: string
    roomNr: string
    isFurnished: boolean
    dhcpConfig: {
        id: number
        mac: string
        ip: string
        disabled: boolean
    }
    nationality: string
    email: string
    phone: string
    hasPaid: boolean
    comment: string
    movedIn: string
}
