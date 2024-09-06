package main

// Pyinstaller - Bundle a Python application and all its dependencies
// Homepage: https://pyinstaller.org/

import (
	"fmt"
	
	"os/exec"
)

func installPyinstaller() {
	// Método 1: Descargar y extraer .tar.gz
	pyinstaller_tar_url := "https://files.pythonhosted.org/packages/5c/df/30b1f66d35defa37e676556acca4eb775b49637bb73054b0c31af134cd8a/pyinstaller-6.10.0.tar.gz"
	pyinstaller_cmd_tar := exec.Command("curl", "-L", pyinstaller_tar_url, "-o", "package.tar.gz")
	err := pyinstaller_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyinstaller_zip_url := "https://files.pythonhosted.org/packages/5c/df/30b1f66d35defa37e676556acca4eb775b49637bb73054b0c31af134cd8a/pyinstaller-6.10.0.zip"
	pyinstaller_cmd_zip := exec.Command("curl", "-L", pyinstaller_zip_url, "-o", "package.zip")
	err = pyinstaller_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyinstaller_bin_url := "https://files.pythonhosted.org/packages/5c/df/30b1f66d35defa37e676556acca4eb775b49637bb73054b0c31af134cd8a/pyinstaller-6.10.0.bin"
	pyinstaller_cmd_bin := exec.Command("curl", "-L", pyinstaller_bin_url, "-o", "binary.bin")
	err = pyinstaller_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyinstaller_src_url := "https://files.pythonhosted.org/packages/5c/df/30b1f66d35defa37e676556acca4eb775b49637bb73054b0c31af134cd8a/pyinstaller-6.10.0.src.tar.gz"
	pyinstaller_cmd_src := exec.Command("curl", "-L", pyinstaller_src_url, "-o", "source.tar.gz")
	err = pyinstaller_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyinstaller_cmd_direct := exec.Command("./binary")
	err = pyinstaller_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
