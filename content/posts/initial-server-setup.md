+++
title = 'Some configurations to start a new server'
date = 2023-08-29
draft = false
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
