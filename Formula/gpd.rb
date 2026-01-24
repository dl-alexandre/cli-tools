class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  version "v0.2.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-darwin-arm64"
      sha256 "8cc8b044384c9e58464e90b7b83a48c43622e1eab8a4f3be54d422bdc47af77e"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-darwin-amd64"
      sha256 "784108b4c8c1b46c2612001529fbb71f4317d61b71894b94fce06e63f03b4116"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-linux-arm64"
      sha256 "7ed8b08df194d4e799d07692416d7cdf1d5047808714b94ef78afb0e051e61e6"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-linux-amd64"
      sha256 "275a151d61bae4e7cfd6068d38edfaa27a80db74cd8358343ffbe05836d2e10a"
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
