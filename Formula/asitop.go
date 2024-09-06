package main

// Asitop - Perf monitoring CLI tool for Apple Silicon
// Homepage: https://tlkh.github.io/asitop/

import (
	"fmt"
	
	"os/exec"
)

func installAsitop() {
	// Método 1: Descargar y extraer .tar.gz
	asitop_tar_url := "https://files.pythonhosted.org/packages/93/bc/8755d818efc33dd758322086a23f08bee5e1f7769c339a8be5c142adbbbc/asitop-0.0.24.tar.gz"
	asitop_cmd_tar := exec.Command("curl", "-L", asitop_tar_url, "-o", "package.tar.gz")
	err := asitop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asitop_zip_url := "https://files.pythonhosted.org/packages/93/bc/8755d818efc33dd758322086a23f08bee5e1f7769c339a8be5c142adbbbc/asitop-0.0.24.zip"
	asitop_cmd_zip := exec.Command("curl", "-L", asitop_zip_url, "-o", "package.zip")
	err = asitop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asitop_bin_url := "https://files.pythonhosted.org/packages/93/bc/8755d818efc33dd758322086a23f08bee5e1f7769c339a8be5c142adbbbc/asitop-0.0.24.bin"
	asitop_cmd_bin := exec.Command("curl", "-L", asitop_bin_url, "-o", "binary.bin")
	err = asitop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asitop_src_url := "https://files.pythonhosted.org/packages/93/bc/8755d818efc33dd758322086a23f08bee5e1f7769c339a8be5c142adbbbc/asitop-0.0.24.src.tar.gz"
	asitop_cmd_src := exec.Command("curl", "-L", asitop_src_url, "-o", "source.tar.gz")
	err = asitop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asitop_cmd_direct := exec.Command("./binary")
	err = asitop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
