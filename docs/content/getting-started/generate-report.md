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
      github:numtide/nixos-facter -- -o facter.json
    ```

This will scan your system and produce a JSON-based report in a file named `facter.json`:

```json title="facter.json"
{
  "version": 2, // (1)!
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

    To use this report in your NixOS configuration, have a look at [NixOS Facter Modules].

[Nix]: https://nixos.org
[Numtide]: https://numtide.com
[Numtide Binary Cache]: https://numtide.cachix.org
[nixos-facter]: https://github.com/nix-community/nixos-facter
[nixpkgs]: https://github.com/nixos/nixpkgs
[System Management BIOS]: https://wiki.osdev.org/System_Management_BIOS
[NixOS Facter Modules]: https://github.com/nix-community/nixos-facter-modules
