package main

// Pythran - Ahead of Time compiler for numeric kernels
// Homepage: https://pythran.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installPythran() {
	// Método 1: Descargar y extraer .tar.gz
	pythran_tar_url := "https://files.pythonhosted.org/packages/73/32/f892675c5009cd4c1895ded3d6153476bf00adb5ad1634d03635620881f5/pythran-0.16.1.tar.gz"
	pythran_cmd_tar := exec.Command("curl", "-L", pythran_tar_url, "-o", "package.tar.gz")
	err := pythran_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythran_zip_url := "https://files.pythonhosted.org/packages/73/32/f892675c5009cd4c1895ded3d6153476bf00adb5ad1634d03635620881f5/pythran-0.16.1.zip"
	pythran_cmd_zip := exec.Command("curl", "-L", pythran_zip_url, "-o", "package.zip")
	err = pythran_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythran_bin_url := "https://files.pythonhosted.org/packages/73/32/f892675c5009cd4c1895ded3d6153476bf00adb5ad1634d03635620881f5/pythran-0.16.1.bin"
	pythran_cmd_bin := exec.Command("curl", "-L", pythran_bin_url, "-o", "binary.bin")
	err = pythran_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythran_src_url := "https://files.pythonhosted.org/packages/73/32/f892675c5009cd4c1895ded3d6153476bf00adb5ad1634d03635620881f5/pythran-0.16.1.src.tar.gz"
	pythran_cmd_src := exec.Command("curl", "-L", pythran_src_url, "-o", "source.tar.gz")
	err = pythran_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythran_cmd_direct := exec.Command("./binary")
	err = pythran_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
