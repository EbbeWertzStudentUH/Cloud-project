import { GET, PUTwithTokenNoResult } from "$lib";
import { updateOpenProject, open_project } from "../../../stores/projects";
import {get} from 'svelte/store'


let project_id = null;

export async function load({ params }) {
    project_id = params.project_id;
    return {fetchProject, getGithubStats};
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

async function getGithubStats(gh_url){
    try {
		const res = await fetch('http://localhost:3010/github-stats', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({githubUrl: gh_url})
		});
		if (res.ok) {
			return await res.json();
		} else {
			console.error('fetch POST to github stats gave error response: ', res.status, res.body);
		}
	} catch (err) {
		console.error('Failed to fetch POST to github stats:', err);
	}
}