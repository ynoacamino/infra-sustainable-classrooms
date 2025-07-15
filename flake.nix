{
  description = "A flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs =
    { nixpkgs, ... }:
    let
      system = "x86_64-linux";
    in
    {
      devShells."${system}".default =
        let
          pkgs = import nixpkgs {
            inherit system;
            config.allowUnfree = true;
          };
        in
        pkgs.mkShell {
          packages = with pkgs; [
            go
            sqlc
            goa
            protobuf
            protoc-gen-go
            protoc-gen-go-grpc
            pnpm
            nodejs
          ];
          buildInputs = [ pkgs.bashInteractive ];
          shellHook = ''
            export GOROOT="${pkgs.go}/share/go"
          '';
        };
    };
}
