package main

// GrinWallet - Official wallet for the cryptocurrency Grin
// Homepage: https://grin.mw

import (
	"fmt"
	
	"os/exec"
)

func installGrinWallet() {
	// Método 1: Descargar y extraer .tar.gz
	grinwallet_tar_url := "https://github.com/mimblewimble/grin-wallet/archive/refs/tags/v5.3.1.tar.gz"
	grinwallet_cmd_tar := exec.Command("curl", "-L", grinwallet_tar_url, "-o", "package.tar.gz")
	err := grinwallet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grinwallet_zip_url := "https://github.com/mimblewimble/grin-wallet/archive/refs/tags/v5.3.1.zip"
	grinwallet_cmd_zip := exec.Command("curl", "-L", grinwallet_zip_url, "-o", "package.zip")
	err = grinwallet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grinwallet_bin_url := "https://github.com/mimblewimble/grin-wallet/archive/refs/tags/v5.3.1.bin"
	grinwallet_cmd_bin := exec.Command("curl", "-L", grinwallet_bin_url, "-o", "binary.bin")
	err = grinwallet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grinwallet_src_url := "https://github.com/mimblewimble/grin-wallet/archive/refs/tags/v5.3.1.src.tar.gz"
	grinwallet_cmd_src := exec.Command("curl", "-L", grinwallet_src_url, "-o", "source.tar.gz")
	err = grinwallet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grinwallet_cmd_direct := exec.Command("./binary")
	err = grinwallet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
