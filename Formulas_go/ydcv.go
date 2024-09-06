package main

// Ydcv - YouDao Console Version
// Homepage: https://github.com/felixonmars/ydcv

import (
	"fmt"
	
	"os/exec"
)

func installYdcv() {
	// Método 1: Descargar y extraer .tar.gz
	ydcv_tar_url := "https://files.pythonhosted.org/packages/1f/29/17124ebfdea8d810774977474a8652018c04c4a6db1ca413189f7e5b9d52/ydcv-0.7.tar.gz"
	ydcv_cmd_tar := exec.Command("curl", "-L", ydcv_tar_url, "-o", "package.tar.gz")
	err := ydcv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ydcv_zip_url := "https://files.pythonhosted.org/packages/1f/29/17124ebfdea8d810774977474a8652018c04c4a6db1ca413189f7e5b9d52/ydcv-0.7.zip"
	ydcv_cmd_zip := exec.Command("curl", "-L", ydcv_zip_url, "-o", "package.zip")
	err = ydcv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ydcv_bin_url := "https://files.pythonhosted.org/packages/1f/29/17124ebfdea8d810774977474a8652018c04c4a6db1ca413189f7e5b9d52/ydcv-0.7.bin"
	ydcv_cmd_bin := exec.Command("curl", "-L", ydcv_bin_url, "-o", "binary.bin")
	err = ydcv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ydcv_src_url := "https://files.pythonhosted.org/packages/1f/29/17124ebfdea8d810774977474a8652018c04c4a6db1ca413189f7e5b9d52/ydcv-0.7.src.tar.gz"
	ydcv_cmd_src := exec.Command("curl", "-L", ydcv_src_url, "-o", "source.tar.gz")
	err = ydcv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ydcv_cmd_direct := exec.Command("./binary")
	err = ydcv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
}
