load("@rules_python//python:defs.bzl", "py_binary")
load("@pip//:requirements.bzl", "requirement")

py_binary(
    name = "kvscluster",
    srcs = [
        "main.py",
    ],
    main = "main.py",
    deps = [
        requirement("hdbscan"),
        requirement("numpy"),
        requirement("scikit_learn"),
    ],
)

py_binary(
    name = "demo",
    srcs = [
        "demo.py",
    ],
    main = "demo.py",
    deps = [
        requirement("hdbscan"),
        requirement("scikit_learn"),
    ],
)