class Unifi < Formula
  desc "Local UniFi Controller CLI"
  homepage "https://github.com/dl-alexandre/Local-UniFi-CLI"
  version "v0.0.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "unifi"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "unifi"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "unifi"
      end
    end
    if Hardware::CPU.intel? && Hardware::CPU.is_64_bit?
      url "https://github.com/dl-alexandre/Local-UniFi-CLI/releases/download/v0.0.2/unifi-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "unifi"
      end
    end
  end

  test do
    system "#{bin}/unifi", "version"
  end
end
