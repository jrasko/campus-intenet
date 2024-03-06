import type {AsyncData} from "#app";
import {authHeader, getBaseURL} from "~/utils/utils";


export function updateDhcp(): Promise<any> {
    return $fetch('/api/write', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
        }
    )
}

export function listServers(filters: ServerFilters): Promise<AsyncData<Server[], any>> {
    return $fetch('/api/net-configs', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            params: {
                server: filters.server,
                disabled: filters.disabled
            },
            parseResponse: jsonTransform<Server[]>
        }
    )
}

export function fetchServer(id: string): Promise<AsyncData<Server[], any>> {
    return $fetch('/api/net-configs/' + id, {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            parseResponse: jsonTransform<Server[]>
        }
    )
}

export function createServer(cfg: Server): Promise<any>{
    return $fetch('/api/net-configs', {
            baseURL: getBaseURL(),
            method: "POST",
            headers: authHeader(),
            body: cfg
        }
    )
}
export function updateServer(id: number, cfg: Server): Promise<any>{
    return $fetch('/api/net-configs/' + id, {
            baseURL: getBaseURL(),
            method: "PUT",
            headers: authHeader(),
            body: cfg
        }
    )
}

export function deleteServer(id: number): Promise<any>{
    return $fetch('/api/net-configs/' + id, {
            baseURL: getBaseURL(),
            method: "DELETE",
            headers: authHeader(),
        }
    )
}
