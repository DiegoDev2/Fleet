package main

// Evtx - Windows XML Event Log parser
// Homepage: https://github.com/omerbenamram/evtx

import (
	"fmt"
	
	"os/exec"
)

func installEvtx() {
	// Método 1: Descargar y extraer .tar.gz
	evtx_tar_url := "https://github.com/omerbenamram/evtx/archive/refs/tags/v0.8.3.tar.gz"
	evtx_cmd_tar := exec.Command("curl", "-L", evtx_tar_url, "-o", "package.tar.gz")
	err := evtx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	evtx_zip_url := "https://github.com/omerbenamram/evtx/archive/refs/tags/v0.8.3.zip"
	evtx_cmd_zip := exec.Command("curl", "-L", evtx_zip_url, "-o", "package.zip")
	err = evtx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	evtx_bin_url := "https://github.com/omerbenamram/evtx/archive/refs/tags/v0.8.3.bin"
	evtx_cmd_bin := exec.Command("curl", "-L", evtx_bin_url, "-o", "binary.bin")
	err = evtx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	evtx_src_url := "https://github.com/omerbenamram/evtx/archive/refs/tags/v0.8.3.src.tar.gz"
	evtx_cmd_src := exec.Command("curl", "-L", evtx_src_url, "-o", "source.tar.gz")
	err = evtx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	evtx_cmd_direct := exec.Command("./binary")
	err = evtx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
