import asyncio
import uvicorn
from threading import Thread
from NotificationService import NotificationService
from RestCommander import RestCommander



notificationService = NotificationService()
rest_commander = RestCommander(notificationService)

# Function to run the REST API server (FastAPI) on port 5000
def run_ws():
    uvicorn.run(notificationService.fast_api_app, host="0.0.0.0", port=5000)

# Function to run the WebSocket server (FastAPI) on port 5001
def run_rest():
    uvicorn.run(rest_commander.fast_api_app, host="0.0.0.0", port=5001)

# Start both servers concurrently using asyncio and threading
if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    # Run both REST API and WebSocket servers concurrently
    rest_thread = Thread(target=run_rest)
    ws_thread = Thread(target=run_ws)

    rest_thread.start()
    ws_thread.start()

    rest_thread.join()
    ws_thread.join()
