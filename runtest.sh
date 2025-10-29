#!/bin/bash
set -e

BASE_FOLDERS_URL="http://localhost:8080/api/folders"
BASE_NOTES_URL="http://localhost:8080/api/notes"

echo "=== Creando carpeta padre ==="
parent_id=$(curl -s -X POST "$BASE_FOLDERS_URL" \
  -H "Content-Type: application/json" \
  -d '{"name":"Carpeta Padre","description":"Carpeta principal"}' \
  | grep -o '"ID"[ ]*:[ ]*[0-9]*' | sed 's/[^0-9]*//g')
echo "Carpeta Padre creada con ID: $parent_id"
echo ""

echo "=== Creando subcarpeta 1 ==="
sub1_id=$(curl -s -X POST "$BASE_FOLDERS_URL" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Subcarpeta 1\",\"description\":\"Primera subcarpeta\",\"parent_folder_id\":$parent_id}" \
  | grep -o '"ID"[ ]*:[ ]*[0-9]*' | sed 's/[^0-9]*//g')
echo "Subcarpeta 1 creada con ID: $sub1_id"
echo ""

echo "=== Creando subcarpeta 2 ==="
sub2_id=$(curl -s -X POST "$BASE_FOLDERS_URL" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Subcarpeta 2\",\"description\":\"Segunda subcarpeta\",\"parent_folder_id\":$parent_id}" \
  | grep -o '"ID"[ ]*:[ ]*[0-9]*' | sed 's/[^0-9]*//g')
echo "Subcarpeta 2 creada con ID: $sub2_id"
echo ""

echo "=== Creando nota en Carpeta Padre ==="
note1_id=$(curl -s -X POST "$BASE_NOTES_URL" \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Nota Padre\",\"body\":\"Contenido de la nota principal\",\"folder_id\":$parent_id}" \
  | grep -o '"ID"[ ]*:[ ]*[0-9]*' | sed 's/[^0-9]*//g')
echo "Nota creada con ID: $note1_id"
echo ""

echo "=== Creando nota en Subcarpeta 1 ==="
note2_id=$(curl -s -X POST "$BASE_NOTES_URL" \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Nota Subcarpeta 1\",\"body\":\"Contenido de la subcarpeta 1\",\"folder_id\":$sub1_id}" \
  | grep -o '"ID"[ ]*:[ ]*[0-9]*' | sed 's/[^0-9]*//g')
echo "Nota creada con ID: $note2_id"
echo ""

for id in $parent_id $sub1_id $sub2_id; do
  echo "=== Obteniendo carpeta con ID $id ==="
  curl -s -X GET "$BASE_FOLDERS_URL/$id"
  echo -e "\n"
done

for id in $note1_id $note2_id; do
  echo "=== Obteniendo nota con ID $id ==="
  curl -s -X GET "$BASE_NOTES_URL/$id"
  echo -e "\n"
done

echo "=== Actualizando Subcarpeta 2 para que sea hija de Subcarpeta 1 ==="
curl -s -X PUT "$BASE_FOLDERS_URL/$sub2_id" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Subcarpeta 2\",\"description\":\"Ahora es hija de Subcarpeta 1\",\"parent_folder_id\":$sub1_id}"
echo -e "\n"

echo "=== Actualizando nota de Subcarpeta 1 ==="
curl -s -X PUT "$BASE_NOTES_URL/$note2_id" \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Nota Subcarpeta 1 Actualizada\",\"body\":\"Contenido actualizado\"}"
echo -e "\n"

echo "=== Listando todas las carpetas ==="
curl -s -X GET "$BASE_FOLDERS_URL"
echo -e "\n"

echo "=== Listando todas las notas ==="
curl -s -X GET "$BASE_NOTES_URL"
echo -e "\n"

echo "=== Eliminando Subcarpeta 1 (ID $sub1_id) ==="
curl -s -X DELETE "$BASE_FOLDERS_URL/$sub1_id"
echo -e "\n"

echo "=== Eliminando nota de Subcarpeta 1 (ID $note2_id) ==="
curl -s -X DELETE "$BASE_NOTES_URL/$note2_id"
echo -e "\n"

echo "=== Intentando obtener Subcarpeta 1 eliminada ==="
curl -s -X GET "$BASE_FOLDERS_URL/$sub1_id"
echo -e "\n"

echo "=== Intentando obtener nota eliminada ==="
curl -s -X GET "$BASE_NOTES_URL/$note2_id"
echo -e "\n"

echo "=== Listando todas las carpetas finales ==="
curl -s -X GET "$BASE_FOLDERS_URL"
echo -e "\n"

echo "=== Listando todas las notas finales ==="
curl -s -X GET "$BASE_NOTES_URL"
echo -e "\n"

