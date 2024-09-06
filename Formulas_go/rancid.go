package main

// Rancid - Really Awesome New Cisco confIg Differ
// Homepage: https://www.shrubbery.net/rancid/

import (
	"fmt"
	
	"os/exec"
)

func installRancid() {
	// Método 1: Descargar y extraer .tar.gz
	rancid_tar_url := "https://www.shrubbery.net/pub/rancid/rancid-3.13.tar.gz"
	rancid_cmd_tar := exec.Command("curl", "-L", rancid_tar_url, "-o", "package.tar.gz")
	err := rancid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rancid_zip_url := "https://www.shrubbery.net/pub/rancid/rancid-3.13.zip"
	rancid_cmd_zip := exec.Command("curl", "-L", rancid_zip_url, "-o", "package.zip")
	err = rancid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rancid_bin_url := "https://www.shrubbery.net/pub/rancid/rancid-3.13.bin"
	rancid_cmd_bin := exec.Command("curl", "-L", rancid_bin_url, "-o", "binary.bin")
	err = rancid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rancid_src_url := "https://www.shrubbery.net/pub/rancid/rancid-3.13.src.tar.gz"
	rancid_cmd_src := exec.Command("curl", "-L", rancid_src_url, "-o", "source.tar.gz")
	err = rancid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rancid_cmd_direct := exec.Command("./binary")
	err = rancid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: iputils")
exec.Command("latte", "install", "iputils")
}
