package main

// Edgevpn - Immutable, decentralized, statically built p2p VPN
// Homepage: https://mudler.github.io/edgevpn

import (
	"fmt"
	
	"os/exec"
)

func installEdgevpn() {
	// Método 1: Descargar y extraer .tar.gz
	edgevpn_tar_url := "https://github.com/mudler/edgevpn/archive/refs/tags/v0.28.3.tar.gz"
	edgevpn_cmd_tar := exec.Command("curl", "-L", edgevpn_tar_url, "-o", "package.tar.gz")
	err := edgevpn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	edgevpn_zip_url := "https://github.com/mudler/edgevpn/archive/refs/tags/v0.28.3.zip"
	edgevpn_cmd_zip := exec.Command("curl", "-L", edgevpn_zip_url, "-o", "package.zip")
	err = edgevpn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	edgevpn_bin_url := "https://github.com/mudler/edgevpn/archive/refs/tags/v0.28.3.bin"
	edgevpn_cmd_bin := exec.Command("curl", "-L", edgevpn_bin_url, "-o", "binary.bin")
	err = edgevpn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	edgevpn_src_url := "https://github.com/mudler/edgevpn/archive/refs/tags/v0.28.3.src.tar.gz"
	edgevpn_cmd_src := exec.Command("curl", "-L", edgevpn_src_url, "-o", "source.tar.gz")
	err = edgevpn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	edgevpn_cmd_direct := exec.Command("./binary")
	err = edgevpn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
