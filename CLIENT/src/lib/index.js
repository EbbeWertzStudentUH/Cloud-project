import { goto } from '$app/navigation';

export async function GETwithToken(path) {
    const token = await getToken();
	try {
		const res = await fetch('http://localhost:3001/' + path, {
			method: 'GET',
			headers: { Authorization: `Bearer ${token}` }
		});

		if (res.ok) {
			const data = await res.json();
			return data;
		} else {
            console.error('fetch GET with token gave error response: ', res.status, res.body);
        }
	} catch (err) {
		console.error('Failed to fetch GET with token:', err);
	}
}

export async function POST(body, path) {
      try {
        const res = await fetch('http://localhost:3001/'+path, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(body),
        });
    
        if (res.ok) {
          const data = await res.json();
          return data;
        } else {
            console.error('fetch POST gave error response: ', res.status, res.body);
        }
      } catch (err) {
        console.error('Failed to fetch POST:', err);
      }
    }

export async function getToken() {
	{
		const token = localStorage.getItem('authToken');
		if (!token) {
            console.log("token is not present, going to login page");
			goto('/login');
		} else {
			return token;
		}
	}
}

