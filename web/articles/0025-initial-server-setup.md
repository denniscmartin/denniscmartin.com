TITLE=Initial server set up
DESCRIPTION=Initial server set up
DATE=29/08/2023
+++
Basic steps to do after creating an Ubuntu VM.

Login as root:

```bash
ssh root@$USERNAME
```

Create new user:

```bash
adduser $USERNAME
```

Granting administrative privileges:

```bash
usermod -aG sudo $USERNAME
```

Set basic firewall:

```bash
ufw allow OpenSSH
ufw enable
ufw status
```

Enable external access to our user (assuming the root account uses SSH key authentication):

```bash
rsync --archive --chown=$USERNAME:$USERNAME ~/.ssh /home/$USERNAME
```