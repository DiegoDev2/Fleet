package main

// Bwa - Burrow-Wheeler Aligner for pairwise alignment of DNA
// Homepage: https://github.com/lh3/bwa

import (
	"fmt"
	
	"os/exec"
)

func installBwa() {
	// Método 1: Descargar y extraer .tar.gz
	bwa_tar_url := "https://github.com/lh3/bwa/archive/refs/tags/v0.7.18.tar.gz"
	bwa_cmd_tar := exec.Command("curl", "-L", bwa_tar_url, "-o", "package.tar.gz")
	err := bwa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bwa_zip_url := "https://github.com/lh3/bwa/archive/refs/tags/v0.7.18.zip"
	bwa_cmd_zip := exec.Command("curl", "-L", bwa_zip_url, "-o", "package.zip")
	err = bwa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bwa_bin_url := "https://github.com/lh3/bwa/archive/refs/tags/v0.7.18.bin"
	bwa_cmd_bin := exec.Command("curl", "-L", bwa_bin_url, "-o", "binary.bin")
	err = bwa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bwa_src_url := "https://github.com/lh3/bwa/archive/refs/tags/v0.7.18.src.tar.gz"
	bwa_cmd_src := exec.Command("curl", "-L", bwa_src_url, "-o", "source.tar.gz")
	err = bwa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bwa_cmd_direct := exec.Command("./binary")
	err = bwa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sse2neon")
	exec.Command("latte", "install", "sse2neon").Run()
}
