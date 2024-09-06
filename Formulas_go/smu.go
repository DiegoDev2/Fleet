package main

// Smu - Simple markup with markdown-like syntax
// Homepage: https://github.com/Gottox/smu

import (
	"fmt"
	
	"os/exec"
)

func installSmu() {
	// Método 1: Descargar y extraer .tar.gz
	smu_tar_url := "https://github.com/Gottox/smu/archive/refs/tags/v1.5.tar.gz"
	smu_cmd_tar := exec.Command("curl", "-L", smu_tar_url, "-o", "package.tar.gz")
	err := smu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smu_zip_url := "https://github.com/Gottox/smu/archive/refs/tags/v1.5.zip"
	smu_cmd_zip := exec.Command("curl", "-L", smu_zip_url, "-o", "package.zip")
	err = smu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smu_bin_url := "https://github.com/Gottox/smu/archive/refs/tags/v1.5.bin"
	smu_cmd_bin := exec.Command("curl", "-L", smu_bin_url, "-o", "binary.bin")
	err = smu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smu_src_url := "https://github.com/Gottox/smu/archive/refs/tags/v1.5.src.tar.gz"
	smu_cmd_src := exec.Command("curl", "-L", smu_src_url, "-o", "source.tar.gz")
	err = smu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smu_cmd_direct := exec.Command("./binary")
	err = smu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
