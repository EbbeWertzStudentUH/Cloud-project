import { writable } from 'svelte/store';

export const projects_list = writable([]);
export const open_project = writable(null);

export function addUserToOpenProject(user) {
	open_project.update((project) => {
		return {
			...project,
			users: [...project.users, user]
		};
	});
}
export function addMilestoneToOpenProject(milestone) {
	open_project.update((project) => {
		return {
			...project,
			milestones: [...project.milestones, milestone]
		};
	});
}
export function updateTask(task_id, changed_fields) {
	open_project.update((project) => {
		return {
			...project,
			milestones: project.milestones.map((milestone) => {
				
				let updatedTasks = milestone.tasks.map((task) =>
					task.id === task_id
						? { ...task, user: { ...task.user, ...changed_fields.user }, ...changed_fields }
						: task
				);
				let finished_tasks = updatedTasks.filter((task) => task.status == 'closed').length;
				return {
					...milestone,
					tasks: updatedTasks,
					num_of_finished_tasks: finished_tasks
				};
			})
		};
	});
}
export function addTaskToMilestone(milestone_id, task) {
	open_project.update((project) => {
		return {
			...project,
			milestones: project.milestones.map((milestone) => {
				if (milestone.id === milestone_id) {
					return { ...milestone, tasks: [...milestone.tasks, task] };
				}
				return milestone;
			})
		};
	});
	open_project.update((project) => {
		return {
			...project,
			milestones: project.milestones.map((milestone) =>
				milestone.id === milestone_id
					? { ...milestone, num_of_tasks: milestone.tasks.length }
					: milestone
			)
		};
	});
}
export function updateOpenProject(project) {
	open_project.set(project);
}

export function updateProjectsList(newList) {
	projects_list.set(newList.map((item) => ({ id: item.id, name: item.name })));
}

export function addProject(project) {
	projects_list.update((items) => [...items, project]);
}
