import { goto } from '$app/navigation';
import { addNotification, notifications } from '../stores/notifications';
import { addMessage } from '../stores/updatemessages';

let ws = null;

export async function GETwithToken(path) {
    const token = getToken();
    if(!token) return;
	try {
		const res = await fetch('http://localhost:3001' + path, {
			method: 'GET',
			headers: { Authorization: `Bearer ${token}` }
		});

		if (res.ok) {
			const data = await res.json();
      console.log(" ===== GET RESULT ===== ");
      console.log(data);
			return data;
		} else {
            console.error('fetch GET with token gave error response: ', res.status, res.body);
        }
	} catch (err) {
		console.error('Failed to fetch GET with token:', err);
	}
}
export async function PUTwithTokenNoResult(path) {
  const token = getToken();
  if(!token) return;
try {
  const res = await fetch('http://localhost:3001' + path, {
    method: 'PUT',
    headers: { Authorization: `Bearer ${token}` }
  });

  if (!res.ok) {
    console.error('fetch GET with token gave error response: ', res.status, res.body);
  }
} catch (err) {
  console.error('Failed to fetch GET with token:', err);
}
}
export async function GET(path) {
try {
  const res = await fetch('http://localhost:3001'+path, {
    method: 'GET',
  });
  if (res.ok) {
    const data = await res.json();
    console.log(" ===== GET RESULT ===== ");
    console.log(data);
    return data;
  } else {
          console.error('fetch GET gave error response: ', res.status, res.body);
      }
} catch (err) {
  console.error('Failed to fetch GET:', err);
}
}
export async function DELETEwithToken(body, path) {
  const token = getToken();
  if(!token) return;
try {
  const res = await fetch('http://localhost:3001' + path, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`
     },
    body: JSON.stringify(body),
  });

  if (res.ok) {
    const data = await res.json();
    console.log(" ===== DELETE RESULT ===== ");
      console.log(data);
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
        const res = await fetch('http://localhost:3001'+path, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(body),
        });
    
        if (res.ok) {
          const data = await res.json();
          console.log(" ===== POST RESULT ===== ");
          console.log(data);
          return data;
        } else {
            console.error('fetch POST gave error response: ', res.status, res.body);
        }
      } catch (err) {
        console.error('Failed to fetch POST:', err);
      }
}
export async function POSTWithToken(body, path) {
      const token = getToken();
      if(!token) return;
      try {
        const res = await fetch('http://localhost:3001'+path, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`
           },
          body: JSON.stringify(body),
        });
    
        if (res.ok) {
          const data = await res.json();
          console.log(" ===== POST RESULT ===== ");
          console.log(data);
          return data;
        } else {
            console.error('fetch POST gave error response: ', res.status, res.body);
        }
      } catch (err) {
        console.error('Failed to fetch POST:', err);
      }
}
export function getToken() {
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

export function killWebSocket(){
  ws.close();
}

export function initializeWebSocket() {
  ws = new WebSocket("ws://localhost:3004/ws");

  ws.onopen = () => {
      console.log("WebSocket connected");
      const token = getToken();
      ws.send(token);
  };

  ws.onmessage = (event) => {
      const resp = JSON.parse(event.data);
      console.log("WebSocket message received:", resp);
      if(resp.type == "notification"){
        addNotification(resp.data);
      }
      addMessage(resp);
  };

  ws.onclose = () => {
      console.log("WebSocket connection closed");
  };

  ws.onerror = (error) => {
      console.error("WebSocket error:", error);
  };
}