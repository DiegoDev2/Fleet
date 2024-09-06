package main

// Ievms - Automated installation of Microsoft IE AppCompat virtual machines
// Homepage: https://xdissent.github.io/ievms/

import (
	"fmt"
	
	"os/exec"
)

func installIevms() {
	// Método 1: Descargar y extraer .tar.gz
	ievms_tar_url := "https://github.com/xdissent/ievms/archive/refs/tags/v0.3.3.tar.gz"
	ievms_cmd_tar := exec.Command("curl", "-L", ievms_tar_url, "-o", "package.tar.gz")
	err := ievms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ievms_zip_url := "https://github.com/xdissent/ievms/archive/refs/tags/v0.3.3.zip"
	ievms_cmd_zip := exec.Command("curl", "-L", ievms_zip_url, "-o", "package.zip")
	err = ievms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ievms_bin_url := "https://github.com/xdissent/ievms/archive/refs/tags/v0.3.3.bin"
	ievms_cmd_bin := exec.Command("curl", "-L", ievms_bin_url, "-o", "binary.bin")
	err = ievms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ievms_src_url := "https://github.com/xdissent/ievms/archive/refs/tags/v0.3.3.src.tar.gz"
	ievms_cmd_src := exec.Command("curl", "-L", ievms_src_url, "-o", "source.tar.gz")
	err = ievms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ievms_cmd_direct := exec.Command("./binary")
	err = ievms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: unar")
exec.Command("latte", "install", "unar")
}
