package main

// AcesContainer - Reference implementation of SMPTE ST2065-4
// Homepage: https://github.com/ampas/aces_container

import (
	"fmt"
	
	"os/exec"
)

func installAcesContainer() {
	// Método 1: Descargar y extraer .tar.gz
	acescontainer_tar_url := "https://github.com/ampas/aces_container/archive/refs/tags/v1.0.2.tar.gz"
	acescontainer_cmd_tar := exec.Command("curl", "-L", acescontainer_tar_url, "-o", "package.tar.gz")
	err := acescontainer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	acescontainer_zip_url := "https://github.com/ampas/aces_container/archive/refs/tags/v1.0.2.zip"
	acescontainer_cmd_zip := exec.Command("curl", "-L", acescontainer_zip_url, "-o", "package.zip")
	err = acescontainer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	acescontainer_bin_url := "https://github.com/ampas/aces_container/archive/refs/tags/v1.0.2.bin"
	acescontainer_cmd_bin := exec.Command("curl", "-L", acescontainer_bin_url, "-o", "binary.bin")
	err = acescontainer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	acescontainer_src_url := "https://github.com/ampas/aces_container/archive/refs/tags/v1.0.2.src.tar.gz"
	acescontainer_cmd_src := exec.Command("curl", "-L", acescontainer_src_url, "-o", "source.tar.gz")
	err = acescontainer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	acescontainer_cmd_direct := exec.Command("./binary")
	err = acescontainer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
