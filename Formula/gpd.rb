class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  version "v0.3.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.2/gpd-darwin-arm64"
      sha256 "bed3abaf5c11b25056dd6cb66cddf9a1dc1ccd49ebe23f567b30359ec10a369d"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.2/gpd-darwin-amd64"
      sha256 "c7ea65b5982651e6063fd01a933e62ea7f828beb05093bcdcca9cb06b2c1114f"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.2/gpd-linux-arm64"
      sha256 "cc3579ad705cff20de8ac3a752a493bd4b3fc7c32c10f9f303c4df7dd1e21e78"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.2/gpd-linux-amd64"
      sha256 "f049c38713049469b35c911f3d8d7373947ea88e26917ec8c9c12e0ad279d5a5"
    end
  end

  def install
    bin.install "gpd-darwin-arm64" => "gpd" if OS.mac? && Hardware::CPU.arm?
    bin.install "gpd-darwin-amd64" => "gpd" if OS.mac? && Hardware::CPU.intel?
    bin.install "gpd-linux-arm64" => "gpd" if OS.linux? && Hardware::CPU.arm?
    bin.install "gpd-linux-amd64" => "gpd" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/gpd", "version"
  end
end
