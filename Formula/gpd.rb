class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  version "v0.2.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-darwin-arm64"
      sha256 "11a21c5957b0a4cbb7972e7e5580af3bf746c2ea601a2ec1781fd70af69b349b"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-darwin-amd64"
      sha256 "60e420f548c51c5080b988a293405c670e59d62c8763489d2bc543dd1429ea4b"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-linux-arm64"
      sha256 "491f13481aefe306bcca6f4b95491ac3506266a11e1620307e116736ff929303"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.0/gpd-linux-amd64"
      sha256 "7213ef6f03784f8754a5aee9b1478e9053a6e63ecd17646b9eca9a524bd82db3"
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
