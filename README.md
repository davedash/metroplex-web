# Metroplex: A Developer Friendly Deployment System

This purpose of this system is to create an easy to use deployment system.

- It is designed for large organizations but should be more than suitable for
  individuals and hobbyists alike.
- It should be easy enough that anybody engineer can use it to deploy their own
  software and feel comfortable knowing that it went out correctly.
- It will be designed with lessons learned at real companies [1][1].

[1]: http://engineering.pinterest.com/post/70621633046/deploying-software-at-pinterest

The full system will be called Metroplex.  
*Metroplex puts things in their place.*

It'll consist of a deployment service called `metroplex-web`.  It will be backed
by mysql and should be horizontally scalable.  Think of it as a simple mysql
backed web-app.

SigmaD will be a daemon each machine.  Its configuration will tell it which
Metroplex API Host to contact
(otherwise we'll just assume a host named 'vector') and send health status.

You can define services from the web frontend and register various services to
nodes.

You may also self-register on the node using SigmaD via command line options.

Once you have services and nodes defined, you can then kick off deploys via the
web console.  A deploy will trigger each host that's associated with a service
to kick off any pre-download scripts, download a build, extract a build, run any
pre-restart scritps, restart a service, run any post-restart scripts and then
signal to vector that the deploy is up to date.

## The technology stack

This system intends to use:

- Go (golang) for the service and daemon
- MySQL (as a backend)
- Bootstrap for a CSS framework
- jQuery for JS

Outside of "Go", I'm familiar enough with the other tools that I'll be able to
maintain some velocity for developing this tool.  I will entertain other
technology as our needs change, but strive for simplicity.

We'll write both the service and the daemon in Go.  Go is a good fit for the
daemon which is annoying if it crashes.  Go is an ok fit for the web service,
but other frameworks might suffice (Python's Bottle, Ruby Sinatra, NodeJS).  At
least there's no context switch.

We'll back this initially with MySQL since it's easy to operate and is present
in many environments.

\* Title is totally a working project title.
