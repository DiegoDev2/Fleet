package main

// Rpmspectool - Utility for handling RPM spec files
// Homepage: https://github.com/nphilipp/rpmspectool

import (
	"fmt"
	
	"os/exec"
)

func installRpmspectool() {
	// Método 1: Descargar y extraer .tar.gz
	rpmspectool_tar_url := "https://files.pythonhosted.org/packages/7d/cc/53ef9a699df75f3f29f672d0bdf7aae162829e2c98f7b7b5f063fd5d3a46/rpmspectool-1.99.10.tar.gz"
	rpmspectool_cmd_tar := exec.Command("curl", "-L", rpmspectool_tar_url, "-o", "package.tar.gz")
	err := rpmspectool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rpmspectool_zip_url := "https://files.pythonhosted.org/packages/7d/cc/53ef9a699df75f3f29f672d0bdf7aae162829e2c98f7b7b5f063fd5d3a46/rpmspectool-1.99.10.zip"
	rpmspectool_cmd_zip := exec.Command("curl", "-L", rpmspectool_zip_url, "-o", "package.zip")
	err = rpmspectool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rpmspectool_bin_url := "https://files.pythonhosted.org/packages/7d/cc/53ef9a699df75f3f29f672d0bdf7aae162829e2c98f7b7b5f063fd5d3a46/rpmspectool-1.99.10.bin"
	rpmspectool_cmd_bin := exec.Command("curl", "-L", rpmspectool_bin_url, "-o", "binary.bin")
	err = rpmspectool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rpmspectool_src_url := "https://files.pythonhosted.org/packages/7d/cc/53ef9a699df75f3f29f672d0bdf7aae162829e2c98f7b7b5f063fd5d3a46/rpmspectool-1.99.10.src.tar.gz"
	rpmspectool_cmd_src := exec.Command("curl", "-L", rpmspectool_src_url, "-o", "source.tar.gz")
	err = rpmspectool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rpmspectool_cmd_direct := exec.Command("./binary")
	err = rpmspectool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: rpm")
exec.Command("latte", "install", "rpm")
}
