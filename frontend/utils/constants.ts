export const tableData: Columns = {
    occupied: {
      header: 'Belegt',
      field: 'member',
      key: 'occupied',
      kind: 'bool',
      banNull: false
    },
    roomNr: {
      key: 'roomNr',
      header: 'Zimmer-Nr.',
      field: 'roomNr',
      kind: 'text'
    },
    wg: {
      header: 'WG',
      field: 'wg',
      key: 'wg',
      kind: 'text'
    },
    firstname: {
        header: 'Vorname',
        field: 'member.firstname',
        key: 'firstname',
        kind: 'text'
    },
    lastname: {
        header: 'Nachname',
        field: 'member.lastname',
        key: 'lastname',
        kind: 'text'
    },
    comment: {
      header: 'Kommentar',
      field: 'member.comment',
      key: 'comment',
      kind: 'text'
    },
    mac: {
        header: 'MAC',
        field: 'member.dhcpConfig.mac',
        key: 'mac',
        kind: 'text'
    },
    email: {
        header: 'E-Mail',
        field: 'member.email',
        key: 'email',
        kind: 'text'
    },
    phone: {
        header: 'Telefonnr.',
        field: 'member.phone',
        key: 'phone',
        kind: 'text'
    },
    ip: {
        header: 'IP',
        field: 'member.dhcpConfig.ip',
        key: 'ip',
        kind: 'text'
    },
    isFurnished: {
        header: 'Möbliert',
        field: 'member.isFurnished',
        key: 'isFurnished',
        kind: 'bool',
        banNull: true
    },
    manufacturer: {
        header: 'Hersteller',
        field: 'member.dhcpConfig.manufacturer',
        key: 'manufacturer',
        kind: 'text'
    },
    movedIn: {
        header: 'Einzugsdatum',
        field: 'member.movedIn',
        key: 'movedIn',
        kind: 'text'
    },
    nationality: {
      header: 'Nationalität',
      field: 'member.nationality',
      key: 'nationality',
      kind: 'text'
    },
    createdAt: {
        header: 'Erstellt',
        field: 'member.createdAt',
        key: 'createdAt',
        kind: 'date'
    },
    updatedAt: {
        header: 'Bearbeitet',
        field: 'member.updatedAt',
        key: 'updatedAt',
        kind: 'date'
    },
    lastEditor: {
        header: 'Editor',
        field: 'member.lastEditor',
        key: 'lastEditor',
        kind: 'text'
    }
}

interface ManageFilterList {
  payment: { header: string, value: boolean | null }[]
  disabled: { header: string, value: boolean | null }[]
  occupied: { header: string, value: boolean | null }[]
  block: { header: string, value: Block | null }[]  
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
    ],
    block: [],
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

interface ServerFilterList {
  disabled: { header: string, value: boolean | null }[]
  server: { header: string, value: boolean | null }[]
}

export const serverFilter: ServerFilterList = {
    server: [
        {
            header: 'Alle',
            value: null
        },
        {
            header: 'Server',
            value: true
        },
        {
            header: 'Mitglieder',
            value: false
        }
    ],
    disabled: [
        {
            header: 'Alle',
            value: null
        },
        {
            header: 'Deaktiviert',
            value: true
        },
        {
            header: 'Aktiviert',
            value: false
        }
    ]
}
