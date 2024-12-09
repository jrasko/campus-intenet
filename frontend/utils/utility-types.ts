interface ManageFilters {
    search: string | null
    payment: boolean | null
    disabled: boolean | null
    wg: string | null
}

interface ManageFilterList {
    payment: { header: string, value: boolean | null }[]
    disabled: { header: string, value: boolean | null }[]
}

type Block = '1' | '2' | '3' | '4' | '5'

interface RoomFilters {
    occupied: boolean | null,
    block: Block[]
}

interface RoomFilterList {
    occupied: { header: string, value: boolean | null }[]
    block: { header: string, value: Block | null }[]
}

interface ServerFilters {
    disabled: boolean | null
    server: boolean | null
}

interface ServerFilterList {
    disabled: { header: string, value: boolean | null }[]
    server: { header: string, value: boolean | null }[]
}

type ColumnFormat = 'text' | 'date' | 'bool'

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
    'movedIn' |
    'nationality' |
    'createdAt' |
    'updatedAt' |
    'lastEditor' |
    'isFurnished'

type Columns = Record<Column, ColumnOptions>
