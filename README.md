# vault_dragon_test

## Local Setup

| Requirement | Version  |
| ----------  | -------- |
| Ansible     | 2.9.12 + |
| Virtual Box | 6.1.12 + |
| Vagrant     | 2.2.10 + |
| Golang      | 1.14 +   |

You'll need to do a couple things on your machine first to get it setup for development.

1. Update `/etc/hosts`:
    Add the following line to your `/etc/hosts` file:

    ```shell
    127.0.0.1 local.vaultdragon.com
    ```

2. Download repo and put under `$GOPATH/src` (ex: I set `$GOPATH` to `~/go`)
3. Set environment variable name `VD_ENV`
   
   ```shell
   export VD_ENV=local
   ```

4. Run vagrant at the **root level of the repository**:

    ```shell
    vagrant up
    ```

    This will start an `Ubuntu 20.04` instance in your virtual box and set up all the requirements.
    If this is your first time running the command, it can take up to 5 - 10 minutes.

5. Now go to `http://local.vaultdragon.com:8181`, you should see the site is up and running, and happy coding.

6. To develop in the code, you can modify the code on you local, and run 

    ```
    make compile
    ```

    this will build Go app binary and sync to `/usr/local/bin` in you virtual box, and then `ssh` to your ubuntu virtual box, restart supervisord

    ```
    vagrant ssh
    sudo service supervisord restart
    ```

## Development Workflow

| Important logs                 | Description           |
| -------------------------------| --------------------- |
| **`/var/log/vaultdragon/vaultdragon.log`** | Application log       |
| **`/var/log/vaultdragon/error.log`**   | Application error log |


### Frequently  Used Commands
 - **`sudo service supervisord <status/restart>`**
 - **`log_vaultdragon`**
 - **`log_error_vaultdragon`**