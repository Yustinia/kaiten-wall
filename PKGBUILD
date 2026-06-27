# Maintainer: Yustinia <yustinia@tuta.io>
pkgname=kaiten
pkgver=0.4.3
pkgrel=1
pkgdesc="A wayland wallpaper switcher"
arch=('x86_64')
url="https://github.com/Yustinia/kaiten-wall"
license=('MIT')
depends=('awww')
makedepends=('go' 'git')
optdepends=('hyprland'
    'niri'
    'matugen'
    'wallust')
options=('!strip' '!debug')
source=("$pkgname-$pkgver::git+https://github.com/Yustinia/kaiten-wall.git#tag=v$pkgver"
    "LICENSE::https://raw.githubusercontent.com/Yustinia/kaiten-wall/refs/heads/main/LICENSE")
sha256sums=('SKIP'
    'SKIP')

build() {
    cd "$srcdir/$pkgname-$pkgver"
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o kaiten ./cmd/
}

package() {
    install -Dm755 "$srcdir/$pkgname-$pkgver/kaiten" "$pkgdir/usr/bin/$pkgname"
    install -Dm644 "$srcdir/LICENSE" "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
