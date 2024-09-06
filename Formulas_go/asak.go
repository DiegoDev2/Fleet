package main

// Asak - Cross-platform audio recording/playback CLI tool with TUI
// Homepage: https://github.com/chaosprint/asak

import (
	"fmt"
	
	"os/exec"
)

func installAsak() {
	// Método 1: Descargar y extraer .tar.gz
	asak_tar_url := "https://github.com/chaosprint/asak/archive/refs/tags/v0.3.3.tar.gz"
	asak_cmd_tar := exec.Command("curl", "-L", asak_tar_url, "-o", "package.tar.gz")
	err := asak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asak_zip_url := "https://github.com/chaosprint/asak/archive/refs/tags/v0.3.3.zip"
	asak_cmd_zip := exec.Command("curl", "-L", asak_zip_url, "-o", "package.zip")
	err = asak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asak_bin_url := "https://github.com/chaosprint/asak/archive/refs/tags/v0.3.3.bin"
	asak_cmd_bin := exec.Command("curl", "-L", asak_bin_url, "-o", "binary.bin")
	err = asak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asak_src_url := "https://github.com/chaosprint/asak/archive/refs/tags/v0.3.3.src.tar.gz"
	asak_cmd_src := exec.Command("curl", "-L", asak_src_url, "-o", "source.tar.gz")
	err = asak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asak_cmd_direct := exec.Command("./binary")
	err = asak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
