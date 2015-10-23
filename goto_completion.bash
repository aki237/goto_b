#!/bin/bash

_goto(){
    wordf=`gob listnick`
    local cur=${COMP_WORDS[COMP_CWORD]}
    local prev=${COMP_WORDS[COMP_CWORD-1]}
    case ${COMP_CWORD} in
	1)
	    COMPREPLY=( $( compgen -W "$wordf" -- $cur))
	    ;;
	2)
	    dirpresent=`gob go $prev`"/.git"
	    if [ -d "$dirpresent" ]
	    then
		dir=`gob branch $prev`
		COMPREPLY=( $( compgen -W "$dir" -- $cur))
	    else
		COMPREPLY=( $( compgen -W "" -- $cur))
	    fi
	    
    esac

    return 0
}

complete -F _goto -o filenames goto
