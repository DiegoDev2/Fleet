package main

// FsUae - Amiga emulator
// Homepage: https://fs-uae.net/

import (
	"fmt"
	
	"os/exec"
)

func installFsUae() {
	// Método 1: Descargar y extraer .tar.gz
	fsuae_tar_url := "https://fs-uae.net/files/FS-UAE/Stable/3.1.66/fs-uae-3.1.66.tar.xz"
	fsuae_cmd_tar := exec.Command("curl", "-L", fsuae_tar_url, "-o", "package.tar.gz")
	err := fsuae_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fsuae_zip_url := "https://fs-uae.net/files/FS-UAE/Stable/3.1.66/fs-uae-3.1.66.tar.xz"
	fsuae_cmd_zip := exec.Command("curl", "-L", fsuae_zip_url, "-o", "package.zip")
	err = fsuae_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fsuae_bin_url := "https://fs-uae.net/files/FS-UAE/Stable/3.1.66/fs-uae-3.1.66.tar.xz"
	fsuae_cmd_bin := exec.Command("curl", "-L", fsuae_bin_url, "-o", "binary.bin")
	err = fsuae_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fsuae_src_url := "https://fs-uae.net/files/FS-UAE/Stable/3.1.66/fs-uae-3.1.66.tar.xz"
	fsuae_cmd_src := exec.Command("curl", "-L", fsuae_src_url, "-o", "source.tar.gz")
	err = fsuae_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fsuae_cmd_direct := exec.Command("./binary")
	err = fsuae_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libmpeg2")
exec.Command("latte", "install", "libmpeg2")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: openal-soft")
exec.Command("latte", "install", "openal-soft")
}
