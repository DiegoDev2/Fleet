package main

// Waon - Wave-to-notes transcriber
// Homepage: https://kichiki.github.io/WaoN/

import (
	"fmt"
	
	"os/exec"
)

func installWaon() {
	// Método 1: Descargar y extraer .tar.gz
	waon_tar_url := "https://github.com/kichiki/WaoN/archive/refs/tags/v0.11.tar.gz"
	waon_cmd_tar := exec.Command("curl", "-L", waon_tar_url, "-o", "package.tar.gz")
	err := waon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	waon_zip_url := "https://github.com/kichiki/WaoN/archive/refs/tags/v0.11.zip"
	waon_cmd_zip := exec.Command("curl", "-L", waon_zip_url, "-o", "package.zip")
	err = waon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	waon_bin_url := "https://github.com/kichiki/WaoN/archive/refs/tags/v0.11.bin"
	waon_cmd_bin := exec.Command("curl", "-L", waon_bin_url, "-o", "binary.bin")
	err = waon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	waon_src_url := "https://github.com/kichiki/WaoN/archive/refs/tags/v0.11.src.tar.gz"
	waon_cmd_src := exec.Command("curl", "-L", waon_src_url, "-o", "source.tar.gz")
	err = waon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	waon_cmd_direct := exec.Command("./binary")
	err = waon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sox")
exec.Command("latte", "install", "sox")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
}
