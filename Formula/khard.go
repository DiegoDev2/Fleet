package main

// Khard - Console carddav client
// Homepage: https://github.com/lucc/khard

import (
	"fmt"
	
	"os/exec"
)

func installKhard() {
	// Método 1: Descargar y extraer .tar.gz
	khard_tar_url := "https://files.pythonhosted.org/packages/0d/00/215a69d2ae96cac511a6594116958bf13e210dd24f78c48f5ffaf039edec/khard-0.19.1.tar.gz"
	khard_cmd_tar := exec.Command("curl", "-L", khard_tar_url, "-o", "package.tar.gz")
	err := khard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	khard_zip_url := "https://files.pythonhosted.org/packages/0d/00/215a69d2ae96cac511a6594116958bf13e210dd24f78c48f5ffaf039edec/khard-0.19.1.zip"
	khard_cmd_zip := exec.Command("curl", "-L", khard_zip_url, "-o", "package.zip")
	err = khard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	khard_bin_url := "https://files.pythonhosted.org/packages/0d/00/215a69d2ae96cac511a6594116958bf13e210dd24f78c48f5ffaf039edec/khard-0.19.1.bin"
	khard_cmd_bin := exec.Command("curl", "-L", khard_bin_url, "-o", "binary.bin")
	err = khard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	khard_src_url := "https://files.pythonhosted.org/packages/0d/00/215a69d2ae96cac511a6594116958bf13e210dd24f78c48f5ffaf039edec/khard-0.19.1.src.tar.gz"
	khard_cmd_src := exec.Command("curl", "-L", khard_src_url, "-o", "source.tar.gz")
	err = khard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	khard_cmd_direct := exec.Command("./binary")
	err = khard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
