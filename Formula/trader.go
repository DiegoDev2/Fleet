package main

// Trader - Star Traders
// Homepage: https://www.zap.org.au/projects/trader/

import (
	"fmt"
	
	"os/exec"
)

func installTrader() {
	// Método 1: Descargar y extraer .tar.gz
	trader_tar_url := "https://ftp.zap.org.au/pub/trader/unix/trader-7.20.tar.xz"
	trader_cmd_tar := exec.Command("curl", "-L", trader_tar_url, "-o", "package.tar.gz")
	err := trader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trader_zip_url := "https://ftp.zap.org.au/pub/trader/unix/trader-7.20.tar.xz"
	trader_cmd_zip := exec.Command("curl", "-L", trader_zip_url, "-o", "package.zip")
	err = trader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trader_bin_url := "https://ftp.zap.org.au/pub/trader/unix/trader-7.20.tar.xz"
	trader_cmd_bin := exec.Command("curl", "-L", trader_bin_url, "-o", "binary.bin")
	err = trader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trader_src_url := "https://ftp.zap.org.au/pub/trader/unix/trader-7.20.tar.xz"
	trader_cmd_src := exec.Command("curl", "-L", trader_src_url, "-o", "source.tar.gz")
	err = trader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trader_cmd_direct := exec.Command("./binary")
	err = trader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
}
