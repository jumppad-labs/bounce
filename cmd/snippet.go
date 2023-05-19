package cmd

var zshSnippet = `
HISTFILE=/run/jumppad/history
HISTSIZE=500000
SAVEHIST=500000
setopt hist_ignore_dups share_history inc_append_history extended_history

promptcmd() {
	
}

precmd_functions=(promptcmd)
`

var bashSnippet = `
HISTFILE=/run/jumppad/history
HISTCONTROL=ignoredups:erasedups
HISTSIZE=500000
HISTFILESIZE=500000
shopt -s histappend

PROMPT_COMMAND+="history -a; history -c; history -r; $PROMPT_COMMAND"
`
