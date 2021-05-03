#!/usr/bin/bash
which bash
# Traee el compilado de vue a el server de go lang
dirActual=$(pwd)
cd ../view/chat/
npm run build
cd $dirActual
cp -r ../view/chat/dist ./views # copiamos dist a view del proyecto en go
# cp -r dist/ ../../chatone1/views/
echo 'Directorio view actualizado con vue'
go build
echo 'Proyecto de go compilado'
