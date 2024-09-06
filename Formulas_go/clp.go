package main

// Clp - Linear programming solver
// Homepage: https://github.com/coin-or/Clp

import (
	"fmt"
	
	"os/exec"
)

func installClp() {
	// Método 1: Descargar y extraer .tar.gz
	clp_tar_url := "https://github.com/coin-or/Clp/archive/refs/tags/releases/1.17.10.tar.gz"
	clp_cmd_tar := exec.Command("curl", "-L", clp_tar_url, "-o", "package.tar.gz")
	err := clp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clp_zip_url := "https://github.com/coin-or/Clp/archive/refs/tags/releases/1.17.10.zip"
	clp_cmd_zip := exec.Command("curl", "-L", clp_zip_url, "-o", "package.zip")
	err = clp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clp_bin_url := "https://github.com/coin-or/Clp/archive/refs/tags/releases/1.17.10.bin"
	clp_cmd_bin := exec.Command("curl", "-L", clp_bin_url, "-o", "binary.bin")
	err = clp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clp_src_url := "https://github.com/coin-or/Clp/archive/refs/tags/releases/1.17.10.src.tar.gz"
	clp_cmd_src := exec.Command("curl", "-L", clp_src_url, "-o", "source.tar.gz")
	err = clp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clp_cmd_direct := exec.Command("./binary")
	err = clp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: coinutils")
exec.Command("latte", "install", "coinutils")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: osi")
exec.Command("latte", "install", "osi")
}
