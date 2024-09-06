package main

// Numpy - Package for scientific computing with Python
// Homepage: https://www.numpy.org/

import (
	"fmt"
	
	"os/exec"
)

func installNumpy() {
	// Método 1: Descargar y extraer .tar.gz
	numpy_tar_url := "https://files.pythonhosted.org/packages/59/5f/9003bb3e632f2b58f5e3a3378902dcc73c5518070736c6740fe52454e8e1/numpy-2.1.1.tar.gz"
	numpy_cmd_tar := exec.Command("curl", "-L", numpy_tar_url, "-o", "package.tar.gz")
	err := numpy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	numpy_zip_url := "https://files.pythonhosted.org/packages/59/5f/9003bb3e632f2b58f5e3a3378902dcc73c5518070736c6740fe52454e8e1/numpy-2.1.1.zip"
	numpy_cmd_zip := exec.Command("curl", "-L", numpy_zip_url, "-o", "package.zip")
	err = numpy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	numpy_bin_url := "https://files.pythonhosted.org/packages/59/5f/9003bb3e632f2b58f5e3a3378902dcc73c5518070736c6740fe52454e8e1/numpy-2.1.1.bin"
	numpy_cmd_bin := exec.Command("curl", "-L", numpy_bin_url, "-o", "binary.bin")
	err = numpy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	numpy_src_url := "https://files.pythonhosted.org/packages/59/5f/9003bb3e632f2b58f5e3a3378902dcc73c5518070736c6740fe52454e8e1/numpy-2.1.1.src.tar.gz"
	numpy_cmd_src := exec.Command("curl", "-L", numpy_src_url, "-o", "source.tar.gz")
	err = numpy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	numpy_cmd_direct := exec.Command("./binary")
	err = numpy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
