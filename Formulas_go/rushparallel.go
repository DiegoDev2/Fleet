package main

// RushParallel - Cross-platform command-line tool for executing jobs in parallel
// Homepage: https://github.com/shenwei356/rush

import (
	"fmt"
	
	"os/exec"
)

func installRushParallel() {
	// Método 1: Descargar y extraer .tar.gz
	rushparallel_tar_url := "https://github.com/shenwei356/rush/archive/refs/tags/v0.5.6.tar.gz"
	rushparallel_cmd_tar := exec.Command("curl", "-L", rushparallel_tar_url, "-o", "package.tar.gz")
	err := rushparallel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rushparallel_zip_url := "https://github.com/shenwei356/rush/archive/refs/tags/v0.5.6.zip"
	rushparallel_cmd_zip := exec.Command("curl", "-L", rushparallel_zip_url, "-o", "package.zip")
	err = rushparallel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rushparallel_bin_url := "https://github.com/shenwei356/rush/archive/refs/tags/v0.5.6.bin"
	rushparallel_cmd_bin := exec.Command("curl", "-L", rushparallel_bin_url, "-o", "binary.bin")
	err = rushparallel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rushparallel_src_url := "https://github.com/shenwei356/rush/archive/refs/tags/v0.5.6.src.tar.gz"
	rushparallel_cmd_src := exec.Command("curl", "-L", rushparallel_src_url, "-o", "source.tar.gz")
	err = rushparallel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rushparallel_cmd_direct := exec.Command("./binary")
	err = rushparallel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
