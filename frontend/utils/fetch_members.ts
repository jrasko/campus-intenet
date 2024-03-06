import type {AsyncData} from "#app";
import {authHeader, getBaseURL} from "~/utils/utils";

export async function getMemberConfigs(filters: ManageFilters): Promise<AsyncData<MemberConfig[], any>> {
    if (filters.search === '') {
        filters.search = null
    }
    return $fetch('/api/members', {
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

export function createMemberConfig(cfg: InputMember): Promise<AsyncData<any, any>> {
    return $fetch('/api/members', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function updateMemberConfig(cfg: InputMember): Promise<AsyncData<any, any>> {
    return $fetch('/api/members/' + cfg.id, {
            baseURL: getBaseURL(),
            method: "PUT",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function getMemberConfigFor(id: string): Promise<AsyncData<MemberConfig, any>> {
    return $fetch('/api/members/' + id, {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            parseResponse: jsonTransform<MemberConfig>
        }
    )
}

export function deleteMemberConfigFor(id: number): Promise<AsyncData<any, any>> {
    return $fetch('/api/members/' + id, {
            baseURL: getBaseURL(),
            method: "DELETE",
            headers: authHeader(),
        }
    )
}

export function resetPayments(): Promise<AsyncData<any, any>> {
    return $fetch('/api/members/resetPayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function getShameList(): Promise<AsyncData<any, any>> {
    return $fetch('/api/shame', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            parseResponse: jsonTransform<ReducedPerson[]>
        }
    )
}

export function togglePayment(id: number): Promise<AsyncData<any, any>> {
    return $fetch('/api/members/' + id + '/togglePayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function loginUser(credentials: Credentials): Promise<AsyncData<LoginResponse, any>> {
    return $fetch('/api/login', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: credentials,
            parseResponse: jsonTransform<LoginResponse>
        }
    );
}
