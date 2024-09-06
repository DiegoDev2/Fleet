package main

// Charm - Tool for managing Juju Charms
// Homepage: https://github.com/juju/charmstore-client

import (
	"fmt"
	
	"os/exec"
)

func installCharm() {
	// Método 1: Descargar y extraer .tar.gz
	charm_tar_url := "https://github.com/juju/charmstore-client/archive/refs/tags/v2.5.2.tar.gz"
	charm_cmd_tar := exec.Command("curl", "-L", charm_tar_url, "-o", "package.tar.gz")
	err := charm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	charm_zip_url := "https://github.com/juju/charmstore-client/archive/refs/tags/v2.5.2.zip"
	charm_cmd_zip := exec.Command("curl", "-L", charm_zip_url, "-o", "package.zip")
	err = charm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	charm_bin_url := "https://github.com/juju/charmstore-client/archive/refs/tags/v2.5.2.bin"
	charm_cmd_bin := exec.Command("curl", "-L", charm_bin_url, "-o", "binary.bin")
	err = charm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	charm_src_url := "https://github.com/juju/charmstore-client/archive/refs/tags/v2.5.2.src.tar.gz"
	charm_cmd_src := exec.Command("curl", "-L", charm_src_url, "-o", "source.tar.gz")
	err = charm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	charm_cmd_direct := exec.Command("./binary")
	err = charm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
