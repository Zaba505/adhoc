load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@rules_python//python:defs.bzl", "py_binary")

py_binary(
    name = "python",
    srcs = ["guess_a_number.py"],
    main = "guess_a_number.py",
)

go_binary(
    name = "go",
    srcs = ["main.go"],
)