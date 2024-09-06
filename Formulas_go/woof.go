package main

// Woof - Ad-hoc single-file webserver
// Homepage: https://www.home.unix-ag.org/simon/woof.html

import (
	"fmt"
	
	"os/exec"
)

func installWoof() {
	// Método 1: Descargar y extraer .tar.gz
	woof_tar_url := "https://github.com/simon-budig/woof/archive/refs/tags/woof-20220202.tar.gz"
	woof_cmd_tar := exec.Command("curl", "-L", woof_tar_url, "-o", "package.tar.gz")
	err := woof_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	woof_zip_url := "https://github.com/simon-budig/woof/archive/refs/tags/woof-20220202.zip"
	woof_cmd_zip := exec.Command("curl", "-L", woof_zip_url, "-o", "package.zip")
	err = woof_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	woof_bin_url := "https://github.com/simon-budig/woof/archive/refs/tags/woof-20220202.bin"
	woof_cmd_bin := exec.Command("curl", "-L", woof_bin_url, "-o", "binary.bin")
	err = woof_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	woof_src_url := "https://github.com/simon-budig/woof/archive/refs/tags/woof-20220202.src.tar.gz"
	woof_cmd_src := exec.Command("curl", "-L", woof_src_url, "-o", "source.tar.gz")
	err = woof_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	woof_cmd_direct := exec.Command("./binary")
	err = woof_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
