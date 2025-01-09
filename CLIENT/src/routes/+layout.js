import { writable } from 'svelte/store';
import { goto } from '$app/navigation';

export async function load() {
    return {};
}

// export async function fetchUser() {
//   const token = localStorage.getItem('authToken');

//   if (token) {
//     try {
//       const res = await fetch('http://localhost:3001/user/authenticate', {
//         method: 'GET',
//         headers: { 'Authorization': `Bearer ${token}` },
//       });

//       if (res.ok) {
//         const data = await res.json();
//         if (data.valid) {
//           user.set({
//             firstName: data.first_name,
//             lastName: data.last_name,
//           });
//           return;
//         }
//       }
//     } catch (err) {
//       console.error('Failed to fetch user:', err);
//     }
//   }

//   user.set(null);
//   localStorage.removeItem('authToken');
// }

// export async function login(email, password) {
//   try {
//     const res = await fetch('http://localhost:3001/user/authenticate', {
//       method: 'POST',
//       headers: { 'Content-Type': 'application/json' },
//       body: JSON.stringify({ email, password }),
//     });

//     if (res.ok) {
//       const data = await res.json();
//       if (data.valid) {
//         localStorage.setItem('authToken', data.token);
//         await fetchUser();
//         goto('/');
//       } else {
//         throw new Error('Invalid email or password.');
//       }
//     } else {
//       throw new Error('Invalid email or password.');
//     }
//   } catch (err) {
//     console.error(err);
//     throw err;
//   }
// }

// export function logout() {
//   localStorage.removeItem('authToken');
//   user.set(null);
//   goto('/login');
// }

// onMount(() => {
    //   const token = localStorage.getItem('authToken');
    //   if (!token) {
    //     goto('/login');
    //   } else {
    //     // Fetch user details or do other authenticated actions
    //     user = { name: 'John Doe' }; // Replace with actual user data
    //   }
    // });