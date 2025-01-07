import asyncio
import uvicorn
from threading import Thread
from dotenv import load_dotenv
import os
load_dotenv()
from NotificationService import NotificationService
from RestCommander import RestCommander

notificationService = NotificationService()
rest_commander = RestCommander(notificationService)

def run_ws():
    print(f"WS on {os.getenv('LISTEN_PORT_WS')}")
    uvicorn.run(notificationService.fast_api_app, host="0.0.0.0", port=int(os.getenv('LISTEN_PORT_WS')))

def run_rest():
    print(f"REST on {os.getenv('LISTEN_PORT_REST')}")
    uvicorn.run(rest_commander.fast_api_app, host="0.0.0.0", port=int(os.getenv('LISTEN_PORT_REST')))

if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    rest_thread = Thread(target=run_rest)
    ws_thread = Thread(target=run_ws)

    rest_thread.start()
    ws_thread.start()

    rest_thread.join()
    ws_thread.join()
