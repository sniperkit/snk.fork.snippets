# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.10

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/local/Cellar/cmake/3.10.2/bin/cmake

# The command to remove a file.
RM = /usr/local/Cellar/cmake/3.10.2/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug

# Utility rule file for lint_src_scanner_c.

# Include the progress variables for this target.
include CMakeFiles/lint_src_scanner_c.dir/progress.make

CMakeFiles/lint_src_scanner_c:
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --blue --bold --progress-dir=/Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Linting scanner.c"
	/usr/local/Cellar/cmake/3.10.2/bin/cmake -E chdir /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7 /usr/bin/perl /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/scripts/checkpatch.pl --no-tree --terse --strict --showfile --summary-file --max-line-length=120 --ignore=COMPARISON_TO_NULL --ignore=GLOBAL_INITIALISERS --ignore=INITIALISED_STATIC --ignore=PREFER_KERNEL_TYPES --ignore=COMPLEX_MACRO --ignore=CAMELCASE --ignore=NEW_TYPEDEFS --ignore=MULTISTATEMENT_MACRO_USE_DO_WHILE --ignore=TRAILING_SEMICOLON --ignore=AVOID_EXTERNS -f src/scanner.c

lint_src_scanner_c: CMakeFiles/lint_src_scanner_c
lint_src_scanner_c: CMakeFiles/lint_src_scanner_c.dir/build.make

.PHONY : lint_src_scanner_c

# Rule to build all files generated by this target.
CMakeFiles/lint_src_scanner_c.dir/build: lint_src_scanner_c

.PHONY : CMakeFiles/lint_src_scanner_c.dir/build

CMakeFiles/lint_src_scanner_c.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/lint_src_scanner_c.dir/cmake_clean.cmake
.PHONY : CMakeFiles/lint_src_scanner_c.dir/clean

CMakeFiles/lint_src_scanner_c.dir/depend:
	cd /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7 /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7 /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug/CMakeFiles/lint_src_scanner_c.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/lint_src_scanner_c.dir/depend

