import type {AsyncData} from "#app";

function getBaseURL() {
    const config = useRuntimeConfig()
    return config.public.baseURL
}

export function getConfigs(filters: ManageFilters): AsyncData<MemberConfig[], any> {
    let reqHeaders = authHeader()
    if (filters.search === '') {
        filters.search = null
    }
    return useFetch('/dhcp', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: reqHeaders,
            params: {
                search: filters.search,
                hasPaid: filters.payment,
                disabled: filters.disabled
            },
            server: false,
            transform: jsonTransform<MemberConfig[]>
        }
    )
}

export function createConfig(cfg: MemberInput): AsyncData<any, any> {
    return useFetch('/dhcp', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function updateConfig(cfg: MemberInput): AsyncData<any, any> {
    return useFetch('/dhcp/' + cfg.id, {
            baseURL: getBaseURL(),
            method: "PUT",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function getConfigFor(id: string): AsyncData<MemberConfig, any> {
    return useFetch('/dhcp/' + id, {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            transform: jsonTransform<MemberConfig>
        }
    )
}

export function deleteConfigFor(id: number): AsyncData<any, any> {
    return useFetch('/dhcp/' + id, {
            baseURL: getBaseURL(),
            method: "DELETE",
            headers: authHeader(),
        }
    )
}

export function resetPayments(): AsyncData<any, any> {
    return useFetch('/dhcp/resetPayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function loginUser(credentials: Credentials): AsyncData<LoginResponse, any> {
    return useFetch('/dhcp/login', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: credentials,
            transform: jsonTransform<LoginResponse>
        }
    );
}

export function updateDhcp(): AsyncData<any, any> {
    return useFetch('/dhcp/write', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function getShameList(): AsyncData<any, any> {
    return useFetch('/dhcp/shame', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            transform: jsonTransform<ReducedPerson[]>
        }
    )
}

export function togglePayment(id: number): AsyncData<any, any> {
    return useFetch('/dhcp/' + id + '/togglePayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function fetchRooms(filters: RoomFilters): AsyncData<Room[], any> {
    return useFetch('/api/rooms', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            params: {
                occupied: filters.occupied,
                block: filters.block
            },
            transform: jsonTransform<Room[]>
        }
    )
}

function authHeader(): any {
    return {
        "Authorization": 'Bearer ' + localStorage.getItem('jwt')
    }
}

function jsonTransform<T>(r: string): T {
    return JSON.parse(r)
}
