import type {AsyncData} from "#app";
import {authHeader} from "~/utils/utils";

export function getBaseURL() {
    const config = useRuntimeConfig()
    return config.public.baseURL
}

export async function getConfigs(filters: ManageFilters): Promise<AsyncData<MemberConfig[], any>> {
    if (filters.search === '') {
        filters.search = null
    }
    return $fetch('/dhcp', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            params: {
                search: filters.search,
                hasPaid: filters.payment,
                disabled: filters.disabled
            },
            server: false,
            parseResponse: jsonTransform<MemberConfig[]>
        }
    )
}

export function createConfig(cfg: InputMember): Promise<AsyncData<any, any>> {
    return $fetch('/dhcp', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function updateConfig(cfg: InputMember): Promise<AsyncData<any, any>> {
    return $fetch('/dhcp/' + cfg.id, {
            baseURL: getBaseURL(),
            method: "PUT",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function getConfigFor(id: string): Promise<AsyncData<MemberConfig, any>> {
    return $fetch('/dhcp/' + id, {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            parseResponse: jsonTransform<MemberConfig>
        }
    )
}

export function deleteConfigFor(id: number): Promise<AsyncData<any, any>> {
    return $fetch('/dhcp/' + id, {
            baseURL: getBaseURL(),
            method: "DELETE",
            headers: authHeader(),
        }
    )
}

export function resetPayments(): Promise<AsyncData<any, any>> {
    return $fetch('/dhcp/resetPayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function loginUser(credentials: Credentials): Promise<AsyncData<LoginResponse, any>> {
    return $fetch('/dhcp/login', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: credentials,
            parseResponse: jsonTransform<LoginResponse>
        }
    );
}

export function updateDhcp(): Promise<AsyncData<any, any>> {
    return $fetch('/dhcp/write', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function getShameList(): Promise<AsyncData<any, any>> {
    return $fetch('/dhcp/shame', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            parseResponse: jsonTransform<ReducedPerson[]>
        }
    )
}

export function togglePayment(id: number): Promise<AsyncData<any, any>> {
    return $fetch('/dhcp/' + id + '/togglePayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function fetchRooms(filters: RoomFilters): Promise<AsyncData<Room[], any>> {
    return $fetch('/api/rooms', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            params: {
                occupied: filters.occupied,
                block: filters.block
            },
            parseResponse: jsonTransform<Room[]>
        }
    )
}
