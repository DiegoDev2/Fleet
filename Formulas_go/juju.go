package main

// Juju - DevOps management tool
// Homepage: https://juju.is/

import (
	"fmt"
	
	"os/exec"
)

func installJuju() {
	// Método 1: Descargar y extraer .tar.gz
	juju_tar_url := "https://launchpad.net/juju/3.5/3.5.3/+download/juju-core_3.5.3.tar.gz"
	juju_cmd_tar := exec.Command("curl", "-L", juju_tar_url, "-o", "package.tar.gz")
	err := juju_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	juju_zip_url := "https://launchpad.net/juju/3.5/3.5.3/+download/juju-core_3.5.3.zip"
	juju_cmd_zip := exec.Command("curl", "-L", juju_zip_url, "-o", "package.zip")
	err = juju_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	juju_bin_url := "https://launchpad.net/juju/3.5/3.5.3/+download/juju-core_3.5.3.bin"
	juju_cmd_bin := exec.Command("curl", "-L", juju_bin_url, "-o", "binary.bin")
	err = juju_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	juju_src_url := "https://launchpad.net/juju/3.5/3.5.3/+download/juju-core_3.5.3.src.tar.gz"
	juju_cmd_src := exec.Command("curl", "-L", juju_src_url, "-o", "source.tar.gz")
	err = juju_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	juju_cmd_direct := exec.Command("./binary")
	err = juju_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
