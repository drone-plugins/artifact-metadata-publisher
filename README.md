# Publishes url to artifact tab via plugin step

Format to output the `PLUGIN_ARTIFACT_FILE` :
```
{
  "kind": "fileUpload/v1",
  "data": {
    "fileArtifacts": [
      {
        "name": "<File name>",
        "url": "<URL of the file>"
      }
    ]
  }
}
```


Inputs to set as part of plugin settings:

1. artifact_file: Location of file to store data output
