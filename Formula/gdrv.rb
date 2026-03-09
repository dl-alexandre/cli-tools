class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.6.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-darwin-arm64.tar.gz"
      sha256 "b5248b1d11a173525b53e3fc0fba0f45b327210c095600baefa2559f5ca58160"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-darwin-amd64.tar.gz"
      sha256 "9337c8198f735744fa0c4d139a5df1656d7e2acac2673ed9ca9de8225b05c323"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-linux-arm64.tar.gz"
      sha256 "10bb3e9d03d329ae3b7631eabaea7ee5b1c3e5e5c5126dad4d4ff862b69d59df"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-linux-amd64.tar.gz"
      sha256 "c6814cade75859644419240cbd8b6fbef49e78d3164d069472e91eec98555d12"
    end
  end

  def install
    bin.install "gdrv"
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
