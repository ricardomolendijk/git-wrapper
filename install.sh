#!/bin/bash
printf "Installing git-wrapper...\n"
mkdir -p ~/bin
go build -o git-wrapper *.go

printf "Moving git-wrapper to ~/bin...\n"
mv git-wrapper ~/bin/git-wrapper
printf "Adding ~/bin to PATH...\n"
if ! grep -q 'export PATH="$HOME/bin:$PATH"' ~/.bashrc; then
    echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
    printf "Added ~/bin to PATH in ~/.bashrc\n"
else
    printf "~/bin is already in PATH in ~/.bashrc\n"
fi
printf "Sourcing ~/.bashrc to update PATH...\n"
source ~/.bashrc
printf "Installation complete!.\n"