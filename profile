for sh in /etc/profile.d/*.sh; do
	 if [ -r "$sh" ] ; then
		 . "$sh"
	 fi
 done
 unset sh
