_gos(){
    wordf=`goto_b listnick`
    local cur=${COMP_WORDS[COMP_CWORD]}
    local prev=${COMP_WORDS[COMP_CWORD-1]}
    case ${COMP_CWORD} in
	1)
	    COMPREPLY=( $( compgen -W "list listnick reg go" -- $cur))
	    ;;
	2)
	    case ${prev} in
		go)
		    COMPREPLY=( $( compgen -W "$wordf" -- $cur));;
	    esac
	    ;;
    esac

    return 0
}

complete -F _gos -o filenames goto_b
