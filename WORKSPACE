workspace(name = "adhoc")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Go + Gazelle
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "685052b498b6ddfe562ca7a97736741d87916fe536623afb7da2824c0211c369",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "5982e5463f171da99e3bdaeff8c0f48283a7a5f396ec5282910b9e8a49c0dd7e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "2d9084d3dd13b86ca2e811d2331f780eb86f6d7cb02b405426e3c80dcbfabf25",
    strip_prefix = "protobuf-3.21.1",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.21.1.zip"],
)


load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_rules_dependencies()

go_register_toolchains(version = "1.18.3")

gazelle_dependencies()

# gazelle:repository go_repository name=org_golang_x_xerrors importpath=golang.org/x/xerrors

# Python
rules_python_version = "0.10.2" # Latest @ 2021-06-23

http_archive(
  name = "rules_python",
  sha256 = "e963db3a1893d3ef61869f2f9d438e09aa1f671ec20be9ebae29ad5c31b4a50c",
  strip_prefix = "rules_python-{}".format(rules_python_version),
  url = "https://github.com/bazelbuild/rules_python/archive/{}.zip".format(rules_python_version),
)

load("@rules_python//python:repositories.bzl", "python_register_toolchains")

python_register_toolchains(
  name = "python3_9",
  # Available versions are listed in @rules_python//python:versions.bzl.
  # We recommend using the same version your team is already standardized on.
  python_version = "3.9",
)

load("@python3_9//:defs.bzl", "interpreter")

load("@rules_python//python:pip.bzl", "pip_parse")

pip_parse(
  name = "py_deps",
  requirements_lock = "//:pip_lock.txt",
  python_interpreter_target = interpreter,
)