package main

// Kew - Command-line music player
// Homepage: https://github.com/ravachol/kew

import (
	"fmt"
	
	"os/exec"
)

func installKew() {
	// Método 1: Descargar y extraer .tar.gz
	kew_tar_url := "https://github.com/ravachol/kew/archive/refs/tags/v2.7.2.tar.gz"
	kew_cmd_tar := exec.Command("curl", "-L", kew_tar_url, "-o", "package.tar.gz")
	err := kew_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kew_zip_url := "https://github.com/ravachol/kew/archive/refs/tags/v2.7.2.zip"
	kew_cmd_zip := exec.Command("curl", "-L", kew_zip_url, "-o", "package.zip")
	err = kew_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kew_bin_url := "https://github.com/ravachol/kew/archive/refs/tags/v2.7.2.bin"
	kew_cmd_bin := exec.Command("curl", "-L", kew_bin_url, "-o", "binary.bin")
	err = kew_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kew_src_url := "https://github.com/ravachol/kew/archive/refs/tags/v2.7.2.src.tar.gz"
	kew_cmd_src := exec.Command("curl", "-L", kew_src_url, "-o", "source.tar.gz")
	err = kew_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kew_cmd_direct := exec.Command("./binary")
	err = kew_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: chafa")
	exec.Command("latte", "install", "chafa").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: freeimage")
	exec.Command("latte", "install", "freeimage").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: opusfile")
	exec.Command("latte", "install", "opusfile").Run()
}
