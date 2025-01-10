import { GETwithToken } from '$lib';
import { writable } from 'svelte/store';
import { goto } from '$app/navigation';


const user = writable(null);

export async function load() {
    return {user, fetchUser, logout};
}

async function fetchUser() {
    const token = localStorage.getItem('authToken');
    if (!token) {
        goto('/login');
    } else {
        let data = await GETwithToken('user/authenticate');
        if (data.valid){
            user.set(data);
        } else {
            localStorage.removeItem('authToken');
            goto('/login');
        }
    }
}

function logout() {
  localStorage.removeItem('authToken');
  user.set(null);
  goto('/login');
}