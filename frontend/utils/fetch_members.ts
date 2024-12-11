import {authHeader, getBaseURL} from "~/utils/utils";

export function createMemberConfig(cfg: InputMember): Promise<any> {
    return $fetch('/api/members', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function updateMemberConfig(cfg: InputMember): Promise<any> {
    return $fetch('/api/members/' + cfg.id, {
            baseURL: getBaseURL(),
            method: "PUT",
            headers: authHeader(),
            body: cfg,
        }
    )
}

export function getMemberConfigFor(id: string): Promise<MemberConfig> {
    return $fetch<MemberConfig>('/api/members/' + id, {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            parseResponse: jsonTransform<MemberConfig>
        }
    )
}

export function deleteMemberConfigFor(id: number): Promise<any> {
    return $fetch('/api/members/' + id, {
            baseURL: getBaseURL(),
            method: "DELETE",
            headers: authHeader(),
        }
    )
}

export function resetPayments(): Promise<any> {
    return $fetch('/api/members/resetPayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function togglePayment(id: number): Promise<any> {
    return $fetch('/api/members/' + id + '/togglePayment', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function toggleNetworkActivation(id: number): Promise<any> {
    return $fetch('/api/net-configs/' + id + '/toggleActivation', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function punish(): Promise<any> {
    return $fetch('/api/members/punish', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    ) 
}

export function loginUser(credentials: Credentials): Promise<LoginResponse> {
    return $fetch<LoginResponse>('/api/login', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: credentials,
            parseResponse: jsonTransform<LoginResponse>
        }
    );
}
