# USER SERVICE
### /user/authenticate
>
#### POST
 > login
 - ``` JSON {email, password}```
 - ```-> JSON {valid, token, user {first_name, last_name, id}}```
---
#### GET
> auth via token
- ```HEADER{Authorisation:Bearer token}```
 - ```-> JSON {valid, token, user {first_name, last_name, id}}```

### /user/create_account
>
#### POST
> register + account setup
- ```JSON {email, password, first_name, last_name}```
- ```-> JSON {valid, token, user {first_name, last_name, id}}```

### /user/friends
>
#### GET
> get vriendenlijst
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON [{first_name, last_name, id}]```
---
#### DELETE
> verwijder wederzijds vriend en krijg ge-update vriendenlijst
- ```HEADER{Authorisation:Bearer token}, JSON {friend_id}```
- ```-> JSON [{first_name, last_name, id}]```
    
### /user/friend-requests
>
#### GET
> get friend requests lijst
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON [{first_name, last_name, id}]```
---
#### POST /user/friend-requests/send
> stuur vriendverzoek naar iemand
- ```HEADER{Authorisation:Bearer token}, JSON {friend_id}```
- ```-> JSON (leeg/message)```
---
#### POST /user/friend-requests/accept
> verwijder vriendverzoek en voeg elkaar als vriend toe, en krijg ge-update vrienden lijst
- ```HEADER{Authorisation:Bearer token}, JSON {friend_id}```
- ```-> JSON [{first_name, last_name, id}]```
---
#### DELETE /user/friend-requests/reject
> verwijder vriendverzoek en krijg ge-update friend requests lijst
- ```HEADER{Authorisation:Bearer token}, JSON {friend_id}```
- ```-> JSON [{first_name, last_name, id}]```

# NOTIFIIER SERVICE
### /notifier/subscribe-friends
#### PUT
> subscribe de notifier connectie naar de friends list
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON (leeg/message)```