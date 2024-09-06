package main

// Sickle - Windowed adaptive trimming for FASTQ files using quality
// Homepage: https://github.com/najoshi/sickle

import (
	"fmt"
	
	"os/exec"
)

func installSickle() {
	// Método 1: Descargar y extraer .tar.gz
	sickle_tar_url := "https://github.com/najoshi/sickle/archive/refs/tags/v1.33.tar.gz"
	sickle_cmd_tar := exec.Command("curl", "-L", sickle_tar_url, "-o", "package.tar.gz")
	err := sickle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sickle_zip_url := "https://github.com/najoshi/sickle/archive/refs/tags/v1.33.zip"
	sickle_cmd_zip := exec.Command("curl", "-L", sickle_zip_url, "-o", "package.zip")
	err = sickle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sickle_bin_url := "https://github.com/najoshi/sickle/archive/refs/tags/v1.33.bin"
	sickle_cmd_bin := exec.Command("curl", "-L", sickle_bin_url, "-o", "binary.bin")
	err = sickle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sickle_src_url := "https://github.com/najoshi/sickle/archive/refs/tags/v1.33.src.tar.gz"
	sickle_cmd_src := exec.Command("curl", "-L", sickle_src_url, "-o", "source.tar.gz")
	err = sickle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sickle_cmd_direct := exec.Command("./binary")
	err = sickle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
