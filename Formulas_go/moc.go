package main

// Moc - Terminal-based music player
// Homepage: https://moc.daper.net/

import (
	"fmt"
	
	"os/exec"
)

func installMoc() {
	// Método 1: Descargar y extraer .tar.gz
	moc_tar_url := "https://ftp.daper.net/pub/soft/moc/stable/moc-2.5.2.tar.bz2"
	moc_cmd_tar := exec.Command("curl", "-L", moc_tar_url, "-o", "package.tar.gz")
	err := moc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moc_zip_url := "https://ftp.daper.net/pub/soft/moc/stable/moc-2.5.2.tar.bz2"
	moc_cmd_zip := exec.Command("curl", "-L", moc_zip_url, "-o", "package.zip")
	err = moc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moc_bin_url := "https://ftp.daper.net/pub/soft/moc/stable/moc-2.5.2.tar.bz2"
	moc_cmd_bin := exec.Command("curl", "-L", moc_bin_url, "-o", "binary.bin")
	err = moc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moc_src_url := "https://ftp.daper.net/pub/soft/moc/stable/moc-2.5.2.tar.bz2"
	moc_cmd_src := exec.Command("curl", "-L", moc_src_url, "-o", "source.tar.gz")
	err = moc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moc_cmd_direct := exec.Command("./binary")
	err = moc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
	fmt.Println("Instalando dependencia: ffmpeg@4")
exec.Command("latte", "install", "ffmpeg@4")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: speex")
exec.Command("latte", "install", "speex")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
