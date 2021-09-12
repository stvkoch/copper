# WIP - Copper

Copper is a hound dog who is brought home by Amos Slade, a hunter. While still in puppy stage, he befriends a fox named Tod, and they become the very best of friends. Reality strikes, though, when he is sent off on a hunting trip with Chief and their master, and he is taught how to be a hunting dog.

Upon arriving home, he reminds Tod of his new status, and Tod seems hurt by the change in their friendship. When Chief sees Tod on their property, Copper saves him from getting killed by Amos Slade. However, when his escape results in Chief almost being killed, he claims revenge on Tod.

After hunting desperately for Tod, Copper is faced against a bear to protect his master, but gets hurt in the process. He finds it hard to stand up to the bear, but Tod comes and saves him. Therefore, in turn, when the bear is gone and Amos Slade tries to take advantage of Tod's injuries, Copper stands in the way to save Tod. Their friendship is repaired.

# The server


## Domain

We try organize the resources using "Package by feature" approach, delegating to the feature the responsability of what they should and how to do. They are the king of your own domain.

But, to avoid that "dependency graph of packages must have no cycles" we have special kind of citizen where they have free pass by the domains.

- models
- repositories

All ways dress your requests and responses right wear, it's a matter of etiquette defining the struct to requests and reponses.

Application layer separete the logic for handle a specific route or a group of routes, also is charge to config the routes that will be handle using the Routes function.

#### Registe handler

First you should record your handler into Router/registry file and define into handler the group route and routes you will catch and perform.

#### Defining the routes

Use Routes function inside of handler.

### Routes/registry.go

As sad, this file will hold all handlers actives into the Cooper WebServer

### Middleware/registry.go

Record the middleware used by Cooper Webserver


## HotReload dev server

https://github.com/codegangsta/gin#basic-usage

```
gin --appPort 4000 --all -i run server.go
```

> “5 days of coding can save 15 minutes of planning”


### Usin Docker-composer

```
cd copper
docker-compose up
```

### Reach api server

```
curl -X GET \
  'http://localhost:3000/api/v1/customer/123' \
  -H 'Accept: */*' \
  -H 'Content-Type: application/json'
```
