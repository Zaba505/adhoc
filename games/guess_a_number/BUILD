load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@rules_python//python:defs.bzl", "py_binary")

py_binary(
    name = "python",
    srcs = ["guess_a_number.py"],
    main = "guess_a_number.py",
)

go_binary(
    name = "go",
    embed = [":guess_a_number_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "guess_a_number_lib",
    srcs = ["main.go"],
    importpath = "github.com/Zaba505/adhoc/games/guess_a_number",
    visibility = ["//visibility:private"],
)

go_test(
    name = "guess_a_number_test",
    srcs = ["main_test.go"],
    embed = [":guess_a_number_lib"],
)
