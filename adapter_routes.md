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
#### POST
> add wederzijds nieuwe vriend en krijg ge-update vriendenlijst
- ```HEADER{Authorisation:Bearer token}, JSON {friend_id}```
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
#### POST
> add nieuwe vriend en krijg ge-update friend requests lijst
- ```HEADER{Authorisation:Bearer token}, JSON {friend_id}```
- ```-> JSON [{first_name, last_name, id}]```
---
#### DELETE
> verwijder vriend en krijg ge-update friend requests lijst
- ```HEADER{Authorisation:Bearer token}, JSON {friend_id}```
- ```-> JSON [{first_name, last_name, id}]```
