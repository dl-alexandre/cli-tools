class Mpr < Formula
  desc "USDA Market News CLI for agricultural commodity data"
  homepage "https://github.com/dl-alexandre/MyMarketNews-CLI"
  version "v0.0.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.2/mpr-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.2/mpr-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.2/mpr-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.2/mpr-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  def install
    bin.install "mpr"
  end

  test do
    system "#{bin}/mpr", "--version"
  end
end
