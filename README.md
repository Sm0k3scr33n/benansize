Originally from the repo:

https://github.com/jhchen/ansize/blob/master/ansize.go


added special characters and a bash terminal script to randomly choose a *.asci file


You can begin by inserting the following code into your bash profile so that any script in the "/etc/profile.d/" directory will be executed. You can add scripts later which you might want to automatically start when a bash terminal window is opened:
```bash
for sh in /etc/profile.d/*.sh; do
	 if [ -r "$sh" ] ; then
		 . "$sh"
	 fi
 done
 unset sh
```

 Next we make a script file which can be called anything and we will place it in the "/etc/profile.d/" directory (this assumes you place your ansi files in the same directory as the ansize program. You may consider placing the asci files in the same directory as your bash script):

 ```bash
 #!/bin/bash
 list=($(find go/ansize/ -name '*.ansi'))
 rand=$[ $RANDOM % ${#list[@]} ]
 cat ${list[$rand]}
 ```
