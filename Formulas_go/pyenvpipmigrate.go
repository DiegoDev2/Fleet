package main

// PyenvPipMigrate - Migrate pip packages from one Python version to another
// Homepage: https://github.com/pyenv/pyenv-pip-migrate

import (
	"fmt"
	
	"os/exec"
)

func installPyenvPipMigrate() {
	// Método 1: Descargar y extraer .tar.gz
	pyenvpipmigrate_tar_url := "https://github.com/pyenv/pyenv-pip-migrate/archive/refs/tags/v20181205.tar.gz"
	pyenvpipmigrate_cmd_tar := exec.Command("curl", "-L", pyenvpipmigrate_tar_url, "-o", "package.tar.gz")
	err := pyenvpipmigrate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyenvpipmigrate_zip_url := "https://github.com/pyenv/pyenv-pip-migrate/archive/refs/tags/v20181205.zip"
	pyenvpipmigrate_cmd_zip := exec.Command("curl", "-L", pyenvpipmigrate_zip_url, "-o", "package.zip")
	err = pyenvpipmigrate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyenvpipmigrate_bin_url := "https://github.com/pyenv/pyenv-pip-migrate/archive/refs/tags/v20181205.bin"
	pyenvpipmigrate_cmd_bin := exec.Command("curl", "-L", pyenvpipmigrate_bin_url, "-o", "binary.bin")
	err = pyenvpipmigrate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyenvpipmigrate_src_url := "https://github.com/pyenv/pyenv-pip-migrate/archive/refs/tags/v20181205.src.tar.gz"
	pyenvpipmigrate_cmd_src := exec.Command("curl", "-L", pyenvpipmigrate_src_url, "-o", "source.tar.gz")
	err = pyenvpipmigrate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyenvpipmigrate_cmd_direct := exec.Command("./binary")
	err = pyenvpipmigrate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyenv")
exec.Command("latte", "install", "pyenv")
}
