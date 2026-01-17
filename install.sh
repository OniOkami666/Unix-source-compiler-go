#! /bin/bash

# Reset
NC='\033[0m'


BLACK='\033[0;30m'
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
WHITE='\033[0;37m'


BOLD_GREEN='\033[1;32m'
BOLD_RED='\033[1;31m'
BOLD_CYAN='\033[1;36m'


set -e 

echo -e "${CYAN}Updating dependencies...${NC}"
sudo apt update

echo -e "${GREEN}Downloading the packages needed${NC}"
sudo apt install perl simh
