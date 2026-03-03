class Unifi < Formula
  desc "Local UniFi Controller CLI"
  homepage "https://github.com/dl-alexandre/Local-UniFi-CLI"
  version "v0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.1/unifi-darwin-arm64"
      sha256 "PLACEHOLDER_SHA256"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.1/unifi-darwin-amd64"
      sha256 "PLACEHOLDER_SHA256"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.1/unifi-linux-arm64"
      sha256 "PLACEHOLDER_SHA256"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.1/unifi-linux-amd64"
      sha256 "PLACEHOLDER_SHA256"
    end
  end

  def install
    bin.install "unifi-darwin-arm64" => "unifi" if OS.mac? && Hardware::CPU.arm?
    bin.install "unifi-darwin-amd64" => "unifi" if OS.mac? && Hardware::CPU.intel?
    bin.install "unifi-linux-arm64" => "unifi" if OS.linux? && Hardware::CPU.arm?
    bin.install "unifi-linux-amd64" => "unifi" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/unifi", "version"
  end
end
