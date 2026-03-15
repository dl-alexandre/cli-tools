class Aca < Formula
  desc "Advance Commerce CLI - Apple App Store Connect and commerce management"
  homepage "https://github.com/dl-alexandre/Advance-Commerce-CLI"
  version "v0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Advance-Commerce-CLI/releases/download/v0.0.1/aca-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Advance-Commerce-CLI/releases/download/v0.0.1/aca-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Advance-Commerce-CLI/releases/download/v0.0.1/aca-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Advance-Commerce-CLI/releases/download/v0.0.1/aca-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  def install
    bin.install "aca"
  end

  test do
    system "#{bin}/aca", "version"
  end
end
