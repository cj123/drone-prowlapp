# drone-prowlapp
a drone.io notify plugin for [ProwlApp](https://www.prowlapp.com).

### setup

follow drone's [official documentation on custom plugins](http://readme.drone.io/devs/plugins/#custom-plugins:dce8ed91d073f65a191aa58c2338afcb) to whitelist the seejy/drone-prowlapp plugin

file: `.drone.yml`
```yml
notify:
  prowlapp:
    apikey: some-prowlapp-apikey
    providerkey: some-prowlapp-providerkey 
```

## developing

### building

see `Makefile`

### debugging

```
$ ./drone-prowlapp <<EOF
{
  "repo": {
    "clone_url": "git://github.com/drone/drone",
    "owner": "drone",
    "name": "drone",
    "full_name": "drone/drone"
  },
  "system": {
    "link_url": "https://beta.drone.io"
  },
  "build": {
    "number": 22,
    "status": "success",
    "started_at": 1421029603,
    "finished_at": 1421029813,
    "message": "Update the Readme",
    "author": "johnsmith",
    "author_email": "john.smith@gmail.com",
    "event": "push",
    "branch": "master",
    "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
    "ref": "refs/heads/master"
  },
  "workspace": {
    "root": "/drone/src",
    "path": "/drone/src/github.com/drone/drone"
  },
  "vargs": {
    "provider_key": "PROWLAPP_PROVIDER_KEY",
    "apikey": "PROWLAPP_APIKEY"
  }
}
EOF
```
