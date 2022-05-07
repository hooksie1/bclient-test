go_binary(
    name="bin",
)

go_package()

docker_image(
    name="bclient",
    repository="hooksie1/bclient",
    image_tags=["{build_args.GIT_BRANCH}-{pants.hash}", "latest"],
    dependencies=[":bin"],
)
