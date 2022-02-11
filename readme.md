# Git Interactive/Improved Commit Experience

*ability to go back might be a good feature*

To override default config path, path to config file in local config:

`git config --local ice.config git-ice.toml`

set personal config 

`git config --global ice.localconfig ~/.git-ice.toml` Note that this can be specified per-project with --local instead of --global

to remove ice config

`git config --local --remove-section ice.config`