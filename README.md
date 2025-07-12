# NixOS Facter

NixOS Facter aims to be an alternative to projects such as [NixOS Hardware] and [nixos-generate-config].
It solves the problem of bootstrapping [NixOS configurations] by deferring decisions about hardware and other
aspects of the target platform to NixOS modules.

We do this by first generating a machine-readable report (JSON) which captures detailed information about the machine
or virtual environment it was executed within.

This report is then passed to a series of [NixOS modules] which can make a variety of decisions,
some simple, some more complex, enabling things like automatic configuration of network controllers or graphics cards,
USB devices, and so on.

## Project Structure

This repository contains the binary for generating the report.

[NixOS Facter Modules] contains the necessary NixOS modules for making use of the report in a NixOS configuration.

For more information, please see the [docs].

## Quick Start

To generate a report using `nixos-facter` from nixpkgs:

```console
# you must run this as root
❯ sudo nix run --option experimental-features "nix-command flakes" nixpkgs#nixos-facter -- -o facter.json
```

To use the latest development version from this flake:

```console
# you must run this as root
❯ sudo nix run \
  --option experimental-features "nix-command flakes" \
  --option extra-substituters https://numtide.cachix.org \
  --option extra-trusted-public-keys numtide.cachix.org-1:2ps1kLBUWjxIneOy1Ik6cQjb41X0iXVXeHigGmycPPE= \
  github:nix-community/nixos-facter -- -o facter.json
```

## Contributing

Contributions are always welcome!

## License

This software is provided free under [GNU GPL v3].

---

This project is supported by [Clan] and [Numtide](https://numtide.com/).

## Clan

[Clan] is a peer-to-peer management framework based on [NixOS] that has:

- a uniform interface
- automated secret management
- automated service setup
- automated backups
- peer-to-peer Mesh VPN

Get in touch and learn more at:

- [clan.lol](https://clan.lol)
- [Matrix channel](https://matrix.to/#/#clan:clan.lol) for live discussions.
- IRC bridge on [hackint#clan](https://chat.hackint.org/#/connect?join=clan) for real-time chat support.

## Numtide

![Numtide Logo](https://codahosted.io/docs/6FCIMTRM0p/blobs/bl-sgSunaXYWX/077f3f9d7d76d6a228a937afa0658292584dedb5b852a8ca370b6c61dabb7872b7f617e603f1793928dc5410c74b3e77af21a89e435fa71a681a868d21fd1f599dd10a647dd855e14043979f1df7956f67c3260c0442e24b34662307204b83ea34de929d)

We’re a team of independent freelancers that love open source.
We help our customers make their project lifecycles more efficient by:

- Providing and supporting useful tools such as this one.
- Building and deploying infrastructure, and offering dedicated DevOps support.
- Building their in-house Nix skills, and integrating Nix with their workflows.
- Developing additional features and tools.
- Carrying out custom research and development.

[Contact us](https://numtide.com/contact) if you have a project in mind,
or if you need help with any of our supported tools, including this one.

We'd love to hear from you.

[Clan]: https://clan.lol
[NixOS configurations]: https://nixos.org/manual/nixos/stable/#sec-configuration-syntax
[NixOS Hardware]: https://github.com/NixOS/nixos-hardware
[NixOS Facter Modules]: https://github.com/nix-community/nixos-facter-modules
[NixOS modules]: https://github.com/nix-community/nixos-facter-modules
[nixos-generate-config]: https://github.com/NixOS/nixpkgs/blob/master/nixos/modules/installer/tools/nixos-generate-config.pl
[Numtide Binary Cache]: https://numtide.cachix.org
[nixos-facter]: https://github.com/nix-community/nixos-facter
[nixpkgs]: https://github.com/nixos/nixpkgs
[docs]: https://nix-community.github.io/nixos-facter
[GNU GPL v3]: ./LICENSE
[NixOS]: https://nixos.org
