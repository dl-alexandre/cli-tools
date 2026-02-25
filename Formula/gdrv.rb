class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.5.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-darwin-arm64.tar.gz"
      sha256 "bbf87d6aa901986a8cb0b5f292dc7d83c22802c7163b2cfc91d19e392c3d3f82"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-darwin-amd64.tar.gz"
      sha256 ""
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-linux-arm64.tar.gz"
      sha256 ""
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-linux-amd64.tar.gz"
      sha256 ""
    end
  end

  def install
    bin.install "gdrv"
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
