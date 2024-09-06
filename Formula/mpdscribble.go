package main

// Mpdscribble - Last.fm reporting client for mpd
// Homepage: https://www.musicpd.org/clients/mpdscribble/

import (
	"fmt"
	
	"os/exec"
)

func installMpdscribble() {
	// Método 1: Descargar y extraer .tar.gz
	mpdscribble_tar_url := "https://www.musicpd.org/download/mpdscribble/0.25/mpdscribble-0.25.tar.xz"
	mpdscribble_cmd_tar := exec.Command("curl", "-L", mpdscribble_tar_url, "-o", "package.tar.gz")
	err := mpdscribble_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpdscribble_zip_url := "https://www.musicpd.org/download/mpdscribble/0.25/mpdscribble-0.25.tar.xz"
	mpdscribble_cmd_zip := exec.Command("curl", "-L", mpdscribble_zip_url, "-o", "package.zip")
	err = mpdscribble_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpdscribble_bin_url := "https://www.musicpd.org/download/mpdscribble/0.25/mpdscribble-0.25.tar.xz"
	mpdscribble_cmd_bin := exec.Command("curl", "-L", mpdscribble_bin_url, "-o", "binary.bin")
	err = mpdscribble_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpdscribble_src_url := "https://www.musicpd.org/download/mpdscribble/0.25/mpdscribble-0.25.tar.xz"
	mpdscribble_cmd_src := exec.Command("curl", "-L", mpdscribble_src_url, "-o", "source.tar.gz")
	err = mpdscribble_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpdscribble_cmd_direct := exec.Command("./binary")
	err = mpdscribble_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libmpdclient")
	exec.Command("latte", "install", "libmpdclient").Run()
}
