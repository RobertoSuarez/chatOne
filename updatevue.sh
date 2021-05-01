#!/usr/bin/env bash
cd ../view/chat/
npm run build
echo 'Compilación terminada'
cp -r dist/ ../../chatone1/views/
echo 'Directori view actualizado con vue'
