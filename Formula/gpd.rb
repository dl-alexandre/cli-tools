class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  version "v0.3.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.1/gpd-darwin-arm64"
      sha256 "2906b5e558ef8fb1cacd3d2b7ee3ff699c8eaa9990f38fd8e9a9a64a2c23aacd"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.1/gpd-darwin-amd64"
      sha256 "fdc54bebcbeed4f494eda6e2648ac1a16f4aecc0479b48e298cfa7ec133b3e4f"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.1/gpd-linux-arm64"
      sha256 "4793ce1b9f3ac46a39ec9f421a0f09d409cce73925520eb2522c32ae79f4e28f"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Play-Developer-CLI/releases/download/v0.3.1/gpd-linux-amd64"
      sha256 "6175239d620af952552d1fef152a016c5d469c17aa07d6fa129020502b3220c0"
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
