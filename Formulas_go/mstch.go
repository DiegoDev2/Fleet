package main

// Mstch - Complete implementation of {{mustache}} templates using modern C++
// Homepage: https://github.com/no1msd/mstch

import (
	"fmt"
	
	"os/exec"
)

func installMstch() {
	// Método 1: Descargar y extraer .tar.gz
	mstch_tar_url := "https://github.com/no1msd/mstch/archive/refs/tags/1.0.2.tar.gz"
	mstch_cmd_tar := exec.Command("curl", "-L", mstch_tar_url, "-o", "package.tar.gz")
	err := mstch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mstch_zip_url := "https://github.com/no1msd/mstch/archive/refs/tags/1.0.2.zip"
	mstch_cmd_zip := exec.Command("curl", "-L", mstch_zip_url, "-o", "package.zip")
	err = mstch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mstch_bin_url := "https://github.com/no1msd/mstch/archive/refs/tags/1.0.2.bin"
	mstch_cmd_bin := exec.Command("curl", "-L", mstch_bin_url, "-o", "binary.bin")
	err = mstch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mstch_src_url := "https://github.com/no1msd/mstch/archive/refs/tags/1.0.2.src.tar.gz"
	mstch_cmd_src := exec.Command("curl", "-L", mstch_src_url, "-o", "source.tar.gz")
	err = mstch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mstch_cmd_direct := exec.Command("./binary")
	err = mstch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
