class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.2.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-darwin-arm64"
      sha256 "c728a5da5901d5a495fad2990b9ef311473b31e556619a862ea9bb3ff8f2a631"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-darwin-amd64"
      sha256 "2501af998752ccb3798efd3571df6a7df34238661f62b11862bc6a3fbf27bc25"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-linux-arm64"
      sha256 "9f4430d30628d1e03a84cdb37bba73320a8975ac4c2ac28220ab7ea1815ac708"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-linux-amd64"
      sha256 "93dd7a6dde7ea2f5c9f7316bf9dcb10145beb60d7caa1a3598d0b6766687c7e1"
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
