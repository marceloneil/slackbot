### Instructions for getting the Keybase buildbot running on Linux

- Create an account called "keybasebuild": `sudo useradd -m
  keybasebuild`
- Add user "keybasebuild" to the "docker" group: `sudo gpasswd -a
  keybasebuild docker`.
- Get an SSH key with access to the keybase GitHub repos and the keybase
  account on aur.archlinux.org.
- Clone several keybase repos into /home/keybasebuild:
  - client
  - kbfs
  - kbfs-beta
  - server-ops
  - slackbot
- Import the code signing PGP secret key. After import, remove the
  password from this key. (TODO: Something more interesting with
  yubikeys.)
- Set up s3cmd (~/.s3cfg) with credentials for
  s3://prerelease.keybase.io.
- Create /home/keybasebuild/keybot.env with the following lines:

    ```
    SLACK_TOKEN=<slack token here>
    SLACK_CHANNEL=bot
    ```

- Start/enable-on-boot the systemd service files.
  - `sudo loginctl enable-linger keybasebuild`
  - With a proper login session as keybasebuild (SSH is one way):
    - `mkdir -p ~/.config/systemd/user`
    - `cp ~/slackbot/systemd/keybase.*.{service,timer} ~/.config/systemd/user/`
    - `systemctl --user enable --now keybase.keybot.service`
    - `systemctl --user enable --now keybase.buildplease.timer`
- Take the bot out of dry-run mode by messaging `!tuxbot toggle-dryrun`