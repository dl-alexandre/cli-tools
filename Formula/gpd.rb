class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  version "v3.0.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v3.0.0/gpd-darwin-arm64"
      sha256 "154b597af961883c8e01283401869f8220886948c539b8025c4fb7e7f5e93be8"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v3.0.0/gpd-darwin-amd64"
      sha256 "11005a358b2257c6d0cf202cd6ec60e3d0e2172a9eb549db5abe41dec22319a2"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v3.0.0/gpd-linux-arm64"
      sha256 "568f49f887e7d024a930a04d8576a95b5327fd34e657210fa9c687e5dc200537"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v3.0.0/gpd-linux-amd64"
      sha256 "0e5f20deb958448a39da8a08e5ca9e4b6ec985a9c353f8912f9688738c87facc"
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
