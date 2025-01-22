# USER SERVICE
### /user
#### GET
> get user name
 - ``` QUERY PARAM {userID}```
 - ```-> JSON {first_name, last_name, id}```


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
### /notifier/subscribe
#### PUT /notifier/subscribe/friends
> subscribe de notifier connectie naar de friends list
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON (leeg/message)```
#### PUT /notifier/subscribe/projects
> subscribe de notifier connectie naar de projects list
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON (leeg/message)```
#### PUT /notifier/subscribe/all
> subscribe de notifier connectie naar de friends list en projects list
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON (leeg/message)```
#### PUT /notifier/subscribe/project
> subscribe de notifier connectie naar een specifiek project (en unsubscribe een enventueel vorig geopend project)
- ```HEADER{Authorisation:Bearer token}, JSON {subscribe_project, unsubscribe_project?}```
- ```-> JSON (leeg/message)```




# PROJECT SERVICE
### /project
#### POST /project
> maak nieuw project en laat notifier de projects list updaten
- ```JSON {name, deadline, github_repo}```
- ```-> JSON (leeg/message)```

#### GET /project/{project_id}
> get volledig project (incl milestones, tasks, problems)
- ```(enkel path segments)```
- ```-> JSON {id, name, users, deadline, github_repo, milestones}```

#### GET /projects
> get alle projecten waar jij deel van bent (minimale prject data, gewoon voor projects list)
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON [{id, name, users, deadline, github_repo, milestones}]```

#### POST /project/{project_id}/user
> Add a user to a project
- ```JSON {user_id}```
- ```-> JSON (leeg/message)```

#### /project/{project_id}/milestone
> maak milestone in project
- ```JSON {name, deadline}```
- ```-> JSON (leeg/message)```

#### POST /project/{project_id}/milestone/{milestone_id}/task
> maak task in milestone
- ```JSON {name}```
- ```-> JSON (leeg/message)```

#### POST /project/{project_id}/task/{task_id}/problem
> maak problem in task
- ```JSON {problem: {id, name, posted_at}}```
- ```-> JSON (leeg/message)```

#### PUT /project/{project_id}/task/{task_id}/problem/{problem_id}/resolve
> Resolve problem in task
- ```JSON {problem_id}```
- ```-> JSON (leeg/message)```

#### PUT /project/{project_id}/task/{task_id}/assign
> Assign een task naar jezelf
- ```HEADER{Authorisation:Bearer token}```
- ```-> JSON (leeg/message)```

#### PUT /project/{project_id}/task/{task_id}/complete
> set task naar complete
- ```(enkel path segments)```
- ```-> JSON (leeg/message)```
