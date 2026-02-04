class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v3.1.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.1/gdrv-darwin-arm64"
      sha256 "ae3016df964a732606806e1d414e9eeb55250a5b25306774e177384dbb3a42bc"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.1/gdrv-darwin-amd64"
      sha256 "17da08445881abd04b1ee5ef268f8205975f6dce372168b72ff1c71e060924bc"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.1/gdrv-linux-arm64"
      sha256 "491924a230dd413fc2239939a78f9f97759d81a99e5e1c8fc69f10afbe145b60"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v3.1.1/gdrv-linux-amd64"
      sha256 "5cec3f569e655f0314df6a6ebc584d4d2ed84f2381fa01e47a853fc550dcae64"
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
