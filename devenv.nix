{pkgs, ...}: {
  packages = with pkgs; [
    turso-cli
    goose
    bootdev-cli
  ];
}
