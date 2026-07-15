# Generate a report

To generate a report, you will need to have [Nix] installed on the target machine.

=== "Nixpkgs"

    ```shell
    sudo nix run nixpkgs#nixos-facter -- -o facter.json
    ```

=== "Flake"

    ```shell
    sudo nix run \
      --option experimental-features "nix-command flakes" \
      --option extra-substituters https://numtide.cachix.org \
      --option extra-trusted-public-keys numtide.cachix.org-1:2ps1kLBUWjxIneOy1Ik6cQjb41X0iXVXeHigGmycPPE= \
      github:nix-community/nixos-facter -- -o facter.json
    ```

This will scan your system and produce a JSON-based report in a file named `facter.json`:

```json title="facter.json"
{
  "version": 1, // (1)!
  "system": "x86_64-linux", // (2)!
  "virtualisation": "none", // (3)!
  "hardware": { // (4)!
    "bios": { ... },
    "bluetooth": [ ... ],
    "bridge": [ ... ],
    "chip_card": [ ... ] ,
    "cpu": [ ... ],
    "disk": [ ... ],
    "graphics_card": [ ... ],
    "hub": [ ... ],
    "keyboard": [ ... ],
    "memory": [ ... ],
    "monitor": [ ... ],
    "mouse": [ ... ],
    "network_controller": [ ... ],
    "network_interface": [ ... ],
    "sound": [ ... ],
    "storage_controller": [ ... ],
    "system": [ ... ],
    "unknown": [ ... ],
    "usb_controller": [ ... ]
  },
  "smbios": { // (5)!
    "bios": { ... },
    "board": { ... },
    "cache": [ ... ],
    "chassis": { ... },
    "config": { ... },
    "language": { ... },
    "memory_array": [ ... ],
    "memory_array_mapped_address": [ ... ],
    "memory_device": [ ... ],
    "memory_device_mapped_address": [ ... ],
    "memory_error": [ ... ],
    "onboard": [ ... ],
    "port_connector": [ ... ],
    "processor": [ ... ],
    "slot": [ ... ],
    "system": { ... }
  }
}
```

1. Used to track major breaking changes in the report format.
2. Architecture of the target machine.
3. Indicates whether the report was generated inside a virtualised environment, and if so, what type.
4. All the various bits of hardware that could be detected.
5. [System Management BIOS] information if available.

!!! tip

    To use this report in your NixOS configuration, add the following to your configuration:

    ```nix
    {
      hardware.facter.reportPath = ./facter.json;
    }
    ```

    See the [nixpkgs documentation](https://search.nixos.org/options?query=facter) for more details.

## Cloud provider metadata

### Hetzner

When running on a [Hetzner](https://www.hetzner.com) instance, facter can capture instance metadata (e.g. the assigned IPv6 network
configuration) from the Hetzner metadata service. It is disabled by default and must be enabled with a flag:

```shell
sudo nixos-facter --cloud-hetzner -o facter.json
```

The captured metadata is included in the report under a `cloud` section:

```json title="facter.json"
{
  "cloud": {
    "hetzner": {
      "hostname": "my-server",
      "instance-id": 123456,
      "region": "eu-central",
      ...
    }
  }
}
```

[Nix]: https://nixos.org
[Numtide]: https://numtide.com
[Numtide Binary Cache]: https://numtide.cachix.org
[nixos-facter]: https://github.com/nix-community/nixos-facter
[nixpkgs]: https://github.com/nixos/nixpkgs
[System Management BIOS]: https://wiki.osdev.org/System_Management_BIOS
