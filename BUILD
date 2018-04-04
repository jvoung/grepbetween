load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/jvoung/gobetween",
)

go_library(
    name = "print_between",
    srcs = ["print_between.go"],
    importpath = "github.com/jvoung/gobetween",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["print_between_test.go"],
    embed = [":print_between"],
)
