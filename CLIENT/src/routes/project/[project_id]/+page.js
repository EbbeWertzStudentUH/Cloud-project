import { GET, POST, PUTwithTokenNoResult } from "$lib";
import { updateOpenProject, updateProjectsList } from "../../../stores/projects";
import { open_project } from '../../../stores/projects';
import {get} from 'svelte/store'



let project_id = null;

export async function load({ params }) {
    project_id = params.project_id;
    return {fetchProject, addUserToProject};
};

async function fetchProject() {
    const prevProj = get(open_project);
    let unsubscribe_project = null;
    if(prevProj){
        unsubscribe_project = prevProj.id;
    }
    await PUTwithTokenNoResult('/notifier/subscribe/project', {subscribe_project:project_id, unsubscribe_project})
    const resp = await GET('/project/'+project_id);
    updateOpenProject(resp);
}

async function addUserToProject(user_id) {
    await POST({user_id}, '/project/'+project_id+'/user', false)
}