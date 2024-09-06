package main

// Blockhash - Perceptual image hash calculation tool
// Homepage: https://github.com/commonsmachinery/blockhash

import (
	"fmt"
	
	"os/exec"
)

func installBlockhash() {
	// Método 1: Descargar y extraer .tar.gz
	blockhash_tar_url := "https://github.com/commonsmachinery/blockhash/archive/refs/tags/v0.3.3.tar.gz"
	blockhash_cmd_tar := exec.Command("curl", "-L", blockhash_tar_url, "-o", "package.tar.gz")
	err := blockhash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blockhash_zip_url := "https://github.com/commonsmachinery/blockhash/archive/refs/tags/v0.3.3.zip"
	blockhash_cmd_zip := exec.Command("curl", "-L", blockhash_zip_url, "-o", "package.zip")
	err = blockhash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blockhash_bin_url := "https://github.com/commonsmachinery/blockhash/archive/refs/tags/v0.3.3.bin"
	blockhash_cmd_bin := exec.Command("curl", "-L", blockhash_bin_url, "-o", "binary.bin")
	err = blockhash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blockhash_src_url := "https://github.com/commonsmachinery/blockhash/archive/refs/tags/v0.3.3.src.tar.gz"
	blockhash_cmd_src := exec.Command("curl", "-L", blockhash_src_url, "-o", "source.tar.gz")
	err = blockhash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blockhash_cmd_direct := exec.Command("./binary")
	err = blockhash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
}
