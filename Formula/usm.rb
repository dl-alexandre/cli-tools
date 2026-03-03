class Usm < Formula
  desc "UniFi Site Manager CLI for cloud-based site management"
  homepage "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI"
  version "v0.0.3"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "usm"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "usm"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "usm"
      end
    end
    if Hardware::CPU.intel? && Hardware::CPU.is_64_bit?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "usm"
      end
    end
  end

  test do
    system "#{bin}/usm", "version"
  end
end
