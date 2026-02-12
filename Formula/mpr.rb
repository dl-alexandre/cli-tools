class Mpr < Formula
  desc "USDA Market News CLI for agricultural commodity data"
  homepage "https://github.com/dl-alexandre/MyMarketNews-CLI"
  version "v0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.1/mpr-darwin-arm64.tar.gz"
      sha256 "f8493bcd1f6d2d853624e4f1f3b6f4aacc12b7357caec03ae281b469b4262b7c"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.1/mpr-darwin-amd64.tar.gz"
      sha256 "b6a480aaa51694b241122974430851b3d64df14479505fefa6b17c9f8c9a339d"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.1/mpr-linux-arm64.tar.gz"
      sha256 "56854b52767610641c7739a63f7fdc14877d7e489d492abdcb0ed1167b1ce627"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/MyMarketNews-CLI/releases/download/v0.0.1/mpr-linux-amd64.tar.gz"
      sha256 "8f2d834883b8668cef5efe3c57d3e04e364df07d211758332962c45447762cc0"
    end
  end

  def install
    bin.install "mpr"
  end

  test do
    system "#{bin}/mpr", "--version"
  end
end
