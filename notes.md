## Mongo

```
# To have launchd start mongodb at login:
ln -sfv /usr/local/opt/mongodb/*.plist ~/Library/LaunchAgents
# Then to load mongodb now:
launchctl load ~/Library/LaunchAgents/homebrew.mxcl.mongodb.plist
#O r, if you don't want/need launchctl, you can just run:
mongod --config /usr/local/etc/mongod.conf
```


## Go

* Bookmarks in /golang

* Use URL routing from http://www.gorillatoolkit.org/pkg/mux


## Site Structure

* / Main page

	- Log in / Create new user
	- default handler

* /tasks 

	- show all user's tasks
	- redirect home if not authenticated

* /tasks/new
	
	- create new task

* /tasks/{someTask}

	- view task with this ID
