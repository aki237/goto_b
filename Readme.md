## Goto v1
### Easily navigate between your projects
The concept is that any nickname given by the user is registered against the current working folder. So later when the use types the command :

```shell
goto <nickname>
```
it goes to that project folder directly.

## Installation
  The zip package shows the folders present in the home directory. So copy the ```goto_b``` file in the bin folder to any of the folders in the ```$PATH``` variable. Now, obviously, without bash completion this is just a waste. So place the .config/goto/ folder in the home directory. As the location of the database is hard coded, it is essential to keep the ```goto.db``` in that given folder (if not present create one). Now place the ```*.bash and goto``` files in the same folder or in some other folder of choice. Let the folder of choice be ```FOLDER/```. Now in ```.bashrc``` add the following lines.
```shell
source $FOLDER/goto_completion.bash
source $FOLDER/goto
source $FOLDER/gos_completion.bash
  ```
  Probably you are good to go.

## Usage

First goto any project folder. Now record this folder path against any random nickname.
```shell
goto_b reg myrandnick
```
or
```shell
goto_b reg myrandnick [PROJECT_FOLDER_PATH]
```
Now this folder nick has been recorded in the database. Now to navigate to that project folder,
```shell
goto [tab]
```
You'll get possible completions.

## Why goto??
This goto project is written, because for me as a go developer switching between projects is just a pain always.
