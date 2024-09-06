package main

// Sipsak - SIP Swiss army knife
// Homepage: https://github.com/nils-ohlmeier/sipsak/

import (
	"fmt"
	
	"os/exec"
)

func installSipsak() {
	// Método 1: Descargar y extraer .tar.gz
	sipsak_tar_url := "https://github.com/nils-ohlmeier/sipsak/releases/download/0.9.8.1/sipsak-0.9.8.1.tar.gz"
	sipsak_cmd_tar := exec.Command("curl", "-L", sipsak_tar_url, "-o", "package.tar.gz")
	err := sipsak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sipsak_zip_url := "https://github.com/nils-ohlmeier/sipsak/releases/download/0.9.8.1/sipsak-0.9.8.1.zip"
	sipsak_cmd_zip := exec.Command("curl", "-L", sipsak_zip_url, "-o", "package.zip")
	err = sipsak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sipsak_bin_url := "https://github.com/nils-ohlmeier/sipsak/releases/download/0.9.8.1/sipsak-0.9.8.1.bin"
	sipsak_cmd_bin := exec.Command("curl", "-L", sipsak_bin_url, "-o", "binary.bin")
	err = sipsak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sipsak_src_url := "https://github.com/nils-ohlmeier/sipsak/releases/download/0.9.8.1/sipsak-0.9.8.1.src.tar.gz"
	sipsak_cmd_src := exec.Command("curl", "-L", sipsak_src_url, "-o", "source.tar.gz")
	err = sipsak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sipsak_cmd_direct := exec.Command("./binary")
	err = sipsak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
