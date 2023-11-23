{
  description = "Daily dose of Data Structure & Algorithm";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = { self, nixpkgs }:
    let
      forAllSystems = nixpkgs.lib.genAttrs nixpkgs.lib.systems.flakeExposed;
    in
    {
      devShells = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShell {
            name = "minmax";
            shellHook = ''
              git config pull.rebase true
              ${pkgs.neo-cowsay}/bin/cowsay -f sage -n "Daily dose of Data Structure & Algorithm"
            '';
            buildInputs = with pkgs; [
              editorconfig-checker
              go
            ];
          };
        }
      );
    };
}
