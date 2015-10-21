## Goto v1
### Easily navigate between your projects
The concept is that any nickname given by the user is registered against the current working folder. So later when the use types the command :

```shell
goto <nickname>
```
it goes to that project folder directly.

## Installation
  First run,
  ```shell
  make deps
  ```
  to get all the dependencies.
  Then run :
  ```shell
  make
  ```
  to install it to your system.
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
