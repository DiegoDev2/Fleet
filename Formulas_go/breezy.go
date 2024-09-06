package main

// Breezy - Version control system implemented in Python with multi-format support
// Homepage: https://github.com/breezy-team/breezy

import (
	"fmt"
	
	"os/exec"
)

func installBreezy() {
	// Método 1: Descargar y extraer .tar.gz
	breezy_tar_url := "https://files.pythonhosted.org/packages/bb/3f/f1b74d0e32c5455e53655bf095724d37e31b2f184b2dddb899cedbbb6c56/breezy-3.3.8.tar.gz"
	breezy_cmd_tar := exec.Command("curl", "-L", breezy_tar_url, "-o", "package.tar.gz")
	err := breezy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	breezy_zip_url := "https://files.pythonhosted.org/packages/bb/3f/f1b74d0e32c5455e53655bf095724d37e31b2f184b2dddb899cedbbb6c56/breezy-3.3.8.zip"
	breezy_cmd_zip := exec.Command("curl", "-L", breezy_zip_url, "-o", "package.zip")
	err = breezy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	breezy_bin_url := "https://files.pythonhosted.org/packages/bb/3f/f1b74d0e32c5455e53655bf095724d37e31b2f184b2dddb899cedbbb6c56/breezy-3.3.8.bin"
	breezy_cmd_bin := exec.Command("curl", "-L", breezy_bin_url, "-o", "binary.bin")
	err = breezy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	breezy_src_url := "https://files.pythonhosted.org/packages/bb/3f/f1b74d0e32c5455e53655bf095724d37e31b2f184b2dddb899cedbbb6c56/breezy-3.3.8.src.tar.gz"
	breezy_cmd_src := exec.Command("curl", "-L", breezy_src_url, "-o", "source.tar.gz")
	err = breezy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	breezy_cmd_direct := exec.Command("./binary")
	err = breezy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
