class Cpw < Formula
  desc "File watcher that copies changed files to a destination"
  homepage "https://github.com/mxvsh/cpw"
  version "VERSION_PLACEHOLDER"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/mxvsh/cpw/releases/download/vVERSION_PLACEHOLDER/cpw-darwin-arm64"
      sha256 "SHA256_PLACEHOLDER_DARWIN_ARM64"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/mxvsh/cpw/releases/download/vVERSION_PLACEHOLDER/cpw-linux-arm64"
      sha256 "SHA256_PLACEHOLDER_LINUX_ARM64"
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/mxvsh/cpw/releases/download/vVERSION_PLACEHOLDER/cpw-linux-armv7"
      sha256 "SHA256_PLACEHOLDER_LINUX_ARM"
    end
  end

  def install
    bin.install Dir["cpw-*"].first => "cpw"
  end

  test do
    system "#{bin}/cpw", "-version"
  end
end 