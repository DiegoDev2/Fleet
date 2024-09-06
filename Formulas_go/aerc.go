package main

// Aerc - Email client that runs in your terminal
// Homepage: https://aerc-mail.org/

import (
	"fmt"
	
	"os/exec"
)

func installAerc() {
	// Método 1: Descargar y extraer .tar.gz
	aerc_tar_url := "https://git.sr.ht/~rjarry/aerc/archive/0.18.2.tar.gz"
	aerc_cmd_tar := exec.Command("curl", "-L", aerc_tar_url, "-o", "package.tar.gz")
	err := aerc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aerc_zip_url := "https://git.sr.ht/~rjarry/aerc/archive/0.18.2.zip"
	aerc_cmd_zip := exec.Command("curl", "-L", aerc_zip_url, "-o", "package.zip")
	err = aerc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aerc_bin_url := "https://git.sr.ht/~rjarry/aerc/archive/0.18.2.bin"
	aerc_cmd_bin := exec.Command("curl", "-L", aerc_bin_url, "-o", "binary.bin")
	err = aerc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aerc_src_url := "https://git.sr.ht/~rjarry/aerc/archive/0.18.2.src.tar.gz"
	aerc_cmd_src := exec.Command("curl", "-L", aerc_src_url, "-o", "source.tar.gz")
	err = aerc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aerc_cmd_direct := exec.Command("./binary")
	err = aerc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: scdoc")
exec.Command("latte", "install", "scdoc")
}
