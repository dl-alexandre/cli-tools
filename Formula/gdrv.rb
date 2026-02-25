class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.5.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-darwin-arm64.tar.gz"
      sha256 "01cbafc52c115bcf1c8dd6bf5e60de4fdb8a64e4cad5c90ba0511153d4697d28"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-darwin-amd64.tar.gz"
      sha256 ""
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-linux-arm64.tar.gz"
      sha256 ""
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-linux-amd64.tar.gz"
      sha256 ""
    end
  end

  def install
    bin.install "gdrv"
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
