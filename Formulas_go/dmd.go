package main

// Dmd - Digital Mars D compiler
// Homepage: https://dlang.org/

import (
	"fmt"
	
	"os/exec"
)

func installDmd() {
	// Método 1: Descargar y extraer .tar.gz
	dmd_tar_url := "https://github.com/dlang/dmd/archive/refs/tags/v2.109.1.tar.gz"
	dmd_cmd_tar := exec.Command("curl", "-L", dmd_tar_url, "-o", "package.tar.gz")
	err := dmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dmd_zip_url := "https://github.com/dlang/dmd/archive/refs/tags/v2.109.1.zip"
	dmd_cmd_zip := exec.Command("curl", "-L", dmd_zip_url, "-o", "package.zip")
	err = dmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dmd_bin_url := "https://github.com/dlang/dmd/archive/refs/tags/v2.109.1.bin"
	dmd_cmd_bin := exec.Command("curl", "-L", dmd_bin_url, "-o", "binary.bin")
	err = dmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dmd_src_url := "https://github.com/dlang/dmd/archive/refs/tags/v2.109.1.src.tar.gz"
	dmd_cmd_src := exec.Command("curl", "-L", dmd_src_url, "-o", "source.tar.gz")
	err = dmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dmd_cmd_direct := exec.Command("./binary")
	err = dmd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ldc")
exec.Command("latte", "install", "ldc")
}
