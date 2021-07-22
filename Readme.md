## Satellite of your project

Allows you not to install on your computer programs that
accompanying developments.  
That service for you:  
* If your project uses many services.   
* If your team has people who do not need to know the entire ecosystem of the project.  
* If you are tired of keeping in mind all commands with the help of which different parts of the project are assembled.  

You can describe all possible command variations in the yaml file and run them, having only docker on your machine. For convenience, the yaml config can be added to the revision of the git, and the bin file can be added to gitignore, because different operating systems will have their own launchers.  

For php you can use [composer package wkit](https://github.com/Mamau/satellite-cli). For downloading last version.   
If you dont use composer, you can just download [bin file](https://github.com/Mamau/satellite/releases), and put in the root of the project.  
For unix system need access:
```bash
chmod +x sat
```
For MacOS need additional rights, because i do not have a developer certificate.

Example:
```bash
sat help
``` 

### Possible services
* #### Docker pull image
*Use it for pull image*
```yaml
services:
  - name: "fresh-img"
    command: "pull"
    image: "node"
    version: "10"
```
Config above allow you run:
```bash
./sat fresh-img
```

* #### Docker exec command through container
*Use it for execute command*
```yaml
services:
  - name: "composer"
    description: "Install composer dependencies"
    version: "2"
    user-id: "$(uid)"
    image-command: "composer"
    work-dir: "/home/www-data"
    volumes:
      - "$(pwd):/home/www-data"
      - "$(pwd)/cache:/tmp"
```
Config above allow you run:
```bash
./sat composer install --ignore-platofrm-reqs
```

* #### Docker start image
*Use it for run your image as detached service*
```yaml
services:
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

* #### Start docker-compose
*Use it for run docker-compose files*
```yaml
services:
  - name: "docker-compose-name"
    docker-compose: true #required
    path: "./path/to/file/docker-compose"
    project-directory: "./path/to/file/docker-compose"
    project-name: "anyName"
    verbose: true
    log-level: "DEBUG"
    detach: true
    command: "up"
    remove-orphans: true
    
  - name: "docker-compose-name-2"
    command: "up"
    docker-compose: true # required
    detach: true

  - name: "docker-compose-down"
    command: "down"
    docker-compose: true # required
```
Config above allow you run:
```bash
./sat docker-compose-name
# or
./sat docker-compose-name-2
# or
./sat docker-compose-down
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

