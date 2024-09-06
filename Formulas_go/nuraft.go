package main

// Nuraft - C++ implementation of Raft core logic as a replication library
// Homepage: https://github.com/eBay/NuRaft

import (
	"fmt"
	
	"os/exec"
)

func installNuraft() {
	// Método 1: Descargar y extraer .tar.gz
	nuraft_tar_url := "https://github.com/eBay/NuRaft/archive/refs/tags/v2.1.0.tar.gz"
	nuraft_cmd_tar := exec.Command("curl", "-L", nuraft_tar_url, "-o", "package.tar.gz")
	err := nuraft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nuraft_zip_url := "https://github.com/eBay/NuRaft/archive/refs/tags/v2.1.0.zip"
	nuraft_cmd_zip := exec.Command("curl", "-L", nuraft_zip_url, "-o", "package.zip")
	err = nuraft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nuraft_bin_url := "https://github.com/eBay/NuRaft/archive/refs/tags/v2.1.0.bin"
	nuraft_cmd_bin := exec.Command("curl", "-L", nuraft_bin_url, "-o", "binary.bin")
	err = nuraft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nuraft_src_url := "https://github.com/eBay/NuRaft/archive/refs/tags/v2.1.0.src.tar.gz"
	nuraft_cmd_src := exec.Command("curl", "-L", nuraft_src_url, "-o", "source.tar.gz")
	err = nuraft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nuraft_cmd_direct := exec.Command("./binary")
	err = nuraft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: asio")
exec.Command("latte", "install", "asio")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
