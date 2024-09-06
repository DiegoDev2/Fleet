package main

// Observerward - Web application and service fingerprint identification tool
// Homepage: https://emo-crab.github.io/observer_ward/

import (
	"fmt"
	
	"os/exec"
)

func installObserverward() {
	// Método 1: Descargar y extraer .tar.gz
	observerward_tar_url := "https://github.com/emo-crab/observer_ward/archive/refs/tags/v2024.8.16.tar.gz"
	observerward_cmd_tar := exec.Command("curl", "-L", observerward_tar_url, "-o", "package.tar.gz")
	err := observerward_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	observerward_zip_url := "https://github.com/emo-crab/observer_ward/archive/refs/tags/v2024.8.16.zip"
	observerward_cmd_zip := exec.Command("curl", "-L", observerward_zip_url, "-o", "package.zip")
	err = observerward_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	observerward_bin_url := "https://github.com/emo-crab/observer_ward/archive/refs/tags/v2024.8.16.bin"
	observerward_cmd_bin := exec.Command("curl", "-L", observerward_bin_url, "-o", "binary.bin")
	err = observerward_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	observerward_src_url := "https://github.com/emo-crab/observer_ward/archive/refs/tags/v2024.8.16.src.tar.gz"
	observerward_cmd_src := exec.Command("curl", "-L", observerward_src_url, "-o", "source.tar.gz")
	err = observerward_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	observerward_cmd_direct := exec.Command("./binary")
	err = observerward_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
