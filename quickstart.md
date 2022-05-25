## quick start

quick start for onepw

### Commands

* help     - `display help`
* version  - `display version`
* init     - `init password box or modify master password`
* add      - `add a new password or update old password`
* import   - `import notes.xlsx to onepw`
* remove   - `remove passwords by ids or (category,account)(aliases rm,del,delete)`
* list     - `list all passwords(aliases ls)`
* find     - `find password by id,category,account,tag,site and so on`
* upgrade  - `upgrade to newest version(aliases up)`
* generate - `a utility command for generating password(aliases gen)`
* info     - `show low level information of password`


### help - `show help information`

```sh
# show help information of onepw
$> onepw
# or
$> onepw help

# show help information of specific command
$> onepw help COMMAND
# or
$> onepw COMMAND -h
```

### version - `show onepw version`

```sh
$> onepw version
# or
$> onepw -v
```

### init - `init password box`
First of all, you should `init` a password box with master password.

```sh
# Will prompt for enter master password
$> onepw init
type the master password:
```

**NOTE**: The master password can be set by ENV variable ONEPW_MASTER.

### add - `add a new command or update old password`

```sh
Options:

  -h, --help
      display help information

  --master[=$ONEPW_MASTER]
      master password

  --debug[=false]
      usage debug mode

  -c, --category
      category of password

  -u, --account
      account of password

  --site
      website of password

  --tag
      tags of password

  --id
      password id for updating

  --pw, --password
      the password

  --cpw, --confirm-password
      confirm password
```

```sh
$> onepw add -c=email -u user@example.com
type the password:
repeat the password:
```

### import - `import notes.xlsx to onepw password data`

```sh
Options:

  -h, --help                     Display help information
      --master[=$ONEPW_MASTER]   Your master password
      --debug[=false]            Enable debug mode
  -n, --notes[=notes.xlsx]       need import notes path
```

```sh
$> onepw import -n notes.xlsx
```

### list - `list all passwords, aliases ls`

```sh
Options:

  -h, --help
      display help information

  --master[=$ONEPW_MASTER]
      master password

  --debug[=false]
      usage debug mode

  --no-header[=false]
      don't print header line
```

```sh
$> onepw list
# or
$> onepw ls
```

### remove - `remove passwords by ids or account, aliases rm/del/delete`

```sh
Usage: onepw rm [ids...] [OPTIONS]

Options:

  -h, --help
      display help information

  --master[=$ONEPW_MASTER]
      master password

  --debug[=false]
      usage debug mode

  -a, --all[=false]
      remove all found passwords
```

### find - `find passwords by id,category,account,...`

```sh
Usage: onepw find <WORD>

Options:

  -h, --help
      display help information

  --master[=$ONEPW_MASTER]
      master password

  --debug[=false]
      usage debug mode

  -p, --just-password[=false]
      only show password

  -f, --just-first[=false]
      only show first result
```

### generate - `generate password, aliases gen`

```sh
Usage: onepw gen [OPTIONS] LEN

Options:

  -h, --help
      display help information

  -n, --number=N[=1]
      number of generated passwords

  -d, --digit[=false]
      whether the password contains digit

  -c, --lower-char[=false]
      whether the password contains lowercase character

  -C, --upper-char[=false]
      whether the password contains uppercase character

  -s, --special-char[=false]
      whether the password contains the special character

  --sset, --special-set
      custom special character set
```

```sh
$> onepw gen 12
FA7vAeZML02r
$> onepw gen 12 -cs
iqva%kj*^!!f
$> onepw gen 16 -cCdS
0g1b^TgAUXAij2KC
```

### info - `show low-level information of password`

```sh
Usage: onepw info <ids...>

Options:

  -h, --help
      display help information

  --master[=$ONEPW_MASTER]
      master password

  --debug[=false]
      usage debug mode

  -a, --all
      show all found passwords
```

## Example

```sh
$> echo "export ONEPW_FILE=~/mypasswords/password.data"
$> echo "export ONEPW_MASTER=MySecret"

# init password box
$> onepw init

# add a new password
$> onepw add -c email -u user@example.com
type the password: 
repeat the password: 
password d9437f07af7c8b035a4fa9513ace449f added

# list all passwords
$> onepw ls
+---------+----------+------------------+----------+---------------------------+
| ID      | CATEGORY | ACCOUNT          | PASSWORD | UPDATED_AT                |
+---------+----------+------------------+----------+---------------------------+
| d9437f0 | email    | user@example.com | 123456   | 2016-04-29T00:54:36+08:00 |
+---------+----------+------------------+----------+---------------------------+

# add a new password
$> onepw add -c github -u hello --pw=123456 --cpw=123456
password 3439d3178f35f56f4c3d6f27e7ccc9a7 added

# list all passwords
$> onepw ls
+---------+----------+------------------+----------+---------------------------+
| ID      | CATEGORY | ACCOUNT          | PASSWORD | UPDATED_AT                |
+---------+----------+------------------+----------+---------------------------+
| 3439d31 | github   | hello            | 123456   | 2016-04-29T00:56:26+08:00 |
+---------+----------+------------------+----------+---------------------------+
| d9437f0 | email    | user@example.com | 123456   | 2016-04-29T00:54:36+08:00 |
+---------+----------+------------------+----------+---------------------------+

# add a new password
$> onepw add -c email -u user2@gmail.com --site=gmail.com --tag=google
type the password:
repeat the password:
password 2ca000f993a665337bebd4700cfd7c6c added

# list all passwords
$> onepw ls
+---------+----------+------------------+----------+---------------------------+
| ID      | CATEGORY | ACCOUNT          | PASSWORD | UPDATED_AT                |
+---------+----------+------------------+----------+---------------------------+
| 2ca000f | email    | user2@gmail.com  | 123456   | 2016-04-29T00:58:49+08:00 |
+---------+----------+------------------+----------+---------------------------+
| 3439d31 | github   | hello            | 123456   | 2016-04-29T00:56:26+08:00 |
+---------+----------+------------------+----------+---------------------------+
| d9437f0 | email    | user@example.com | 123456   | 2016-04-29T00:54:36+08:00 |
+---------+----------+------------------+----------+---------------------------+

# find passwords
$> onepw find mail
+---------+-------+------------------+--------+---------------------------+
| 2ca000f | email | user2@gmail.com  | 123456 | 2016-04-29T00:58:49+08:00 |
+---------+-------+------------------+--------+---------------------------+
| d9437f0 | email | user@example.com | 123456 | 2016-04-29T00:54:36+08:00 |
+---------+-------+------------------+--------+---------------------------+

# find first password
$> onepw find mail -f
+---------+-------+------------------+--------+---------------------------+
| 2ca000f | email | user2@gmail.com  | 123456 | 2016-04-29T00:58:49+08:00 |
+---------+-------+------------------+--------+---------------------------+

# find passwords, but only show password
$> onepw find mail -p
123456
123456

# ^TRY:
# onepw find mail -pf

$> onepw find hello
+---------+--------+-------+--------+---------------------------+
| 3439d31 | github | hello | 123456 | 2016-04-29T00:56:26+08:00 |
+---------+--------+-------+--------+---------------------------+

# remove passwords
$> onepw rm 343
deleted passwords:
3439d3178f35f56f4c3d6f27e7ccc9a7

# list all passwords
$> onepw ls
+---------+----------+------------------+----------+---------------------------+
| ID      | CATEGORY | ACCOUNT          | PASSWORD | UPDATED_AT                |
+---------+----------+------------------+----------+---------------------------+
| 2ca000f | email    | user2@gmail.com  | 123456   | 2016-04-29T00:58:49+08:00 |
+---------+----------+------------------+----------+---------------------------+
| d9437f0 | email    | user@example.com | 123456   | 2016-04-29T00:54:36+08:00 |
+---------+----------+------------------+----------+---------------------------+
```
