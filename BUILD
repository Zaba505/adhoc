load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/Zaba505/adhoc
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=go_dependencies.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
