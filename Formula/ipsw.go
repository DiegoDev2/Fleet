package main

// Ipsw - Research tool for iOS & macOS devices
// Homepage: https://blacktop.github.io/ipsw

import (
	"fmt"
	
	"os/exec"
)

func installIpsw() {
	// Método 1: Descargar y extraer .tar.gz
	ipsw_tar_url := "https://github.com/blacktop/ipsw/archive/refs/tags/v3.1.540.tar.gz"
	ipsw_cmd_tar := exec.Command("curl", "-L", ipsw_tar_url, "-o", "package.tar.gz")
	err := ipsw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipsw_zip_url := "https://github.com/blacktop/ipsw/archive/refs/tags/v3.1.540.zip"
	ipsw_cmd_zip := exec.Command("curl", "-L", ipsw_zip_url, "-o", "package.zip")
	err = ipsw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipsw_bin_url := "https://github.com/blacktop/ipsw/archive/refs/tags/v3.1.540.bin"
	ipsw_cmd_bin := exec.Command("curl", "-L", ipsw_bin_url, "-o", "binary.bin")
	err = ipsw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipsw_src_url := "https://github.com/blacktop/ipsw/archive/refs/tags/v3.1.540.src.tar.gz"
	ipsw_cmd_src := exec.Command("curl", "-L", ipsw_src_url, "-o", "source.tar.gz")
	err = ipsw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipsw_cmd_direct := exec.Command("./binary")
	err = ipsw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
