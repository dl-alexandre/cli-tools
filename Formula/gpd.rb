class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  version "v0.3.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.0/gpd-darwin-arm64"
      sha256 "63c8d3bc8543786c96183bf5cb669b91fc74db80a454c771a42e4860ff92bec3"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.0/gpd-darwin-amd64"
      sha256 "12241a3cd4d946e4fde60d954d0ff6c4d8a5b93f62525a3f9736453a6295abec"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.0/gpd-linux-arm64"
      sha256 "b8078bde82c90b9402de3e32d507cae60159d218fb5dea4c7c1e285b7528e1cb"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.0/gpd-linux-amd64"
      sha256 "f7505e5289d57813db95bb503836f62e385bac4dd3d3a76d53a34f4ef75b5b08"
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
