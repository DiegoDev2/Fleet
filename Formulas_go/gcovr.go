package main

// Gcovr - Reports from gcov test coverage program
// Homepage: https://gcovr.com/

import (
	"fmt"
	
	"os/exec"
)

func installGcovr() {
	// Método 1: Descargar y extraer .tar.gz
	gcovr_tar_url := "https://files.pythonhosted.org/packages/ed/9b/119d9b9501a9d0bc91be6b163be98125a9345e37871f4f3243b112d456e6/gcovr-7.2.tar.gz"
	gcovr_cmd_tar := exec.Command("curl", "-L", gcovr_tar_url, "-o", "package.tar.gz")
	err := gcovr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gcovr_zip_url := "https://files.pythonhosted.org/packages/ed/9b/119d9b9501a9d0bc91be6b163be98125a9345e37871f4f3243b112d456e6/gcovr-7.2.zip"
	gcovr_cmd_zip := exec.Command("curl", "-L", gcovr_zip_url, "-o", "package.zip")
	err = gcovr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gcovr_bin_url := "https://files.pythonhosted.org/packages/ed/9b/119d9b9501a9d0bc91be6b163be98125a9345e37871f4f3243b112d456e6/gcovr-7.2.bin"
	gcovr_cmd_bin := exec.Command("curl", "-L", gcovr_bin_url, "-o", "binary.bin")
	err = gcovr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gcovr_src_url := "https://files.pythonhosted.org/packages/ed/9b/119d9b9501a9d0bc91be6b163be98125a9345e37871f4f3243b112d456e6/gcovr-7.2.src.tar.gz"
	gcovr_cmd_src := exec.Command("curl", "-L", gcovr_src_url, "-o", "source.tar.gz")
	err = gcovr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gcovr_cmd_direct := exec.Command("./binary")
	err = gcovr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
