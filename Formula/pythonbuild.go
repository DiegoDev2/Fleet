package main

// PythonBuild - Simple, correct PEP 517 build frontend
// Homepage: https://github.com/pypa/build

import (
	"fmt"
	
	"os/exec"
)

func installPythonBuild() {
	// Método 1: Descargar y extraer .tar.gz
	pythonbuild_tar_url := "https://files.pythonhosted.org/packages/ce/9e/2d725d2f7729c6e79ca62aeb926492abbc06e25910dd30139d60a68bcb19/build-1.2.1.tar.gz"
	pythonbuild_cmd_tar := exec.Command("curl", "-L", pythonbuild_tar_url, "-o", "package.tar.gz")
	err := pythonbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonbuild_zip_url := "https://files.pythonhosted.org/packages/ce/9e/2d725d2f7729c6e79ca62aeb926492abbc06e25910dd30139d60a68bcb19/build-1.2.1.zip"
	pythonbuild_cmd_zip := exec.Command("curl", "-L", pythonbuild_zip_url, "-o", "package.zip")
	err = pythonbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonbuild_bin_url := "https://files.pythonhosted.org/packages/ce/9e/2d725d2f7729c6e79ca62aeb926492abbc06e25910dd30139d60a68bcb19/build-1.2.1.bin"
	pythonbuild_cmd_bin := exec.Command("curl", "-L", pythonbuild_bin_url, "-o", "binary.bin")
	err = pythonbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonbuild_src_url := "https://files.pythonhosted.org/packages/ce/9e/2d725d2f7729c6e79ca62aeb926492abbc06e25910dd30139d60a68bcb19/build-1.2.1.src.tar.gz"
	pythonbuild_cmd_src := exec.Command("curl", "-L", pythonbuild_src_url, "-o", "source.tar.gz")
	err = pythonbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonbuild_cmd_direct := exec.Command("./binary")
	err = pythonbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
