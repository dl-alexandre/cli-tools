class Ask < Formula
  desc "App Store Kit CLI for in-app purchases and subscriptions"
  homepage "https://github.com/dl-alexandre/App-StoreKit-CLI"
  url "https://github.com/dl-alexandre/App-StoreKit-CLI/archive/refs/tags/v0.0.2.tar.gz"
  sha256 "e3ba4bc74331a2d80a694fe623d8cfa22410fd244aaad9959f2ecfaa20f5b5a0"
  version "v0.0.2"
  license "MIT"

  depends_on "go" => :build

  def install
    cd "cmd/ask" do
      system "go", "build", "-o", bin/"ask", "."
    end
  end

  test do
    system "#{bin}/ask", "version"
  end
end
