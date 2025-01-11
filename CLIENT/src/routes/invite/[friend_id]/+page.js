import { GET, POSTWithToken } from "$lib";


export async function load({ params }) {
    const { friend_id } = params;
    const friend_resp = await GET(`/user?id=${friend_id}`);
    return {
        friend: friend_resp,
        sendRequest
    };
}


async function sendRequest(friend_id){
    await POSTWithToken({friend_id}, '/user/friend-requests/send');
}