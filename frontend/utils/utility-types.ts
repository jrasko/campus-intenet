interface ManageFilters {
    search: string | null,
    payment: boolean | null
    disabled: boolean | null
}

type Block = '1' | '2' | '3' | '4' | '5'

interface RoomFilters {
    occupied: boolean | null,
    block: Block[]
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

