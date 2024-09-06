package main

// Xxh - Bring your favorite shell wherever you go through the ssh
// Homepage: https://github.com/xxh/xxh

import (
	"fmt"
	
	"os/exec"
)

func installXxh() {
	// Método 1: Descargar y extraer .tar.gz
	xxh_tar_url := "https://files.pythonhosted.org/packages/d6/ac/fb40368ff37fbdd00d041e241cc0d7a50cdac7bc6ae54dcb9f1349acdde6/xxh-xxh-0.8.14.tar.gz"
	xxh_cmd_tar := exec.Command("curl", "-L", xxh_tar_url, "-o", "package.tar.gz")
	err := xxh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xxh_zip_url := "https://files.pythonhosted.org/packages/d6/ac/fb40368ff37fbdd00d041e241cc0d7a50cdac7bc6ae54dcb9f1349acdde6/xxh-xxh-0.8.14.zip"
	xxh_cmd_zip := exec.Command("curl", "-L", xxh_zip_url, "-o", "package.zip")
	err = xxh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xxh_bin_url := "https://files.pythonhosted.org/packages/d6/ac/fb40368ff37fbdd00d041e241cc0d7a50cdac7bc6ae54dcb9f1349acdde6/xxh-xxh-0.8.14.bin"
	xxh_cmd_bin := exec.Command("curl", "-L", xxh_bin_url, "-o", "binary.bin")
	err = xxh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xxh_src_url := "https://files.pythonhosted.org/packages/d6/ac/fb40368ff37fbdd00d041e241cc0d7a50cdac7bc6ae54dcb9f1349acdde6/xxh-xxh-0.8.14.src.tar.gz"
	xxh_cmd_src := exec.Command("curl", "-L", xxh_src_url, "-o", "source.tar.gz")
	err = xxh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xxh_cmd_direct := exec.Command("./binary")
	err = xxh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
