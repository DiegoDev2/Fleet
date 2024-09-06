package main

// Scons - Substitute for classic 'make' tool with autoconf/automake functionality
// Homepage: https://www.scons.org/

import (
	"fmt"
	
	"os/exec"
)

func installScons() {
	// Método 1: Descargar y extraer .tar.gz
	scons_tar_url := "https://files.pythonhosted.org/packages/b9/76/a2c1293642f9a448f2d012cabf525be69ca5abf4af289bc0935ac1554ee8/scons-4.8.1.tar.gz"
	scons_cmd_tar := exec.Command("curl", "-L", scons_tar_url, "-o", "package.tar.gz")
	err := scons_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scons_zip_url := "https://files.pythonhosted.org/packages/b9/76/a2c1293642f9a448f2d012cabf525be69ca5abf4af289bc0935ac1554ee8/scons-4.8.1.zip"
	scons_cmd_zip := exec.Command("curl", "-L", scons_zip_url, "-o", "package.zip")
	err = scons_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scons_bin_url := "https://files.pythonhosted.org/packages/b9/76/a2c1293642f9a448f2d012cabf525be69ca5abf4af289bc0935ac1554ee8/scons-4.8.1.bin"
	scons_cmd_bin := exec.Command("curl", "-L", scons_bin_url, "-o", "binary.bin")
	err = scons_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scons_src_url := "https://files.pythonhosted.org/packages/b9/76/a2c1293642f9a448f2d012cabf525be69ca5abf4af289bc0935ac1554ee8/scons-4.8.1.src.tar.gz"
	scons_cmd_src := exec.Command("curl", "-L", scons_src_url, "-o", "source.tar.gz")
	err = scons_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scons_cmd_direct := exec.Command("./binary")
	err = scons_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
