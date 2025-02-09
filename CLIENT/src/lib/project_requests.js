import { POST, PUTwithTokenNoResult } from "$lib";


export async function addUserToProject(user_id, project_id) {
    await POST({user_id}, '/project/'+project_id+'/user', false)
}

export async function makeMilestoneInProject(project_id, name, deadline) {
    await POST({name, deadline}, '/project/'+project_id+'/milestone', false)
}
export async function makeTaskInMilestone(project_id, milestone_id, name) {
    await POST({name}, '/project/'+project_id+'/milestone/'+milestone_id+"/task", false)
}
export async function assignToTask(project_id, task_id) {
    await PUTwithTokenNoResult('/project/'+project_id+'/task/'+task_id+"/assign")
}
export async function completeTask(project_id, task_id) {
    await PUTwithTokenNoResult('/project/'+project_id+'/task/'+task_id+"/complete")
}
export async function makeProblemInTask(project_id, task_id, problem_name) {
    await POST({problem:{name:problem_name, posted_at:''}}, '/project/'+project_id+'/task/'+task_id+"/problem", false)
}
export async function ResolveProblemInTask(project_id, task_id, problem_id) {
    await PUTwithTokenNoResult('/project/'+project_id+'/task/'+task_id+'/problem/'+problem_id+"/resolve")
}