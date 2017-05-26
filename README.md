Origionally from the repo:

https://github.com/jhchen/ansize/blob/master/ansize.go


added special characters and a bash terminal script to randomly choose a *.asci file


You can begin by inserting the following code into your bash profile so that any script in the profile.d directory will be executed. You can add scripts later which you might want to automatically start when a bash terminal window is opened:
```bash
for sh in /etc/profile.d/*.sh; do
	 if [ -r "$sh" ] ; then
		 . "$sh"
	 fi
 done
 unset sh
```

 next we make a scipt file which can be called anything and we will place it in the "/etc/profile.d/" directory
