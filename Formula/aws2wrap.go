package main

// Aws2Wrap - Script to export current AWS SSO credentials or run a sub-process with them
// Homepage: https://github.com/linaro-its/aws2-wrap

import (
	"fmt"
	
	"os/exec"
)

func installAws2Wrap() {
	// Método 1: Descargar y extraer .tar.gz
	aws2wrap_tar_url := "https://files.pythonhosted.org/packages/6d/c7/8afdf4d0c7c6e2072c73a0150f9789445af33381a611d33333f4c9bf1ef6/aws2-wrap-1.4.0.tar.gz"
	aws2wrap_cmd_tar := exec.Command("curl", "-L", aws2wrap_tar_url, "-o", "package.tar.gz")
	err := aws2wrap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aws2wrap_zip_url := "https://files.pythonhosted.org/packages/6d/c7/8afdf4d0c7c6e2072c73a0150f9789445af33381a611d33333f4c9bf1ef6/aws2-wrap-1.4.0.zip"
	aws2wrap_cmd_zip := exec.Command("curl", "-L", aws2wrap_zip_url, "-o", "package.zip")
	err = aws2wrap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aws2wrap_bin_url := "https://files.pythonhosted.org/packages/6d/c7/8afdf4d0c7c6e2072c73a0150f9789445af33381a611d33333f4c9bf1ef6/aws2-wrap-1.4.0.bin"
	aws2wrap_cmd_bin := exec.Command("curl", "-L", aws2wrap_bin_url, "-o", "binary.bin")
	err = aws2wrap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aws2wrap_src_url := "https://files.pythonhosted.org/packages/6d/c7/8afdf4d0c7c6e2072c73a0150f9789445af33381a611d33333f4c9bf1ef6/aws2-wrap-1.4.0.src.tar.gz"
	aws2wrap_cmd_src := exec.Command("curl", "-L", aws2wrap_src_url, "-o", "source.tar.gz")
	err = aws2wrap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aws2wrap_cmd_direct := exec.Command("./binary")
	err = aws2wrap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
