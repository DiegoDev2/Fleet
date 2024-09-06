package main

// Ipython - Interactive computing in Python
// Homepage: https://ipython.org/

import (
	"fmt"
	
	"os/exec"
)

func installIpython() {
	// Método 1: Descargar y extraer .tar.gz
	ipython_tar_url := "https://files.pythonhosted.org/packages/57/24/d4fabaca03c8804bf0b8d994c8ae3a20e57e9330d277fb43d83e558dec5e/ipython-8.27.0.tar.gz"
	ipython_cmd_tar := exec.Command("curl", "-L", ipython_tar_url, "-o", "package.tar.gz")
	err := ipython_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipython_zip_url := "https://files.pythonhosted.org/packages/57/24/d4fabaca03c8804bf0b8d994c8ae3a20e57e9330d277fb43d83e558dec5e/ipython-8.27.0.zip"
	ipython_cmd_zip := exec.Command("curl", "-L", ipython_zip_url, "-o", "package.zip")
	err = ipython_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipython_bin_url := "https://files.pythonhosted.org/packages/57/24/d4fabaca03c8804bf0b8d994c8ae3a20e57e9330d277fb43d83e558dec5e/ipython-8.27.0.bin"
	ipython_cmd_bin := exec.Command("curl", "-L", ipython_bin_url, "-o", "binary.bin")
	err = ipython_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipython_src_url := "https://files.pythonhosted.org/packages/57/24/d4fabaca03c8804bf0b8d994c8ae3a20e57e9330d277fb43d83e558dec5e/ipython-8.27.0.src.tar.gz"
	ipython_cmd_src := exec.Command("curl", "-L", ipython_src_url, "-o", "source.tar.gz")
	err = ipython_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipython_cmd_direct := exec.Command("./binary")
	err = ipython_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
