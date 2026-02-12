class Ams < Formula
  desc "Apple Maps Server API CLI for location services"
  homepage "https://github.com/dl-alexandre/Apple-Map-Server-CLI"
  url "https://github.com/dl-alexandre/Apple-Map-Server-CLI/archive/refs/tags/v0.0.2.tar.gz"
  sha256 "c0ffecb01f957e95caf589927c2bca7a8dca7569a56837e7ea5f0f276ba25562"
  version "v0.0.2"
  license "MIT"

  depends_on "go" => :build

  def install
    cd "cmd/ams" do
      system "go", "build", "-o", bin/"ams", "."
    end
  end

  test do
    system "#{bin}/ams", "version"
  end
end
