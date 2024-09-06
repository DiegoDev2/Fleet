package main

// Grayskull - Recipe generator for Conda
// Homepage: https://github.com/conda/grayskull

import (
	"fmt"
	
	"os/exec"
)

func installGrayskull() {
	// Método 1: Descargar y extraer .tar.gz
	grayskull_tar_url := "https://files.pythonhosted.org/packages/01/ed/8262d7838937c69e1f67c24a74787ffca5be0ae15fb40c3d79852345eea3/grayskull-2.7.1.tar.gz"
	grayskull_cmd_tar := exec.Command("curl", "-L", grayskull_tar_url, "-o", "package.tar.gz")
	err := grayskull_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grayskull_zip_url := "https://files.pythonhosted.org/packages/01/ed/8262d7838937c69e1f67c24a74787ffca5be0ae15fb40c3d79852345eea3/grayskull-2.7.1.zip"
	grayskull_cmd_zip := exec.Command("curl", "-L", grayskull_zip_url, "-o", "package.zip")
	err = grayskull_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grayskull_bin_url := "https://files.pythonhosted.org/packages/01/ed/8262d7838937c69e1f67c24a74787ffca5be0ae15fb40c3d79852345eea3/grayskull-2.7.1.bin"
	grayskull_cmd_bin := exec.Command("curl", "-L", grayskull_bin_url, "-o", "binary.bin")
	err = grayskull_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grayskull_src_url := "https://files.pythonhosted.org/packages/01/ed/8262d7838937c69e1f67c24a74787ffca5be0ae15fb40c3d79852345eea3/grayskull-2.7.1.src.tar.gz"
	grayskull_cmd_src := exec.Command("curl", "-L", grayskull_src_url, "-o", "source.tar.gz")
	err = grayskull_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grayskull_cmd_direct := exec.Command("./binary")
	err = grayskull_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
