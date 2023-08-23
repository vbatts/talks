---
author: Vincent Batts (vbatts)
---

# OCI images

and their background

---

## Timeline

* "specs" (2015-05-05) ([DockerCon 15 moment](https://youtu.be/at72dhg-SZY?t=6594))
* split to "runtime-spec" and "image-spec" (2016-03-23)
* created "distribution-spec" from [docker/docker#9015](https://github.com/docker/docker/issues/9015) (2018-03-08)

(Should've basically gone in reverse!)

* v1.0.0 for runtime and image spec on 2017-07-19
* v1.0.0 for distribution spec on 2021-04-26
* v1.1.0-rc.1 currently for runtime-spec
* v1.1.0-rc3 currently for distribution-spec
* v1.1.0-rc4 currently for image-spec

---

## Spec or source?

Open Containers Initiative really tried to _only_ be specifications.

But straight away there was [runc](https://github.com/opencontainers/runc)...

There are also `-tools` repos corresponding to specs (though these are poorly adopted).

Then there was [go-digest](https://github.com/opencontainers/go-digest), [umoci](https://github.com/opencontainers/umoci), and [selinux](https://github.com/opencontainers/selinux) as well.

---

## First, Useful tools

[crane](https://github.com/google/go-containerregistry/tree/main/cmd/crane)

Install like:
```bash
VERSION=$(curl -s "https://api.github.com/repos/google/go-containerregistry/releases/latest" | jq -r '.tag_name')
OS=Linux       # or Darwin, Windows
ARCH=x86_64    # or arm64, x86_64, armv6, i386, s390x
curl -sL "https://github.com/google/go-containerregistry/releases/download/${VERSION}/go-containerregistry_${OS}_${ARCH}.tar.gz" > go-containerregistry.tar.gz
tar -zxvf go-containerregistry.tar.gz -C /usr/local/bin/ crane
```

or just `go install github.com/google/go-containerregistry/cmd/crane@latest` 😄

also, [oras](https://github.com/oras-project/oras) is good to be familiar with if you're importing their libraries.

also also, [skopeo](https://github.com/containers/skopeo).

---

## Now, about the image

For the _most part_ it is a namechange of the corresponding docker data structures.

[image-spec/media-types.md](https://github.com/opencontainers/image-spec/blob/main/media-types.md) has more inforation.

The word "image" is effectively legacy that is here to stay.

---

# Relationship

```
~~~graph-easy --as=boxart
digraph G {
  {
    imageIndex [label="Image Index\n<<optional>>\napplication/vnd.oci.image.index.v1+json"]
    {
      rank=same
      manifest [label="Image manifest\napplication/vnd.oci.image.manifest.v1+json"]
    }
    config [label="Image config JSON\napplication/vnd.oci.image.config.v1+json"]
    layer [label="Layer tar archive\napplication/vnd.oci.image.layer.v1.tar\napplication/vnd.oci.image.layer.v1.tar+gzip\napplication/vnd.oci.image.layer.nondistributable.v1.tar\napplication/vnd.oci.image.layer.nondistributable.v1.tar+gzip"]
  }

  imageIndex -> imageIndex [label="1..*"]
  imageIndex -> manifest [label="1..*"]
  manifest -> config [label="1..1"]
  manifest -> layer [label="1..*"]
  manifest -> manifest [label="0..1"];
}
~~~
```

---

## "Artifacts"

Given the backwards evolution of OCI from image-spec to distribution-spec, packaging any objects that are not container specific feel like a hack.

But folks have been doing so for a while.

And registries arbitrarily chose how to blocklist or allowlist `mediaTypes`.

But not stopped masquerading `mediaTypes`.
Even Red Hat needing to publish src.rpms "in like manner" as the container images built from their binaries, in [BuildSourceImage](https://github.com/containers/BuildSourceImage).

[Guidance for an empty descriptor](https://github.com/opencontainers/image-spec/blob/main/manifest.md#guidance-for-an-empty-descriptor)

---

## Hands-on ...

```bash
crane --help
```

```
~~~crane --help
replaceme
~~~
```

---

## Hands-on ...

```bash
crane manifest docker.io/alpine
```
```json
{
  "manifests": [
    {
      "digest": "sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33",
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "platform": {
        "architecture": "amd64",
        "os": "linux"
      },
      "size": 528
    },
    {
      "digest": "sha256:f748290eb66ad6f938e25dd348acfb3527a422e280b7547b1cdfaf38d4492c4b",
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "platform": {
        "architecture": "arm",
        "os": "linux",
        "variant": "v6"
      },
      "size": 528
    },
[...]
```

```bash
crane -v manifest docker.io/alpine
```

---

## Hands-on ...

```bash
crane manifest docker.io/alpine@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33
```
```json
{
  "schemaVersion": 2,
  "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
  "config": {
    "mediaType": "application/vnd.docker.container.image.v1+json",
    "size": 1471,
    "digest": "sha256:7e01a0d0a1dcd9e539f8e9bbd80106d59efbdf97293b3d38f5d7a34501526cdb"
  },
  "layers": [
    {
      "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
      "size": 3401613,
      "digest": "sha256:7264a8db6415046d36d16ba98b79778e18accee6ffa71850405994cffa9be7de"
    }
  ]
}
```

---

## Hands-on ...

```bash
crane config docker.io/alpine@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33
```

Show the contents of the `application/vnd.docker.container.image.v1+json` from the manifest

---

## edit

```bash
docker pull docker.io/alpine
docker tag docker.io/alpine r.batts.cloud/alpine-test:latest
docker push r.batts.cloud/alpine-test:latest
crane edit manifest r.batts.cloud/alpine-test:latest
```

use `application/vnd.oci.image.manifest.v1+json`, `application/vnd.oci.image.config.v1+json`, and `application/vnd.oci.image.layer.v1.tar+gzip`.

```bash
docker rmi r.batts.cloud/alpine-test:latest
docker pull r.batts.cloud/alpine-test@sha256:b0050a40f615a0225485155f238034ac1cfd581e8071ba78aed5ca272e4dc9bb
docker run --rm r.batts.cloud/alpine-test@sha256:b0050a40f615a0225485155f238034ac1cfd581e8071ba78aed5ca272e4dc9bb echo howdy
crane manifest r.batts.cloud/alpine-test@sha256:b0050a40f615a0225485155f238034ac1cfd581e8071ba78aed5ca272e4dc9bb | jq .
crane manifest r.batts.cloud/alpine-test@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33 | jq .
crane delete r.batts.cloud/alpine-test@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33
crane delete r.batts.cloud/alpine-test@sha256:b0050a40f615a0225485155f238034ac1cfd581e8071ba78aed5ca272e4dc9bb

```

---

## links

- [past talks](https://github.com/vbatts/talks)
- [slide tool used](https://github.com/maaslalani/slides)

