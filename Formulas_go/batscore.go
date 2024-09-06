package main

// BatsCore - Bash Automated Testing System
// Homepage: https://github.com/bats-core/bats-core

import (
	"fmt"
	
	"os/exec"
)

func installBatsCore() {
	// Método 1: Descargar y extraer .tar.gz
	batscore_tar_url := "https://github.com/bats-core/bats-core/archive/refs/tags/v1.11.0.tar.gz"
	batscore_cmd_tar := exec.Command("curl", "-L", batscore_tar_url, "-o", "package.tar.gz")
	err := batscore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	batscore_zip_url := "https://github.com/bats-core/bats-core/archive/refs/tags/v1.11.0.zip"
	batscore_cmd_zip := exec.Command("curl", "-L", batscore_zip_url, "-o", "package.zip")
	err = batscore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	batscore_bin_url := "https://github.com/bats-core/bats-core/archive/refs/tags/v1.11.0.bin"
	batscore_cmd_bin := exec.Command("curl", "-L", batscore_bin_url, "-o", "binary.bin")
	err = batscore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	batscore_src_url := "https://github.com/bats-core/bats-core/archive/refs/tags/v1.11.0.src.tar.gz"
	batscore_cmd_src := exec.Command("curl", "-L", batscore_src_url, "-o", "source.tar.gz")
	err = batscore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	batscore_cmd_direct := exec.Command("./binary")
	err = batscore_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
}
