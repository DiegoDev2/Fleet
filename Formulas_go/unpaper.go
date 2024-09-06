package main

// Unpaper - Post-processing for scanned/photocopied books
// Homepage: https://www.flameeyes.com/projects/unpaper

import (
	"fmt"
	
	"os/exec"
)

func installUnpaper() {
	// Método 1: Descargar y extraer .tar.gz
	unpaper_tar_url := "https://www.flameeyes.com/files/unpaper-7.0.0.tar.xz"
	unpaper_cmd_tar := exec.Command("curl", "-L", unpaper_tar_url, "-o", "package.tar.gz")
	err := unpaper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unpaper_zip_url := "https://www.flameeyes.com/files/unpaper-7.0.0.tar.xz"
	unpaper_cmd_zip := exec.Command("curl", "-L", unpaper_zip_url, "-o", "package.zip")
	err = unpaper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unpaper_bin_url := "https://www.flameeyes.com/files/unpaper-7.0.0.tar.xz"
	unpaper_cmd_bin := exec.Command("curl", "-L", unpaper_bin_url, "-o", "binary.bin")
	err = unpaper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unpaper_src_url := "https://www.flameeyes.com/files/unpaper-7.0.0.tar.xz"
	unpaper_cmd_src := exec.Command("curl", "-L", unpaper_src_url, "-o", "source.tar.gz")
	err = unpaper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unpaper_cmd_direct := exec.Command("./binary")
	err = unpaper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
}
