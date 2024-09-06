package main

// Normalize - Adjust volume of audio files to a standard level
// Homepage: https://www.nongnu.org/normalize/

import (
	"fmt"
	
	"os/exec"
)

func installNormalize() {
	// Método 1: Descargar y extraer .tar.gz
	normalize_tar_url := "https://savannah.nongnu.org/download/normalize/normalize-0.7.7.tar.gz"
	normalize_cmd_tar := exec.Command("curl", "-L", normalize_tar_url, "-o", "package.tar.gz")
	err := normalize_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	normalize_zip_url := "https://savannah.nongnu.org/download/normalize/normalize-0.7.7.zip"
	normalize_cmd_zip := exec.Command("curl", "-L", normalize_zip_url, "-o", "package.zip")
	err = normalize_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	normalize_bin_url := "https://savannah.nongnu.org/download/normalize/normalize-0.7.7.bin"
	normalize_cmd_bin := exec.Command("curl", "-L", normalize_bin_url, "-o", "binary.bin")
	err = normalize_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	normalize_src_url := "https://savannah.nongnu.org/download/normalize/normalize-0.7.7.src.tar.gz"
	normalize_cmd_src := exec.Command("curl", "-L", normalize_src_url, "-o", "source.tar.gz")
	err = normalize_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	normalize_cmd_direct := exec.Command("./binary")
	err = normalize_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mad")
exec.Command("latte", "install", "mad")
}
