#!/bin/bash

# some ANSI colors for nicer logs
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # no color

echo -e "${YELLOW}🔍 Running go test...${NC}"
go test -v -cover ./...
if [ $? -ne 0 ]; then
    echo -e "${RED}🚨 Tests failed. Commit aborted.${NC}"
    exit 1
fi
echo -e "${GREEN}✅ Tests passed${NC}"

echo -e "${YELLOW}🔍 Running gosec...${NC}"
gosec -include=G101 ./...
if [ $? -ne 0 ]; then
    # TODO: including only needed rules and set severity to MEDIUM-HIGH
    echo -e "${RED}🚨 gosec found security issues. Commit aborted.${NC}"
    exit 1
fi
echo -e "${GREEN}✅ gosec passed${NC}"

echo -e "${YELLOW}🔍 Running gitleaks...${NC}"
gitleaks detect --no-git --verbose
if [ $? -ne 0 ]; then
    echo -e "${RED}🚨 gitleaks found secrets. Commit aborted.${NC}"
    exit 1
fi
echo -e "${GREEN}✅ gitleaks passed${NC}"

echo -e "${GREEN}🎉 All checks passed. Happy committing!${NC}"
