package main

// Tt - Command-line utility to manage Tarantool applications
// Homepage: https://github.com/tarantool/tt

import (
	"fmt"
	
	"os/exec"
)

func installTt() {
	// Método 1: Descargar y extraer .tar.gz
	tt_tar_url := "https://github.com/tarantool/tt/releases/download/v2.4.0/tt-2.4.0-complete.tar.gz"
	tt_cmd_tar := exec.Command("curl", "-L", tt_tar_url, "-o", "package.tar.gz")
	err := tt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tt_zip_url := "https://github.com/tarantool/tt/releases/download/v2.4.0/tt-2.4.0-complete.zip"
	tt_cmd_zip := exec.Command("curl", "-L", tt_zip_url, "-o", "package.zip")
	err = tt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tt_bin_url := "https://github.com/tarantool/tt/releases/download/v2.4.0/tt-2.4.0-complete.bin"
	tt_cmd_bin := exec.Command("curl", "-L", tt_bin_url, "-o", "binary.bin")
	err = tt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tt_src_url := "https://github.com/tarantool/tt/releases/download/v2.4.0/tt-2.4.0-complete.src.tar.gz"
	tt_cmd_src := exec.Command("curl", "-L", tt_src_url, "-o", "source.tar.gz")
	err = tt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tt_cmd_direct := exec.Command("./binary")
	err = tt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: mage")
	exec.Command("latte", "install", "mage").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: bash-completion")
	exec.Command("latte", "install", "bash-completion").Run()
}
