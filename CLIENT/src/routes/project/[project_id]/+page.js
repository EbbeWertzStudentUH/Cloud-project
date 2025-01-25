import { GET } from "$lib";
import { updateOpenProject, updateProjectsList } from "../../../stores/projects";

let project_id = null;

export async function load({ params }) {
    project_id = params.project_id;
    return {fetchProject};
};

async function fetchProject() {
    const resp = await GET('/project/'+project_id);
    updateOpenProject(resp);
}