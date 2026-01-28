class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v3.1.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.0/gdrv-darwin-arm64"
      sha256 "624d932a2ae0af4aa76c524823192c7a95bae0cfc2e4b39eb1d030f94482d89f"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.0/gdrv-darwin-amd64"
      sha256 "226196f4c1f58bb8635db2eed72ec48c7478c55f2be1f943609dc3aef10916f5"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.0/gdrv-linux-arm64"
      sha256 "fabc27421cb6ff6ea18a24c9fa5f6d610cd7d23dfac93f064d34b2e136900fff"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.0/gdrv-linux-amd64"
      sha256 "cf8716cbd38ac81877771906f21548b4c033fa80bcc8aa7b369ba7a07ab5f82c"
    end
  end

  def install
    bin.install "gdrv-darwin-arm64" => "gdrv" if OS.mac? && Hardware::CPU.arm?
    bin.install "gdrv-darwin-amd64" => "gdrv" if OS.mac? && Hardware::CPU.intel?
    bin.install "gdrv-linux-arm64" => "gdrv" if OS.linux? && Hardware::CPU.arm?
    bin.install "gdrv-linux-amd64" => "gdrv" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
