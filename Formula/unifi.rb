class Unifi < Formula
  desc "Local UniFi Controller CLI"
  homepage "https://github.com/dl-alexandre/Local-UniFi-CLI"
  version "v0.0.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-darwin-arm64"
      sha256 "c423aea0cf303db174bfdad2184661898ccab52e11c9b8a4906f350dd6d64d14"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-darwin-amd64"
      sha256 "4392fb499c84d4590b717cc6d1f1490eda214c34ac80b4f134d399579e8ffa27"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-linux-arm64"
      sha256 "a18168cfa8671d08f12e31f7c8488f670d7f61ea24f025bd8b7881b2ff869390"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-linux-amd64"
      sha256 "52fa76ba3b3d0a2a94198217a232c448432d90610d7c4e4e3b29f4319cfe3148"
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
