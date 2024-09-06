package main

// FleetCli - Manage large fleets of Kubernetes clusters
// Homepage: https://github.com/rancher/fleet

import (
	"fmt"
	
	"os/exec"
)

func installFleetCli() {
	// Método 1: Descargar y extraer .tar.gz
	fleetcli_tar_url := "https://github.com/rancher/fleet/archive/refs/tags/v0.10.1.tar.gz"
	fleetcli_cmd_tar := exec.Command("curl", "-L", fleetcli_tar_url, "-o", "package.tar.gz")
	err := fleetcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fleetcli_zip_url := "https://github.com/rancher/fleet/archive/refs/tags/v0.10.1.zip"
	fleetcli_cmd_zip := exec.Command("curl", "-L", fleetcli_zip_url, "-o", "package.zip")
	err = fleetcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fleetcli_bin_url := "https://github.com/rancher/fleet/archive/refs/tags/v0.10.1.bin"
	fleetcli_cmd_bin := exec.Command("curl", "-L", fleetcli_bin_url, "-o", "binary.bin")
	err = fleetcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fleetcli_src_url := "https://github.com/rancher/fleet/archive/refs/tags/v0.10.1.src.tar.gz"
	fleetcli_cmd_src := exec.Command("curl", "-L", fleetcli_src_url, "-o", "source.tar.gz")
	err = fleetcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fleetcli_cmd_direct := exec.Command("./binary")
	err = fleetcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
