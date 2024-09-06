package main

// Jtbl - Convert JSON and JSON Lines to terminal, CSV, HTTP, and markdown tables
// Homepage: https://github.com/kellyjonbrazil/jtbl

import (
	"fmt"
	
	"os/exec"
)

func installJtbl() {
	// Método 1: Descargar y extraer .tar.gz
	jtbl_tar_url := "https://files.pythonhosted.org/packages/9e/7c/b21f3383ca611b56dbc281081cca73a24274c3f39654e7fe196d73a67af6/jtbl-1.6.0.tar.gz"
	jtbl_cmd_tar := exec.Command("curl", "-L", jtbl_tar_url, "-o", "package.tar.gz")
	err := jtbl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jtbl_zip_url := "https://files.pythonhosted.org/packages/9e/7c/b21f3383ca611b56dbc281081cca73a24274c3f39654e7fe196d73a67af6/jtbl-1.6.0.zip"
	jtbl_cmd_zip := exec.Command("curl", "-L", jtbl_zip_url, "-o", "package.zip")
	err = jtbl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jtbl_bin_url := "https://files.pythonhosted.org/packages/9e/7c/b21f3383ca611b56dbc281081cca73a24274c3f39654e7fe196d73a67af6/jtbl-1.6.0.bin"
	jtbl_cmd_bin := exec.Command("curl", "-L", jtbl_bin_url, "-o", "binary.bin")
	err = jtbl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jtbl_src_url := "https://files.pythonhosted.org/packages/9e/7c/b21f3383ca611b56dbc281081cca73a24274c3f39654e7fe196d73a67af6/jtbl-1.6.0.src.tar.gz"
	jtbl_cmd_src := exec.Command("curl", "-L", jtbl_src_url, "-o", "source.tar.gz")
	err = jtbl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jtbl_cmd_direct := exec.Command("./binary")
	err = jtbl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
