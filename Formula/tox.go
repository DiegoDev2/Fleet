package main

// Tox - Generic Python virtualenv management and test command-line tool
// Homepage: https://tox.wiki/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installTox() {
	// Método 1: Descargar y extraer .tar.gz
	tox_tar_url := "https://files.pythonhosted.org/packages/38/d6/e3713386deb7df92ef383693544eccb18260ce4325fabcf5aac6594c1d95/tox-4.18.0.tar.gz"
	tox_cmd_tar := exec.Command("curl", "-L", tox_tar_url, "-o", "package.tar.gz")
	err := tox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tox_zip_url := "https://files.pythonhosted.org/packages/38/d6/e3713386deb7df92ef383693544eccb18260ce4325fabcf5aac6594c1d95/tox-4.18.0.zip"
	tox_cmd_zip := exec.Command("curl", "-L", tox_zip_url, "-o", "package.zip")
	err = tox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tox_bin_url := "https://files.pythonhosted.org/packages/38/d6/e3713386deb7df92ef383693544eccb18260ce4325fabcf5aac6594c1d95/tox-4.18.0.bin"
	tox_cmd_bin := exec.Command("curl", "-L", tox_bin_url, "-o", "binary.bin")
	err = tox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tox_src_url := "https://files.pythonhosted.org/packages/38/d6/e3713386deb7df92ef383693544eccb18260ce4325fabcf5aac6594c1d95/tox-4.18.0.src.tar.gz"
	tox_cmd_src := exec.Command("curl", "-L", tox_src_url, "-o", "source.tar.gz")
	err = tox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tox_cmd_direct := exec.Command("./binary")
	err = tox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
