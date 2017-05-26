#!/bin/bash
list=($(find go/ansize/ -name '*.ansi'))
rand=$[ $RANDOM % ${#list[@]} ]
cat ${list[$rand]}
