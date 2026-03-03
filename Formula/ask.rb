class Ask < Formula
  desc "App Store Kit CLI for in-app purchases and subscriptions"
  homepage "https://github.com/dl-alexandre/App-StoreKit-CLI"
  version "v0.0.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/App-StoreKit-CLI/releases/download/v0.0.2/ask-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "ask"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/App-StoreKit-CLI/releases/download/v0.0.2/ask-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "ask"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/dl-alexandre/App-StoreKit-CLI/releases/download/v0.0.2/ask-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "ask"
      end
    end
    if Hardware::CPU.intel? && Hardware::CPU.is_64_bit?
      url "https://github.com/dl-alexandre/App-StoreKit-CLI/releases/download/v0.0.2/ask-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"

      def install
        bin.install "ask"
      end
    end
  end

  test do
    system "#{bin}/ask", "version"
  end
end
