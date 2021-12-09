#!/bin/bash
git add -A
git commit -am "$1"
git push -u origin main
echo "SUBIDA A GITHUB FINALIZADA"