package main

// Libmpd - Higher level access to MPD functions
// Homepage: https://gmpc.fandom.com/wiki/Gnome_Music_Player_Client

import (
	"fmt"
	
	"os/exec"
)

func installLibmpd() {
	// Método 1: Descargar y extraer .tar.gz
	libmpd_tar_url := "https://www.musicpd.org/download/libmpd/11.8.17/libmpd-11.8.17.tar.gz"
	libmpd_cmd_tar := exec.Command("curl", "-L", libmpd_tar_url, "-o", "package.tar.gz")
	err := libmpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmpd_zip_url := "https://www.musicpd.org/download/libmpd/11.8.17/libmpd-11.8.17.zip"
	libmpd_cmd_zip := exec.Command("curl", "-L", libmpd_zip_url, "-o", "package.zip")
	err = libmpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmpd_bin_url := "https://www.musicpd.org/download/libmpd/11.8.17/libmpd-11.8.17.bin"
	libmpd_cmd_bin := exec.Command("curl", "-L", libmpd_bin_url, "-o", "binary.bin")
	err = libmpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmpd_src_url := "https://www.musicpd.org/download/libmpd/11.8.17/libmpd-11.8.17.src.tar.gz"
	libmpd_cmd_src := exec.Command("curl", "-L", libmpd_src_url, "-o", "source.tar.gz")
	err = libmpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmpd_cmd_direct := exec.Command("./binary")
	err = libmpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
