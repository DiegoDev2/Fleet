package main

// Kconf - CLI for managing multiple kubeconfigs
// Homepage: https://github.com/particledecay/kconf

import (
	"fmt"
	
	"os/exec"
)

func installKconf() {
	// Método 1: Descargar y extraer .tar.gz
	kconf_tar_url := "https://github.com/particledecay/kconf/archive/refs/tags/v2.0.0.tar.gz"
	kconf_cmd_tar := exec.Command("curl", "-L", kconf_tar_url, "-o", "package.tar.gz")
	err := kconf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kconf_zip_url := "https://github.com/particledecay/kconf/archive/refs/tags/v2.0.0.zip"
	kconf_cmd_zip := exec.Command("curl", "-L", kconf_zip_url, "-o", "package.zip")
	err = kconf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kconf_bin_url := "https://github.com/particledecay/kconf/archive/refs/tags/v2.0.0.bin"
	kconf_cmd_bin := exec.Command("curl", "-L", kconf_bin_url, "-o", "binary.bin")
	err = kconf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kconf_src_url := "https://github.com/particledecay/kconf/archive/refs/tags/v2.0.0.src.tar.gz"
	kconf_cmd_src := exec.Command("curl", "-L", kconf_src_url, "-o", "source.tar.gz")
	err = kconf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kconf_cmd_direct := exec.Command("./binary")
	err = kconf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
