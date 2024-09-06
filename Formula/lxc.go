package main

// Lxc - CLI client for interacting with LXD
// Homepage: https://ubuntu.com/lxd

import (
	"fmt"
	
	"os/exec"
)

func installLxc() {
	// Método 1: Descargar y extraer .tar.gz
	lxc_tar_url := "https://github.com/canonical/lxd/releases/download/lxd-6.1/lxd-6.1.tar.gz"
	lxc_cmd_tar := exec.Command("curl", "-L", lxc_tar_url, "-o", "package.tar.gz")
	err := lxc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lxc_zip_url := "https://github.com/canonical/lxd/releases/download/lxd-6.1/lxd-6.1.zip"
	lxc_cmd_zip := exec.Command("curl", "-L", lxc_zip_url, "-o", "package.zip")
	err = lxc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lxc_bin_url := "https://github.com/canonical/lxd/releases/download/lxd-6.1/lxd-6.1.bin"
	lxc_cmd_bin := exec.Command("curl", "-L", lxc_bin_url, "-o", "binary.bin")
	err = lxc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lxc_src_url := "https://github.com/canonical/lxd/releases/download/lxd-6.1/lxd-6.1.src.tar.gz"
	lxc_cmd_src := exec.Command("curl", "-L", lxc_src_url, "-o", "source.tar.gz")
	err = lxc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lxc_cmd_direct := exec.Command("./binary")
	err = lxc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
