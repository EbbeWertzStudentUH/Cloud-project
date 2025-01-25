import { writable } from "svelte/store";

export const projects_list = writable([]);
export const open_project = writable(null);

export function updateOpenProject(project){
    open_project.set(project);
}

export function updateProjectsList(newList){
    projects_list.set(newList.map(item => ({id: item.id,name: item.name})));
}

export function addProject(project){
    projects_list.update(items => [...items, project]);
}