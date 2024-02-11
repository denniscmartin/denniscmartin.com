+++
title = 'My template for C projects'
date = 2023-08-29
draft = true
+++

This is the script I use to create a basic C project fast. It creates the following files:

- A `CMakeLists.txt` file with a configuration.
- MIT `LICENSE.md` and `README.md` with the project name as a title.
- A `.gitignore` to ignore any folder called `build`.
- A simple `hello world` C file.
- A `.clang-format` file with the LINUX C style.

If you would like to use it, create a file with the
following code (I named it `c-init`) and place it in you `/bin` folder. Make sure the
folder is added to your `PATH`.

```bash
#!/bin/bash

if [ -z "$1" ]; then
        echo "Project name is missing."
        exit 1
fi

project_name=$1
project_dir=$(pwd)/$project_name

mkdir $project_dir
cd $project_dir

touch CMakeLists.txt
touch LICENSE.md
touch README.md
touch .gitignore

curl -LJO https://gist.githubusercontent.com/denniscmartin/\
67ca8777425c3a3ae0831089d36c557b/raw/517d5ff8a058780bb7a3403d7781d3a0a5fe8391/\
.clang-format

mkdir build
mkdir src
touch src/main.c

cat <<EOF >CMakeLists.txt
cmake_minimum_required(VERSION 3.1)

project($project_name
        VERSION 1.0
        LANGUAGES C
)

set(CMAKE_C_STANDARD 17)

add_custom_target(run
        COMMAND $project_name
        DEPENDS $project_name
        WORKING_DIRECTORY \${CMAKE_PROJECT_DIR}
)

add_executable($project_name src/main.c)
EOF

cat <<EOF >.gitignore
**/build/
EOF

cat <<EOF >src/main.c
#include <stdio.h>

int main() {
        printf("Hello, world!\n");

        return 0;
}
EOF

echo "Project created succesfully"
tree
```

## Usage

```bash
your-filename $PROJECT_NAME
```

In my case is:

```bash
c-init the-next-million-dollar-idea
```
