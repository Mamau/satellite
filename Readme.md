## Starter service

Позволяет не устанавливать на свой компьютер программы которые 
сопутствуют разработке, такие как: composer, bower, yarn.  
Так же позволяет описать в yaml файле в разделе **services** нужные образы и использовать их.  
Через раздел **macros** можно задать часто используемые команды и запускать все одной командой. 
```yaml
commands:
  composer:
    version: "2"
    user-id: "1000"
  bower:
    user-id: "1000"
  yarn:
    version: "10"
    user-id: "1000"

macros:
  - init:
    name: "init"
    commands:
      - "composer install --ignore-platform-reqs"
      - "yarn install"
      - "bower install"
  - build:
    name: "build-front"
    commands:
      - "yarn build-dev"
  - dump:
    name: "dump-and-show-php-version"
    commands:
      - "php -v"
      - "mysql mysqldump -u username -ppassword database_name  > db.sql"

services:
  - php:
    name: "php"
    version: "7.4"
  - mysql:
    name: "mysql"
    version: "5.7"

```
Исходя из этого конфига можно запускать одной командой группы действий:
```bash
./starter macros init # запустит установку composer, yarn и bower
./starter macros build-front # запустить сборку фронта
./starter macros dump-and-show-php-version # покажет версию php и сделает дамп mysql
```
Отличие блоков commands от services в том, что в commands заранее определены сервисы и у них можно описать конфиг, в блоке services можно описать только запуск докер команды с использованием pre и post-commands.  
На машине необходим только Docker.  
Для php есть [композер пакет wkit](https://github.com/Mamau/web-kit) через этот
пакет можно скачать последнюю версию   
Если по каким-то причинам нет возможности использовать композер 
Можно скачать нужный вам [бинарь тут](https://github.com/Mamau/starter/releases), и положите в корень проекта.  
Для юникс систем надо будет дать права на исполнения:
```bash
chmod +x starter
```
MacOS требует чтобы такие файлы были подписаны ключом разработчика apple, у меня
нет этого ключа и таким пользователям придется делать дополнительные действия для того
чтобы этот бинарь можно было запустить.

Пример
```bash
starter help
``` 
Покажет доступные команды.  
Если вы работаете с php проектом, вы наверняка используете composer,
но если вам не хочется его устанавливать локально, то Starter ваш выбор.
Положив в корень проекта бинарь запустите
```bash
starter composer i
``` 
Эта команда выполнит установку всех зависимостей и скинет vendor
в текущую директорию.  
Установить пакет:
```bash
starter composer require some/package
``` 
Для дополнительных настроек, в корне проекта можно создать starter.yaml, 
В секции pre-commands можно перечислить все команды, которые необходимо
выполнить перед основной, добавить репозиторий или настроить конфиг...  
Если требуется в команду передать параметры вида --param=value, то нужно их
передавать через "--", пример:  
```bash
starter yarn install -- --param=value --param2=value2
```
пример с передачей параметра в основную команду, и дочерних параметров дочереней команде
```bash
starter yarn install --version=12 -- --param=value --param2=value2
```
В данном примере --version=12 передается в основную команду, и будет взят
образ ноды с 12 версией и параметры --param=value --param2=value2 будут
переданы в yarn контейнер и выполнится команда: yarn install --param=value --param2=value2  
### Доступные атрибуты для команд в starter.yaml
*Многие атрибуты дублируют докер*  
  
Название | Значение | Описание
--- | --- | ---
`version` | строка | Версия образа, например композер должен быть 1.9 версии, значит там стоит поставить 1.9, перетирает флаг команды --version
`user-id` | строка | От какого пользователя будут создаваться (волюмироваться) файлы 
`work-dir` | строка | Рабочая директория скрипта 
`environment-variables` | список | Список env переменных со значениями (SOME_VAR=someVal), можно использовать в связке с post/pre-commands
`add-hosts` | список | Добавить хост в контейнер 
`ports` | список | Прокинуть нужные порты 
`volumes` | список | Прокинуть нужные волюмы, по умолчанию прокидывается текущая директория в домашнюю образа 
`dns` | список | Прокинуть нужные dns 
`post-commands` | список | Список команд которые должны выполнится перед основной 
`pre-commands` | список | Список команд которые должны исполнятся после основной 


### Примеры конфигураций
*[можно посмотреть тут](https://github.com/Mamau/starter/tree/master/example-config)*  

