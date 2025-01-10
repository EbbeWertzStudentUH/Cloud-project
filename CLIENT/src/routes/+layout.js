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
        console.log("token is not present, going to login page");
        goto('/login');
    } else {
        console.log(`validating token ${token}...`);
        let data = await GETwithToken('user/authenticate');
        if (data.valid){
            console.log("token valid!");
            user.set(data.user);
        } else {
            console.log("token did not validate, going to login page");
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