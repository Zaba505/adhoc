load("@pip//:requirements.bzl", "all_whl_requirements")
load("@rules_python//gazelle/manifest:defs.bzl", "gazelle_python_manifest")
load("@rules_python//gazelle/modules_mapping:def.bzl", "modules_mapping")
load("@rules_python//gazelle:def.bzl", "GAZELLE_PYTHON_RUNTIME_DEPS")
load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")

gazelle_binary(
    name = "gazelle_binary",
    languages = [
        "@rules_python//gazelle",  # Use gazelle from rules_python.
        "@bazel_gazelle//language/go",  # Built-in rule from gazelle for Golang.
        "@bazel_gazelle//language/proto",  # Built-in rule from gazelle for Protos.
        # Any languages that depend on Gazelle's proto plugin must come after it.
        # "@external_repository//language/gazelle",  External languages can be added here.
    ],
    visibility = ["//visibility:public"],
)

# gazelle:prefix github.com/Zaba505/adhoc
gazelle(
    name = "gazelle",
    data = GAZELLE_PYTHON_RUNTIME_DEPS,
    gazelle = "//:gazelle_binary",
)

# Python
# This rule fetches the metadata for python packages we depend on. That data is
# required for the gazelle_python_manifest rule to update our manifest file.
modules_mapping(
    name = "modules_map",
    wheels = all_whl_requirements,
)

# Gazelle python extension needs a manifest file mapping from
# an import to the installed package that provides it.
# This macro produces two targets:
# - //:gazelle_python_manifest.update can be used with `bazel run`
#   to recalculate the manifest
# - //:gazelle_python_manifest.test is a test target ensuring that
#   the manifest doesn't need to be updated
gazelle_python_manifest(
    name = "gazelle_python_manifest",
    modules_mapping = ":modules_map",
    # This is what we called our `pip_install` rule, where third-party
    # python libraries are loaded in BUILD files.
    pip_repository_name = "pip",
    # When using pip_parse instead of pip_install, set the following.
    pip_repository_incremental = True,
    # This should point to wherever we declare our python dependencies
    # (the same as what we passed to the modules_mapping rule in WORKSPACE)
    requirements = "//:requirements_lock.txt",
)
