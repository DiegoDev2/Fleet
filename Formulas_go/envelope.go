package main

// Envelope - Environment variables CLI tool
// Homepage: https://github.com/mattrighetti/envelope

import (
	"fmt"
	
	"os/exec"
)

func installEnvelope() {
	// Método 1: Descargar y extraer .tar.gz
	envelope_tar_url := "https://github.com/mattrighetti/envelope/archive/refs/tags/0.3.11.tar.gz"
	envelope_cmd_tar := exec.Command("curl", "-L", envelope_tar_url, "-o", "package.tar.gz")
	err := envelope_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	envelope_zip_url := "https://github.com/mattrighetti/envelope/archive/refs/tags/0.3.11.zip"
	envelope_cmd_zip := exec.Command("curl", "-L", envelope_zip_url, "-o", "package.zip")
	err = envelope_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	envelope_bin_url := "https://github.com/mattrighetti/envelope/archive/refs/tags/0.3.11.bin"
	envelope_cmd_bin := exec.Command("curl", "-L", envelope_bin_url, "-o", "binary.bin")
	err = envelope_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	envelope_src_url := "https://github.com/mattrighetti/envelope/archive/refs/tags/0.3.11.src.tar.gz"
	envelope_cmd_src := exec.Command("curl", "-L", envelope_src_url, "-o", "source.tar.gz")
	err = envelope_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	envelope_cmd_direct := exec.Command("./binary")
	err = envelope_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
