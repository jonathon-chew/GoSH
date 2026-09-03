# CHANGELOG

	## v0.1.0 

		### NEW
		1. check for file nad no longer errors out and breaks if can't write to the file
		1. quick helper script to find new builtins/common functions that can be added to the git and are not just templates
		1. adding previous attempts at recreating bash functionality and the underlying shell commands and options
		1. basic implimentation
		### UPDATES
		1. import path to the projects better name
		1. remove logging into printline statements and fatals as a shell this should never fail and break out of executable
		1. go-sh to GoSH as a better name for the project
		1. functions and other parts to allow for the refactor of logic into their own files
		### REFACTORS
		1. everything into their own files and packages where relevent, refactor input into it's own function as this only takes in a single line, this will need to parse multiple lines going forward...
		1. moving key logic to their own files