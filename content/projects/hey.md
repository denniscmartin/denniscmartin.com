+++
title = 'Hey! Network programming in C'
date = 2023-07-21
thumbnail = 'projects-hey-thumbnail.png'
draft = false
tags = ['networking', 'c-language', 'programming', 'docker']
+++


Hey! is a simple TCP streaming app developed in C using POSIX APIs. Both client and 
server are containerized using Docker.

You can find the code in my [Github](https://github.com/denniscmartin/hey)

I also recorded a [Youtube video](https://www.youtube.com/watch?v=r3CQ0euv6TQ)

## Usage

Make sure you have Docker installed.

```bash
./run.sh
```

Using docker compose, this script build the Docker images (server and client), create a 
user-defined bridge network, and run both containers. After that, you will be provided 
with the client shell to send messages to the server.

To close the connection type `exit` from the client shell and press `ENTER`.

Some things that I find interested:

- Run docker `logs --follow hey-server-1` in another terminal. Doing that you can see the messages arriving at the server.
- If you have Wireshark installed you can use it to sniff the packets sent in the Docker network. This is good to understand TCP.

# What could you do next?

A good exercise will be to implement logic for some commands in the server. For example, 
you could program the server to send to the client the current date every time the server 
receives the command `date`.