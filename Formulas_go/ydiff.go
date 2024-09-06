package main

// Ydiff - View colored diff with side by side and auto pager support
// Homepage: https://github.com/ymattw/ydiff

import (
	"fmt"
	
	"os/exec"
)

func installYdiff() {
	// Método 1: Descargar y extraer .tar.gz
	ydiff_tar_url := "https://files.pythonhosted.org/packages/22/e1/df78c47d98070228bc1589df6aa2b2c7ae1d4b35f9312dbb2d7a122a0f19/ydiff-1.3.tar.gz"
	ydiff_cmd_tar := exec.Command("curl", "-L", ydiff_tar_url, "-o", "package.tar.gz")
	err := ydiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ydiff_zip_url := "https://files.pythonhosted.org/packages/22/e1/df78c47d98070228bc1589df6aa2b2c7ae1d4b35f9312dbb2d7a122a0f19/ydiff-1.3.zip"
	ydiff_cmd_zip := exec.Command("curl", "-L", ydiff_zip_url, "-o", "package.zip")
	err = ydiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ydiff_bin_url := "https://files.pythonhosted.org/packages/22/e1/df78c47d98070228bc1589df6aa2b2c7ae1d4b35f9312dbb2d7a122a0f19/ydiff-1.3.bin"
	ydiff_cmd_bin := exec.Command("curl", "-L", ydiff_bin_url, "-o", "binary.bin")
	err = ydiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ydiff_src_url := "https://files.pythonhosted.org/packages/22/e1/df78c47d98070228bc1589df6aa2b2c7ae1d4b35f9312dbb2d7a122a0f19/ydiff-1.3.src.tar.gz"
	ydiff_cmd_src := exec.Command("curl", "-L", ydiff_src_url, "-o", "source.tar.gz")
	err = ydiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ydiff_cmd_direct := exec.Command("./binary")
	err = ydiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
