#!/bin/bash

# Install instructions
# 1. make this script executable with `sudo chmod +x install.sh`
# 2. run script `./install.sh`
# ---- 

# copy executable binary
sudo cp runik /usr/local/bin && \

# copy desktop icon
sudo cp appicon.png /usr/share/icons/hicolor/256x256/apps/runik.png && \

# copy .desktop entry
sudo cp runik.desktop /usr/share/applications && \

# update icon cache
sudo update-desktop-database ~/.local/share/applications && \

echo "Runik successfully installed."
