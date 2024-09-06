package main

// Dynaconf - Configuration Management for Python
// Homepage: https://www.dynaconf.com/

import (
	"fmt"
	
	"os/exec"
)

func installDynaconf() {
	// Método 1: Descargar y extraer .tar.gz
	dynaconf_tar_url := "https://files.pythonhosted.org/packages/56/1a/324f1bf234cc4f98445305fd8719245318466e310e05caea7ef052748ecd/dynaconf-3.2.6.tar.gz"
	dynaconf_cmd_tar := exec.Command("curl", "-L", dynaconf_tar_url, "-o", "package.tar.gz")
	err := dynaconf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dynaconf_zip_url := "https://files.pythonhosted.org/packages/56/1a/324f1bf234cc4f98445305fd8719245318466e310e05caea7ef052748ecd/dynaconf-3.2.6.zip"
	dynaconf_cmd_zip := exec.Command("curl", "-L", dynaconf_zip_url, "-o", "package.zip")
	err = dynaconf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dynaconf_bin_url := "https://files.pythonhosted.org/packages/56/1a/324f1bf234cc4f98445305fd8719245318466e310e05caea7ef052748ecd/dynaconf-3.2.6.bin"
	dynaconf_cmd_bin := exec.Command("curl", "-L", dynaconf_bin_url, "-o", "binary.bin")
	err = dynaconf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dynaconf_src_url := "https://files.pythonhosted.org/packages/56/1a/324f1bf234cc4f98445305fd8719245318466e310e05caea7ef052748ecd/dynaconf-3.2.6.src.tar.gz"
	dynaconf_cmd_src := exec.Command("curl", "-L", dynaconf_src_url, "-o", "source.tar.gz")
	err = dynaconf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dynaconf_cmd_direct := exec.Command("./binary")
	err = dynaconf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
