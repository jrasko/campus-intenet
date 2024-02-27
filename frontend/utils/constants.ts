export const tableData: Columns = {
    firstname: {
        header: 'Vorname',
        field: 'firstname',
        key: 'firstname',
        kind: 'text'
    },
    lastname: {
        header: 'Nachname',
        field: 'lastname',
        key: 'lastname',
        kind: 'text'
    },
    mac: {
        header: 'MAC',
        field: 'dhcpConfig.mac',
        key: 'mac',
        kind: 'text'
    },
    roomNr: {
        key: 'roomNr',
        header: 'Zimmer-Nr.',
        field: 'room.roomNr',
        kind: 'text'
    },
    wg: {
        header: 'WG',
        field: 'room.wg',
        key: 'wg',
        kind: 'text'
    },
    email: {
        header: 'E-Mail',
        field: 'email',
        key: 'email',
        kind: 'text'
    },
    phone: {
        header: 'Telefonnr.',
        field: 'phone',
        key: 'phone',
        kind: 'text'
    },
    ip: {
        header: 'IP',
        field: 'dhcpConfig.ip',
        key: 'ip',
        kind: 'text'
    },
    manufacturer: {
        header: 'Hersteller',
        field: 'dhcpConfig.manufacturer',
        key: 'manufacturer',
        kind: 'text'
    },
    comment: {
        header: 'Kommentar',
        field: 'comment',
        key: 'comment',
        kind: 'text'
    },
    createdAt: {
        header: 'Erstellt',
        field: 'createdAt',
        key: 'createdAt',
        kind: 'date'
    },
    updatedAt: {
        header: 'Bearbeitet',
        field: 'updatedAt',
        key: 'updatedAt',
        kind: 'date'
    },
    lastEditor: {
        header: 'Editor',
        field: 'lastEditor',
        key: 'lastEditor',
        kind: 'text'
    }
}

export const manageFilter: ManageFilterList = {
    payment: [
        {
            header: 'Alle',
            value: null
        },
        {
            header: 'Bezahlt',
            value: true
        },
        {
            header: 'Nicht Bezahlt',
            value: false
        }
    ],
    disabled: [
        {
            header: 'Alle',
            value: null
        },
        {
            header: 'Aktiviert',
            value: false
        },
        {
            header: 'Deaktiviert',
            value: true
        }
    ]
}

export const roomFiler: RoomFilterList = {
    block: [
        {
            header: 'Alle',
            value: null
        }
    ],
    occupied: [
        {
            header: 'Alle',
            value: null
        },
        {
            header: 'Belegt',
            value: true
        },
        {
            header: 'Unbelegt',
            value: false
        }
    ]
}
