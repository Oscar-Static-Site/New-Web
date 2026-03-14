{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};

      node = pkgs.nodejs_20;
      pnpm = node.pkgs.pnpm;

      # Create a derivation for the astro project directory
      astroDir = pkgs.runCommand "astro-project"
        {
          nativeBuildInputs = [ pkgs.makeWrapper ];
          buildInputs = [ node ];
        }
        ''
          mkdir -p "$out"
          cp -r ${./astro} "$out/"
          cp ${./pnpm-lock.yaml} "$out/"
        '';
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        name = "New-Web Development Environment";

        buildInputs = with pkgs; [
          node
          pnpm
          nodePackages.typescript
          nodePackages.prettier
          nodePackages.eslint
        ];

        shellHook = ''
          echo "Welcome to the New-Web development environment!"
          echo "=============================================="
          echo "Available commands:"
          echo "  pnpm run dev       - Start development server"
          echo "  pnpm run build     - Build the site (includes pagefind)"
          echo "  pnpm run preview   - Preview built site"
          echo "  pnpm run format    - Format code with prettier"
          echo "  pnpm run check     - Run TypeScript type checking"
          echo ""
          echo "Note: Run 'pnpm install' first if dependencies are missing"
        '';

        DENO_ENABLE = 1;
      };

      packages.${system}.default = astroDir;

      apps.${system}.dev = {
        type = "app";
        program = "${pkgs.writeShellScriptBin "astro-dev" ''
          export PATH="${node}/bin:${pnpm}/bin:$PATH"
          cd ./astro/
          if [ ! -d "node_modules" ]; then
            pnpm install --frozen-lockfile
          fi
          exec pnpm run dev
        ''}/bin/astro-dev";
      };
    };
}

