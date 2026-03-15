class Ups < Formula
  desc "UPS CLI - Track packages and manage UPS shipments"
  homepage "https://github.com/dl-alexandre/UPS-CLI"
  version "v0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/UPS-CLI/releases/download/v0.0.1/ups-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/UPS-CLI/releases/download/v0.0.1/ups-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/UPS-CLI/releases/download/v0.0.1/ups-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/UPS-CLI/releases/download/v0.0.1/ups-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  def install
    bin.install "ups"
  end

  test do
    system "#{bin}/ups", "version"
  end
end
