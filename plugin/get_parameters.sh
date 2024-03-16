#!/bin/sh

cue cmd dump | yq -p yaml -o json 'select(document_index == 0)' | jq '[{
  name: "Cue-parameters",
  title: "Cue Parameters",
  collectionType: "map",
  map: [leaf_paths as $path | {"key": $path | join("."), "value": getpath($path)|tostring}] | from_entries
}]'
