class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  version "v0.2.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.1/gpd-darwin-arm64"
      sha256 "458df1c93840ba54787332a242d05262d87a228f7291a5ed57e92dc54d1f88bd"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.1/gpd-darwin-amd64"
      sha256 "e2bb22bf4fe6e38d844c5c10be355d3088715babee820141e572d2dfd4376ea0"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.1/gpd-linux-arm64"
      sha256 "99c6bf9104bddc7edce1ea1fcf92966ff74c0b0e3983ae904885509a5dcb1cd8"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.2.1/gpd-linux-amd64"
      sha256 "33e4e8b4735c537673af6a6d4ec428c804c2b5293ea2a167aa775ab6e35d0873"
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
