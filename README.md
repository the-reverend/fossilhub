fossilhub
=========

use nginx and go to fill in some gaps in the fossil ui

## prerequisites

* install fossil and nginx

		sudo apt-get install fossil nginx

* install go; follow the instructions [here](http://golang.org/doc/install)

## setup your fossil repository

fossil repositories are fully contained in a single file.  i strongly recommend you adhere to the naming convention of a .fossil extension, e.g. myproject.fossil, because fossil server expects this.  unlike other dvcs systems, e.g. git, fossil does not create a dot-folder of metadata in your source path.  my personal convention is to create a folder e.g. ~/fossil/ to keep all my fossil repositories and create project folders e.g. ~/src/myproject/ in which to keep open checkouts.  you can have as many fossil repositories as you have space.  it makes it simple to serve them all if they are all stored in the same folder, e.g. ~/fossil/.

	$ mkdir ~/fossil
	$ fossil new ~/fossil/myproject.fossil
	$ mkdir ~/src
	$ mkdir ~/src/myproject
	$ cd ~/src/myproject
	$ fossil open ~/fossil/myproject.fossil

## xinetd

xinetd will execute the fossil server to service each request.  there is a sample xinetd configuration file in this project.  modify it to your taste and copy it or link it to /etc/xinetd.d/.  in the sample, only local traffic will be served.  also, add the following line to /etc/services :

	# Local services
	fossil       8080/tcp      # fossil scm

## fossilhub

fossil will serve each of your repositories as subfolders, e.g. localhost:8080/yourproject/.  unfortunately, fossil will not serve a root page that lists all your repositories, i.e. localhost:8080/ will be a 404.  this project seeks to fill that gap.  currently, fossilhub serves a page that shows all the repositories stored in ~/fossil/*.  more features that can obviously flow from this include: add/remove repository, permissions management, backups, social interraction, etc.

	todo : describe build process and fossilhub configuration

## nginx

i use nginx to proxy the fossil server and fossilhub under a subdomain so that i don't have to remember port numbers.  there is a sample nginx configuration file in this project.  modify it to your taste and copy it or link it to /etc/nginx/sites-enabled/.  this will glue fossil server and fossilhub to your subdomain such that e.g. fossil.yourdomain.com serves the hub homepage with nice links to each of your repositories, e.g. fossil.yourdomain.com/yourproject.
