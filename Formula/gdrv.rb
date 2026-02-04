class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.4.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.0/gdrv-darwin-arm64"
      sha256 "b5299e5ef264dd39e2cd9b23fe799af8cf68cf2e92fd1df6090d784d5d8738b1"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.0/gdrv-darwin-amd64"
      sha256 "8275f81e8590aabf66341fd41e716af00808e9a398f71be5b3e7ebc6c34d8a38"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.0/gdrv-linux-arm64"
      sha256 "f8f2794b6db9fa5e8d9b33fe61b3250b4625fae0b587b2cc0dd977a58cbf0d0f"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.0/gdrv-linux-amd64"
      sha256 "30b2a115814fc91de6823267494217071858bb488ac34b2ec9c2e8e246ff4bfb"
    end
  end

  def install
    bin.install "gdrv-darwin-arm64" => "gdrv" if OS.mac? && Hardware::CPU.arm?
    bin.install "gdrv-darwin-amd64" => "gdrv" if OS.mac? && Hardware::CPU.intel?
    bin.install "gdrv-linux-arm64" => "gdrv" if OS.linux? && Hardware::CPU.arm?
    bin.install "gdrv-linux-amd64" => "gdrv" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
