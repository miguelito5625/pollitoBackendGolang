#!/bin/bash
echo ""
if [ -z "$1" ]
  then
    echo "Tienes que agregar el mensaje para el commit"
    exit 1
fi
git add -A
git commit -am "$1"
git push -u origin main
echo "SUBIDA A GITHUB FINALIZADA"