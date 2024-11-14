# Maintainer: Your Name <your.email@domain.com>
pkgname=portico
pkgver=0.1.0
pkgrel=1
pkgdesc="Remote Access Display Manager"
arch=('x86_64')
url="https://github.com/yourusername/portico"
license=('MIT')
depends=('libx11' 'libxrandr' 'pam' 'xorg-server')
makedepends=('go' 'git')
backup=('etc/portico/config.yaml')
source=("git+https://github.com/yourusername/portico.git")
sha256sums=('SKIP')

build() {
    cd "$srcdir/$pkgname"
    go build \
        -trimpath \
        -ldflags "-s -w" \
        ./cmd/portico
}

package() {
    cd "$srcdir/$pkgname"
    
    # Binary
    install -Dm755 portico "$pkgdir/usr/bin/portico"
    
    # Service file
    install -Dm644 configs/portico.service "$pkgdir/usr/lib/systemd/system/portico.service"
    
    # Config
    install -Dm644 configs/portico.yaml "$pkgdir/etc/portico/config.yaml"
}
