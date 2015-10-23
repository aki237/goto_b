#!/bin/bash
cnt=${1000}

_goto(){
    unset cnt
    wordf=`goto_b listnick`
    i=0
    for word in $w
    do
	cnt[i]=${word}
	i=$[$i+1]
    done
    local cur=${COMP_WORDS[COMP_CWORD]}
    case "$cur" in
	*)
	    COMPREPLY=( $( compgen -W "$wordf" -- $cur))
    esac

    return 0
}

complete -F _goto -o filenames goto
