_gob(){
    wordf=`gob listnick`
    local cur=${COMP_WORDS[COMP_CWORD]}
    local prev=${COMP_WORDS[COMP_CWORD-1]}
    local prevev=${COMP_WORDS[COMP_CWORD-2]}
    case ${COMP_CWORD} in
	1)
	    COMPREPLY=( $( compgen -W "list listnick reg go branch" -- $cur))
	    ;;
	2)
	    case ${prev} in
		go)
		    COMPREPLY=( $( compgen -W "$wordf" -- $cur));;
		branch)
		    COMPREPLY=( $( compgen -W "$wordf" -- $cur));;
   	    esac
	    ;;
	3)
	    case ${prevev} in
		reg)
		    COMPREPLY=( $( compgen -o plusdirs  -f -X '!*.txt' -- $cur ));;
	    esac
	    ;;
    esac

    return 0
}

complete -F _gob -o filenames gob
