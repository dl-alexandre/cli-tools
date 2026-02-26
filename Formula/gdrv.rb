class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.5.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.2/gdrv-darwin-arm64.tar.gz"
      sha256 "0ee0274943256a5d17604bea4e7cf49f6dd0270b0a526dae9caf731865fc2734"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.2/gdrv-darwin-amd64.tar.gz"
      sha256 "f03242d26e34078c33dad0864ffc40daad6b9f729834be1a38f065a170f87d4e"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.2/gdrv-linux-arm64.tar.gz"
      sha256 "177ebdab929f2d71ead577b87eff95e9bf581e04311322e97996796c0cd54d30"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.2/gdrv-linux-amd64.tar.gz"
      sha256 "23bc781dd41ecdbc3c476da1646be81f130256e0979f5c2567727904f93ad099"
    end
  end

  def install
    bin.install "gdrv"
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
