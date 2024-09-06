package main

// Osi - Open Solver Interface
// Homepage: https://github.com/coin-or/Osi

import (
	"fmt"
	
	"os/exec"
)

func installOsi() {
	// Método 1: Descargar y extraer .tar.gz
	osi_tar_url := "https://github.com/coin-or/Osi/archive/refs/tags/releases/0.108.11.tar.gz"
	osi_cmd_tar := exec.Command("curl", "-L", osi_tar_url, "-o", "package.tar.gz")
	err := osi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osi_zip_url := "https://github.com/coin-or/Osi/archive/refs/tags/releases/0.108.11.zip"
	osi_cmd_zip := exec.Command("curl", "-L", osi_zip_url, "-o", "package.zip")
	err = osi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osi_bin_url := "https://github.com/coin-or/Osi/archive/refs/tags/releases/0.108.11.bin"
	osi_cmd_bin := exec.Command("curl", "-L", osi_bin_url, "-o", "binary.bin")
	err = osi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osi_src_url := "https://github.com/coin-or/Osi/archive/refs/tags/releases/0.108.11.src.tar.gz"
	osi_cmd_src := exec.Command("curl", "-L", osi_src_url, "-o", "source.tar.gz")
	err = osi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osi_cmd_direct := exec.Command("./binary")
	err = osi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: coinutils")
	exec.Command("latte", "install", "coinutils").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
