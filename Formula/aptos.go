package main

// Aptos - Layer 1 blockchain built to support fair access to decentralized assets for all
// Homepage: https://aptosfoundation.org/

import (
	"fmt"
	
	"os/exec"
)

func installAptos() {
	// Método 1: Descargar y extraer .tar.gz
	aptos_tar_url := "https://github.com/aptos-labs/aptos-core/archive/refs/tags/aptos-cli-v4.1.0.tar.gz"
	aptos_cmd_tar := exec.Command("curl", "-L", aptos_tar_url, "-o", "package.tar.gz")
	err := aptos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aptos_zip_url := "https://github.com/aptos-labs/aptos-core/archive/refs/tags/aptos-cli-v4.1.0.zip"
	aptos_cmd_zip := exec.Command("curl", "-L", aptos_zip_url, "-o", "package.zip")
	err = aptos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aptos_bin_url := "https://github.com/aptos-labs/aptos-core/archive/refs/tags/aptos-cli-v4.1.0.bin"
	aptos_cmd_bin := exec.Command("curl", "-L", aptos_bin_url, "-o", "binary.bin")
	err = aptos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aptos_src_url := "https://github.com/aptos-labs/aptos-core/archive/refs/tags/aptos-cli-v4.1.0.src.tar.gz"
	aptos_cmd_src := exec.Command("curl", "-L", aptos_src_url, "-o", "source.tar.gz")
	err = aptos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aptos_cmd_direct := exec.Command("./binary")
	err = aptos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustfmt")
	exec.Command("latte", "install", "rustfmt").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: zip")
	exec.Command("latte", "install", "zip").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
