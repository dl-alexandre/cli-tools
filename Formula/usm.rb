class Usm < Formula
  desc "UniFi Site Manager CLI for cloud-based site management"
  homepage "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI"
  version "v1.1.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v1.1.1/usm-darwin-arm64"
      sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v1.1.1/usm-darwin-amd64"
      sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v1.1.1/usm-linux-arm64"
      sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v1.1.1/usm-linux-amd64"
      sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5"
    end
  end

  def install
    bin.install "usm-darwin-arm64" => "usm" if OS.mac? && Hardware::CPU.arm?
    bin.install "usm-darwin-amd64" => "usm" if OS.mac? && Hardware::CPU.intel?
    bin.install "usm-linux-arm64" => "usm" if OS.linux? && Hardware::CPU.arm?
    bin.install "usm-linux-amd64" => "usm" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/usm", "version"
  end
end
