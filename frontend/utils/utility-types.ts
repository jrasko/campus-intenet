export interface RoomFilters {
    occupied?: boolean
    search?: string
    payment?: boolean
    disabled?: boolean
    wg?: string
    block?: Block[]
}

export type Block = ('1' | '2' | '3' | '4' | '5') 

export interface ServerFilters {
    disabled: boolean | null
    server: boolean | null
}
export type Columns = Record<Column, ColumnOptions>

export type Column =
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
  'movedIn' |
  'nationality' |
  'createdAt' |
  'updatedAt' |
  'lastEditor' |
  'isFurnished' |
  'occupied'

export interface ColumnOptions {
    key: string
    header: string
    field: string
    kind: ('text' | 'date' | 'bool')
    banNull?: boolean
}
