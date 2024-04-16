{
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = import nixpkgs {
            inherit system;
          };
        in
        {
          devShell = with pkgs; mkShell {
            buildInputs = [
              git
              go_1_22
              templ
            ];
          };
        }
      );
}
