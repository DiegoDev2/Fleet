package main

// Rtorrent - Ncurses BitTorrent client based on libtorrent-rakshasa
// Homepage: https://github.com/rakshasa/rtorrent

import (
	"fmt"
	
	"os/exec"
)

func installRtorrent() {
	// Método 1: Descargar y extraer .tar.gz
	rtorrent_tar_url := "https://github.com/rakshasa/rtorrent/releases/download/v0.9.8/rtorrent-0.9.8.tar.gz"
	rtorrent_cmd_tar := exec.Command("curl", "-L", rtorrent_tar_url, "-o", "package.tar.gz")
	err := rtorrent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rtorrent_zip_url := "https://github.com/rakshasa/rtorrent/releases/download/v0.9.8/rtorrent-0.9.8.zip"
	rtorrent_cmd_zip := exec.Command("curl", "-L", rtorrent_zip_url, "-o", "package.zip")
	err = rtorrent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rtorrent_bin_url := "https://github.com/rakshasa/rtorrent/releases/download/v0.9.8/rtorrent-0.9.8.bin"
	rtorrent_cmd_bin := exec.Command("curl", "-L", rtorrent_bin_url, "-o", "binary.bin")
	err = rtorrent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rtorrent_src_url := "https://github.com/rakshasa/rtorrent/releases/download/v0.9.8/rtorrent-0.9.8.src.tar.gz"
	rtorrent_cmd_src := exec.Command("curl", "-L", rtorrent_src_url, "-o", "source.tar.gz")
	err = rtorrent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rtorrent_cmd_direct := exec.Command("./binary")
	err = rtorrent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libtorrent-rakshasa")
	exec.Command("latte", "install", "libtorrent-rakshasa").Run()
	fmt.Println("Instalando dependencia: xmlrpc-c")
	exec.Command("latte", "install", "xmlrpc-c").Run()
}
