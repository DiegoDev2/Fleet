package main

// Coinutils - COIN-OR utilities
// Homepage: https://github.com/coin-or/CoinUtils

import (
	"fmt"
	
	"os/exec"
)

func installCoinutils() {
	// Método 1: Descargar y extraer .tar.gz
	coinutils_tar_url := "https://github.com/coin-or/CoinUtils/archive/refs/tags/releases/2.11.12.tar.gz"
	coinutils_cmd_tar := exec.Command("curl", "-L", coinutils_tar_url, "-o", "package.tar.gz")
	err := coinutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coinutils_zip_url := "https://github.com/coin-or/CoinUtils/archive/refs/tags/releases/2.11.12.zip"
	coinutils_cmd_zip := exec.Command("curl", "-L", coinutils_zip_url, "-o", "package.zip")
	err = coinutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coinutils_bin_url := "https://github.com/coin-or/CoinUtils/archive/refs/tags/releases/2.11.12.bin"
	coinutils_cmd_bin := exec.Command("curl", "-L", coinutils_bin_url, "-o", "binary.bin")
	err = coinutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coinutils_src_url := "https://github.com/coin-or/CoinUtils/archive/refs/tags/releases/2.11.12.src.tar.gz"
	coinutils_cmd_src := exec.Command("curl", "-L", coinutils_src_url, "-o", "source.tar.gz")
	err = coinutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coinutils_cmd_direct := exec.Command("./binary")
	err = coinutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
