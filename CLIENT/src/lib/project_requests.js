import { POST } from "$lib";


export async function addUserToProject(user_id, project_id) {
    await POST({user_id}, '/project/'+project_id+'/user', false)
}

export async function makeMilestoneInProject(project_id, name, deadline) {
    await POST({name, deadline}, '/project/'+project_id+'/milestone', false)
}
export async function makeTaskInMilestone(project_id, milestone_id, name) {
    console.log('/project/'+project_id+'/milestone/'+milestone_id+"/task")
    await POST({name}, '/project/'+project_id+'/milestone/'+milestone_id+"/task", false)
}