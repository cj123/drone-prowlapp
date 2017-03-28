# drone-prowlapp
a [drone.io](http://readme.drone.io/) notify plugin for [ProwlApp](https://www.prowlapp.com).

### setup

follow drone's [official documentation on custom plugins](http://readme.drone.io/plugins/plugin-overview/) to whitelist the seejy/drone-prowlapp plugin

file: `.drone.yml`
```yml
pipeline:
  prowlapp:
    image: seejy/drone-prowlapp
    apikey: some-prowlapp-apikey
    provider_key: some-prowlapp-providerkey 
```

## developing

### building

see `Makefile`