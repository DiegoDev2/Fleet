package main

// Sngrep - Command-line tool for displaying SIP calls message flows
// Homepage: https://github.com/irontec/sngrep

import (
	"fmt"
	
	"os/exec"
)

func installSngrep() {
	// Método 1: Descargar y extraer .tar.gz
	sngrep_tar_url := "https://github.com/irontec/sngrep/archive/refs/tags/v1.8.2.tar.gz"
	sngrep_cmd_tar := exec.Command("curl", "-L", sngrep_tar_url, "-o", "package.tar.gz")
	err := sngrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sngrep_zip_url := "https://github.com/irontec/sngrep/archive/refs/tags/v1.8.2.zip"
	sngrep_cmd_zip := exec.Command("curl", "-L", sngrep_zip_url, "-o", "package.zip")
	err = sngrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sngrep_bin_url := "https://github.com/irontec/sngrep/archive/refs/tags/v1.8.2.bin"
	sngrep_cmd_bin := exec.Command("curl", "-L", sngrep_bin_url, "-o", "binary.bin")
	err = sngrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sngrep_src_url := "https://github.com/irontec/sngrep/archive/refs/tags/v1.8.2.src.tar.gz"
	sngrep_cmd_src := exec.Command("curl", "-L", sngrep_src_url, "-o", "source.tar.gz")
	err = sngrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sngrep_cmd_direct := exec.Command("./binary")
	err = sngrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
