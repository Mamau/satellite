## Satellite of your project

#### This is just a wrapper for docker commands and nothing more
This utility just transform .yaml command presentation for docker command.  
This:
```yaml
docker:
  run:
    - name: "my-image"
      detach: true
      clean-up: true
      image: "gitlab.com/my/image"
      environment-variables:
        - "PHP_IDE_CONFIG=serverName=192.168.0.1"
      volumes:
        - "$(pwd):/home/www"
      dns:
        - "8.8.8.8"
      ports:
        - "127.0.0.1:443:443"
        - "80:80"
```
Will be:
```bash
docker run -d --rm -p 127.0.0.1:443:443 -p 80:80 --dns=8.8.8.8 -v /path/to/current/dir:/home/www --name my-image gitlab.com/my/image
```
And you don't need to remember your big commands. Git will remember it.

Allows you not to install on your computer programs that
accompanying developments.  
That service for you:  
* If your project uses many services.   
* If your team has people who do not need to know the entire ecosystem of the project.  
* If you are tired of keeping in mind all commands with the help of which different parts of the project are assembled.  

You can describe all possible command variations in the yaml file and run them, having only docker on your machine. For convenience, the yaml config can be added to the revision of the git, and the bin file can be added to gitignore, because different operating systems will have their own launchers.  

For php you can use [composer package wkit](https://github.com/Mamau/satellite-cli). For downloading last version.   
If you don't use composer, you can just download [bin file](https://github.com/Mamau/satellite/releases), and put in the root of the project.  
For unix system need access:
```bash
chmod +x sat
```
For MacOS need additional rights, because i do not have a developer certificate.

Example:
```bash
sat help
``` 
### The following docker commands are currently supported:
* #### [run](https://docs.docker.com/engine/reference/commandline/run/)
* #### [exec](https://docs.docker.com/engine/reference/commandline/exec/)
* #### [pull](https://docs.docker.com/engine/reference/commandline/pull/)
Commands above allow you to run docker command of the same name (e.g. docker run anyImage) 
* #### [docker-compose](https://docs.docker.com/compose/reference/)
Docker-compose allow you to run command using docker-compose file (e.g. docker-compose exec... | docker-compose run...)  
[Examples](https://github.com/Mamau/satellite/tree/master/example)

### Environment variables
Satellite support .env file. All variables you can define there.
Then use it in .yaml file like this:
```yaml
docker:
  run:
    - name: "my-command"
      user: "${USER_ID_FROM_ENV_FILE}"     
```
In .env file you can define CONFIG_NAME variable.
This variable set name of config which will be used by satellite.
If it not set, satellite will search "satellite.yaml"  
Satellite support use follow commands:  
* **$(pwd)** - will get your current dir path  
* **$(gatewayHost)** - will get **Gateway** host from command: "_docker network inspect bridge_"  
If you define a network property - satellite will create it automatically
### Possible services
* #### Docker pull image
*Use it for pull image*
```yaml
docker:
  pull:
    - name: "fresh-img"
      image: "node"
      version: "10"
      description: "pull node 10"
```
Config above allow you run:
```bash
./sat fresh-img
```

* #### Docker exec command through container
*Use it for execute command*
```yaml
docker:
  exec:
    - name: "console"
      interactive: true
      tty: true
      bin-bash: true
      user: "${USER_ID}" #this will get from .env file
      container-name: "my_project"
      workdir: "/home/www/my_project_dir"
      beginning: "php bin/console"
      description: "execute console command of container"
```
Config above allow you run:
```bash
./sat console cache-clear
```
Exec command have a feature field "beginning". This is first part of your command.  
When you will run command above in fact you will run this:
```bash
docker exec -it --user 1000 --workdir=/home/www/my_project_dir my_project php bin/console cache-clear
```

* #### Docker start image
*Use it for run your image as detached service*
```yaml
docker:
  run:
    - name: "my-image"
      detach: true
      clean-up: true
      image: "gitlab.com/my/image"
      environment-variables:
        - "PHP_IDE_CONFIG=serverName=192.168.0.1"
      volumes:
        - "$(pwd):/home/www"
      dns:
        - "8.8.8.8"
      ports:
        - "127.0.0.1:443:443"
        - "80:80"
```
Config above allow you run:
```bash
./sat my-image
```

* #### Docker run image with commands 
*Use it for install dependencies*
```yaml
docker:
  run:
    - name: "my-composer"
      image: "composer"
      clean-up: true
      user: "${USER_ID}" #this will get from .env file
      interactive: true
      workdir: "/home/www-data"
      tty: true
      beginning: "composer"
      version: "1.9"
      volumes:
        - "$(pwd):/home/www-data"
```
Config above allow you run:
```bash
./sat my-composer install --ignore-platform-reqs
```

Run command have a feature field "beginning". This is first part of your command.  
When you will run command above in fact you will run this:
```bash
docker run -it --user 1000 --workdir=/home/www-data composer:1.9 composer install --ignore-platform-reqs
```

* #### Start docker-compose
*Use it for run docker-compose files*
```yaml
docker-compose:
  up:
    - name: "run-docker-compose"
      path: "./path/to/file/docker-compose"
      project-directory: "./path/to/file/docker-compose"
      project-name: "anyName"
      description: "run docker compose data"
      verbose: true
      log-level: "DEBUG"
```
Config above allow you run:
```bash
./sat docker-compose-name run
```

* #### Uniting command from service section
*Use it for run 2 command*
```yaml
macros:
  - name: "init"
    commands:
      - "composer i --ignore-platform-reqs"
      - "yarn install"
```
Config above allow you run:
```bash
./sat macros init
```


### Example of config
*[You can see here](https://github.com/Mamau/satellite/tree/master/example)*  

