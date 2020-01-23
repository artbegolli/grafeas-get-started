# Grafeas Getting Started

This is a getting started guide for the setup and running of grafeas on a local kubernetes cluster. It will also go through the creation of notes, occurrences and attestation in preparation for integration with the ocibuilder.

The getting started guide will assume you are in this projects working directory.

## Deploying Grafeas

You can deploy grafeas to kubernetes using the `grafeas.yaml` manifest in this repository.

```
kubectl apply -f grafeas.yaml
```

Create a secure tunnel to your Grafeas server

```
kubectl port-forward \
  $(kubectl get pods -l app=grafeas -o jsonpath='{.items[0].metadata.name}') \
  8080:8080
```

## Pushing a Note to Grafeas

Notes define a high level piece of metadata *e.g.* storing information about the builder of a build process. Notes are often
created by providers doing the analysis. The note ID must be unique per project.

Note names must follow the format `/projects/<project_id>/notes/<note_id>`
 
Only edited by note owner - read only for users who have access to occurences referencing them.

You can post a note to grafeus using a standard post curl

```
curl -X POST \
  'http://127.0.0.1:8080/v1beta1/projects/image-signing/notes?noteId=production' \
  -d @./resources/note.json
```

You should then be able to view all your project notes here: http://127.0.0.1:8080/v1beta1/projects/image-signing/notes

## Pushing an Occurrence to Grafeas

An occurence is an instance of a Note - describing how and when a given note occurs on the resource associated with the occurrence.

*e.g.* an occurrence of a note about build details would describe the container images that resulted from a build

Occurrence names should follow the format `/projects/<project_id>/occurrences/<occurrence_id>`

The occurrence ID must be unique per project and is often random. Typically, occurrences are stored in separate projects than those where notes are created.

Write access to occurrences should only be granted to users who have access to link a note to the occurrence. Any user can have read access to occurrences.

You can post occurences to grafeas using a standard post curl

```
curl -X POST \
  'http://127.0.0.1:8080/v1beta1/projects/image-signing/occurrences' \
  -d @resources/occurrence.json
```

### Pushing an Occurrence using the [Grafeas Client API](https://github.com/grafeas/client-go/tree/master/0.1.0)

You can push images to Grafeas from go using the Grafeas Go client API.

An easy to import version of the client which isn't broken is available at <github.com/ocibuilder/gofeas>.

```
go get github.com/ocibuilder/gofeas
```
