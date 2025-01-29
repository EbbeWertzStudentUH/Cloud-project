import { POST } from "$lib";


export async function addUserToProject(user_id, project_id) {
    await POST({user_id}, '/project/'+project_id+'/user', false)
}
