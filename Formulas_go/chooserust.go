package main

// ChooseRust - Human-friendly and fast alternative to cut and (sometimes) awk
// Homepage: https://github.com/theryangeary/choose

import (
	"fmt"
	
	"os/exec"
)

func installChooseRust() {
	// Método 1: Descargar y extraer .tar.gz
	chooserust_tar_url := "https://github.com/theryangeary/choose/archive/refs/tags/v1.3.4.tar.gz"
	chooserust_cmd_tar := exec.Command("curl", "-L", chooserust_tar_url, "-o", "package.tar.gz")
	err := chooserust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chooserust_zip_url := "https://github.com/theryangeary/choose/archive/refs/tags/v1.3.4.zip"
	chooserust_cmd_zip := exec.Command("curl", "-L", chooserust_zip_url, "-o", "package.zip")
	err = chooserust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chooserust_bin_url := "https://github.com/theryangeary/choose/archive/refs/tags/v1.3.4.bin"
	chooserust_cmd_bin := exec.Command("curl", "-L", chooserust_bin_url, "-o", "binary.bin")
	err = chooserust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chooserust_src_url := "https://github.com/theryangeary/choose/archive/refs/tags/v1.3.4.src.tar.gz"
	chooserust_cmd_src := exec.Command("curl", "-L", chooserust_src_url, "-o", "source.tar.gz")
	err = chooserust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chooserust_cmd_direct := exec.Command("./binary")
	err = chooserust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
