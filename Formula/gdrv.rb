class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.5.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.0/gdrv-darwin-arm64"
      sha256 "556b9eaebb1c7aa07cab2ef34d8852e4c758ce15d7d2676ed1fbcde03d2e5bc4"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.0/gdrv-darwin-amd64"
      sha256 "922d2fbcd512023f419a841175a829f4371abfb8fda9d7b99d4f1fb192780836"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.0/gdrv-linux-arm64"
      sha256 "646bf862f40b68bcecf0f8b9cbfc5a0d3e235f331c9b610122f353cb4550569d"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.0/gdrv-linux-amd64"
      sha256 "e57a7e14f11c428315dc04b4ba601544d16810e6eff622049b24303570f856e4"
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
