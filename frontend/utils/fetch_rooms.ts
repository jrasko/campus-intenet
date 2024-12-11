import {authHeader, getBaseURL} from "~/utils/utils";

export function fetchRooms(filters: RoomFilters): Promise<Room[]> {
    return $fetch<Room[]>('/api/rooms', {
            baseURL: getBaseURL(),
            method: "GET",
            headers: authHeader(),
            params: {
                occupied: filters.occupied,
                search: filters.search,
                disabled: filters.disabled,
                wg: filters.wg,
                payment: filters.payment
            },
            parseResponse: jsonTransform<Room[]>
        }
    )
}
