import type {AsyncData} from "#app";
import {authHeader, getBaseURL} from "~/utils/utils";

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
