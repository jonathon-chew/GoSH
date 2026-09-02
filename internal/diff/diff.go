package diff

import (
	"fmt"
	"os"
	"strings"
)

func version() {
	fmt.Print("v0.0.1")
}

func ignoreNewLines(file1, file2 string) {}

func ignoreSpaces(file1, file2 string) {}

func ignoreCase(file1, file2 string) {
	file1Bytes, ErrReadingFile1 := os.ReadFile(file1)
	if ErrReadingFile1 != nil {
		return
	}

	file2Bytes, ErrReadingFile2 := os.ReadFile(file2)
	if ErrReadingFile2 != nil {
		return
	}

	file1Split, file2Split := strings.Split(string(file1Bytes), "\n"), strings.Split(string(file2Bytes), "\n")

	file1Map, file2Map := make(map[string]int), make(map[string]int)

	for _, line := range file1Split {
		file1Map[line]++
	}

	for _, line := range file2Split {
		file2Map[line]++
	}

	for k1, v1 := range file1Map {
		for k2, v2 := range file2Map {
			if k1 == k2 && v1 == v2 {
				delete(file1Map, k1)
				delete(file2Map, k2)
			}
		}
	}

}

func Diff() {

	var arguments []string
	if len(os.Args) > 1 {
		arguments = os.Args[1:]
	} else {
		return
	}

	for argumentNumber, argument := range arguments {
		switch argument {
		case "-B":
			if len(os.Args[argumentNumber+1:]) > 2 {
				file1 := os.Args[argumentNumber+1]
				file2 := os.Args[argumentNumber+2]
				ignoreNewLines(file1, file2)
			}
		case "-b":
			if len(os.Args[argumentNumber+1:]) > 2 {
				file1 := os.Args[argumentNumber+1]
				file2 := os.Args[argumentNumber+2]
				ignoreSpaces(file1, file2)
			}
		case "-i":
			if len(os.Args[argumentNumber+1:]) > 2 {
				file1 := os.Args[argumentNumber+1]
				file2 := os.Args[argumentNumber+2]
				ignoreCase(file1, file2)
			}
		case "-v":
			version()
		}
	}

	/*
		 -C ______ --context ______
		        Like -c but produces a diff with ______ lines of context.

		-c      Produces a diff with 3 lines of context.  With -c the output
		        format is modified slightly: the output begins with
		        identification of the files involved and their creation dates and
		        then each change is separated by a line with fifteen *'s.  The
		        lines removed from _____ are marked with ‘- ’; those added to
		        _____ are marked ‘+ ’.  Lines which are changed from one file to
		        the other are marked in both files with ‘! ’.  Changes which lie
		        within 3 lines of each other are grouped together on output.

		-D ______ --ifdef ______
		        Creates a merged version of _____ and _____ on the standard
		        output, with C preprocessor controls included so that a
		        compilation of the result without defining ______ is equivalent
		        to compiling _____, while defining ______ will yield _____.

		-e --ed
		        Produces output in a form suitable as input for the editor
		        utility, ed(1), which can then be used to convert file1 into
		        file2.

		        Note that when comparing directories with -e, the resulting file
		        may no longer be interpreted as an ed(1) script.  Output is added
		        to indicate which file each set of ed(1) commands applies to.
		        These hunks can be manually extracted to produce an ed(1) script,
		        which can also be applied with patch(1).

		-f --forward-ed
		        Identical output to that of the -e flag, but in reverse order.
		        It cannot be digested by ed(1).

		--help  This option prints a summary to stdout and exits with status 0.

		-n      Produces a script similar to that of -e, but in the opposite
		        order and with a count of changed lines on each insert or delete
		        command.  This is the form used by rcsdiff.

		-q --brief
		        Just print a line when the files differ.  Does not output a list
		        of changes.

		-U ______ --unified ______
		        Like -u but produces a diff with ______ lines of context.

		-u      Produces a _______ diff with 3 lines of context.  A unified diff
		        is similar to the context diff produced by the -c option.
		        However, unlike with -c, all lines to be changed (added and/or
		        removed) are present in a single section.

		--version
		        This option prints a version string to stdout and exits with
		        status 0.

		-y --side-by-side
		        Output in two columns with a marker between them.  The marker can
		        be one of the following:

		              space   Corresponding lines are identical.
		              '|'     Corresponding lines are different.
		              '<'     Files differ and only the first file contains the
		                      line.
		              '>'     Files differ and only the second file contains the
		                      line.

		Comparison options:

		-A _____ --algorithm ____
		        Configure the algorithm used when comparing files.  diff supports
		        3 algorithms:

		        myers   The Myers diff algorithm finds the shortest edit which
		                transforms one input into the other.  It generally runs
		                in O(N+D²) time, requiring O(N) space, where N is the sum
		                of the lengths of the inputs and D is the length of the
		                difference between them, with a theoretical O(N·D) worst
		                case.  If it encounters worst-case input, the
		                implementation used by diff falls back to a less optimal
		                but faster algorithm.
		        patience
		                The Patience variant of the Myers algorithm attempts to
		                create more aesthetically pleasing diff output by
		                logically grouping lines.
		        stone   The Stone algorithm (commonly known as Hunt-McIlroy or
		                Hunt-Szymanski) looks for the longest common subsequence
		                between compared files.  Stone encounters worst case
		                performance when there are long common subsequences.  In
		                large files this can lead to a significant performance
		                impact.  The Stone algorithm is maintained for
		                compatibility.

		        The default algorithm is affected by the POSIXLY_CORRECT and
		        POSIX_PEDANTIC environment variables.  When either variable is
		        set the default algorithm will be “stone”.  Otherwise, the diff
		        utility defaults to the Myers algorithm, but will fall back to
		        the Stone algorithm if the input or output options are not
		        supported by the Myers implementation.

		-a --text
		        Treat all files as ASCII text.  Normally diff will simply print
		        “Binary files ... differ” if files contain binary characters.
		        Use of this option forces diff to produce a diff.

		-B --ignore-blank-lines
		        Causes chunks that include only blank lines to be ignored.

		-b --ignore-space-change
		        Causes trailing blanks (spaces and tabs) to be ignored, and other
		        strings of blanks to compare equal.

		--color=[____]
		        Color the additions green, and removals red, or the value in the
		        DIFFCOLORS environment variable.  The possible values of ____ are
		        “never”, “always” and “auto”.  auto will use color if the output
		        is a tty and the COLORTERM environment variable is set to a non-
		        empty string.

		-d --minimal
		        Try very hard to produce a diff as small as possible.  This may
		        consume a lot of processing power and memory when processing
		        large files with many changes.

		-F ________ --show-function-line _______
		        Like -p, but display the last line that matches provided pattern.

		-I _______ --ignore-matching-lines _______
		        Ignores changes, insertions, and deletions whose lines match the
		        extended regular expression _______.  Multiple -I patterns may be
		        specified.  All lines in the change must match some pattern for
		        the change to be ignored.  See re_format(7) for more information
		        on regular expression patterns.

		-i --ignore-case
		        Ignores the case of letters.  E.g., “A” will compare equal to
		        “a”.

		-l --paginate
		        Pass the output through pr(1) to paginate it.

		-L _____ --label _____
		        Print _____ instead of the first (and second, if this option is
		        specified twice) file name and time in the context or unified
		        diff header.

		-p --show-c-function
		        With unified and context diffs, show with each change the first
		        40 characters of the last line before the context beginning with
		        a letter, an underscore or a dollar sign.  For C and Objective-C
		        source code following standard layout conventions, this will show
		        the prototype of the function the change applies to.

		-T --initial-tab
		        Print a tab rather than a space before the rest of the line for
		        the normal, context or unified output formats.  This makes the
		        alignment of tabs in the line consistent.

		-t --expand-tabs
		        Will expand tabs in output lines.  Normal or -c output adds
		        character(s) to the front of each line which may screw up the
		        indentation of the original source lines and make the output
		        listing difficult to interpret.  This option will preserve the
		        original source's indentation.

		-w --ignore-all-blanks
		        Is similar to -b --ignore-space-change but causes whitespace
		        (blanks and tabs) to be totally ignored.  E.g., “if ( a == b )”
		        will compare equal to “if(a==b)”.

		-W ______ --width ______
		        Output at most ______ columns when using side by side format.
		        The default value is 130.  Note that unless

		-t      was specified, diff will always align the second column to a tab
		        stop, so values of --width smaller than approximately five times
		        the value of --tabsize may yield surprising results.

		--changed-group-format ____
		        Format input groups in the provided

		        the format is a string with special keywords:

		        %<  lines from FILE1

		        %<  lines from FILE2

		--ignore-file-name-case
		        ignore case when comparing file names

		--no-ignore-file-name-case
		        do not ignore case when comparing file names (default)

		--normal
		        default diff output

		--speed-large-files
		        stub option for compatibility with GNU diff

		--strip-trailing-cr
		        strip carriage return on input files

		--suppress-common-lines
		        Do not output common lines when using the side by side format

		--tabsize ______
		        Number of spaces representing a tab (default 8)

		Directory comparison options:

		-N --new-file
		        If a file is found in only one directory, act as if it was found
		        in the other directory too but was of zero size.

		-P --unidirectional-new-file
		        If a file is found only in ____, act as if it was found in ____
		        too but was of zero size.

		-r --recursive
		        Causes application of diff recursively to common subdirectories
		        encountered.

		-S ____ --starting-file ____
		        Re-starts a directory diff in the middle, beginning with file
		        ____.

		-s --report-identical-files
		        Causes diff to report files which are the same, which are
		        otherwise not mentioned.

		-X ____ --exclude-from ____
		        Exclude files and subdirectories from comparison whose basenames
		        match lines in ____.  Multiple -X options may be specified.

		-x _______ --exclude _______
		        Exclude files and subdirectories from comparison whose basenames
		        match _______.  Patterns are matched using shell-style globbing
		        via fnmatch(3).  Multiple -x options may be specified.
	*/
}
