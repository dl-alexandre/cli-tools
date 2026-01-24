class Gdrive < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.2.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.0/gdrive-darwin-arm64"
      sha256 "643afba14e68108f13acfe069ff212acb7c1d9593e4416acaf27958aab57bb27"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.0/gdrive-darwin-amd64"
      sha256 "e68995e919f87468b093262f67143b3eee89c72d355b14bade5764e41d462d0f"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.0/gdrive-linux-arm64"
      sha256 "91ae39487d030251b79ffd7199090972ba3bed53dec2cefba36f198ac4c08621"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.0/gdrive-linux-amd64"
      sha256 "69dfae30bdae4cdf915ee13a5f94157658436bcf274de485458488f47deb7a7c"
    end
  end

  def install
    bin.install "gdrive-darwin-arm64" => "gdrive" if OS.mac? && Hardware::CPU.arm?
    bin.install "gdrive-darwin-amd64" => "gdrive" if OS.mac? && Hardware::CPU.intel?
    bin.install "gdrive-linux-arm64" => "gdrive" if OS.linux? && Hardware::CPU.arm?
    bin.install "gdrive-linux-amd64" => "gdrive" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/gdrive", "version"
  end
end
