## Satellite of your project

Allows you not to install on your computer programs that
accompanying developments.  
That service for you:  
* If your project uses many services.   
* If your team has people who do not need to know the entire ecosystem of the project.  
* If you are tired of keeping in mind all commands with the help of which different parts of the project are assembled.  

You can describe all possible command variations in the yaml file and run them, having only docker on your machine. For convenience, the yaml config can be added to the revision of the git, and the bin file can be added to gitignore, because different operating systems will have their own launchers.  

```yaml
services:
  - name: "yarn"
    image: "node"
    version: "14"
    user-id: "1000"
    image-command: "yarn"
    work-dir: "/home/node"
    volumes:
      - "$(pwd):/home/node"
      - "$(pwd)/cache:/tmp"

  - name: "composer"
    version: "2"
    user-id: "1000"
    image-command: "composer"
    work-dir: "/home/www-data"
    volumes:
      - "$(pwd):/home/www-data"
      - "$(pwd)/cache:/tmp"

macros:
  - name: "init"
    commands:
      - "composer i --ignore-platform-reqs"
      - "yarn install"

```
From config described above, you can run commands"
```bash
./stlt yarn install 
#in fact will run: docker run -ti -u 1000 --workdir=/home/node -v $(pwd):/home/node -v $(pwd)/cache:/tmp node:14 yarn install
./stlt composer i --ignore-platform-reqs 
#in fact will run: docker run -ti -u 1000 --workdir=/home/www-data -v $(pwd):/home/www-data -v $(pwd)/cache:/tmp composer:2 composer i --ignore-platform-reqs
./stlt macros init # will run all commands from section "commands"
```
For php you can use [composer package wkit](https://github.com/Mamau/web-kit). For downloading last version.   
If you dont use composer, you can just download [bin file](https://github.com/Mamau/satellite/releases), and put in the root of the project.  
For unix system need access:
```bash
chmod +x stlt
```
For MacOS need additional rights, because i do not have a developer certificate.

Example:
```bash
stlt help
``` 
### Example of config
*[You can see here](https://github.com/Mamau/satellite/tree/master/example)*  

