package main

// Pluto - CLI tool to help discover deprecated apiVersions in Kubernetes
// Homepage: https://fairwinds.com

import (
	"fmt"
	
	"os/exec"
)

func installPluto() {
	// Método 1: Descargar y extraer .tar.gz
	pluto_tar_url := "https://github.com/FairwindsOps/pluto/archive/refs/tags/v5.20.2.tar.gz"
	pluto_cmd_tar := exec.Command("curl", "-L", pluto_tar_url, "-o", "package.tar.gz")
	err := pluto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pluto_zip_url := "https://github.com/FairwindsOps/pluto/archive/refs/tags/v5.20.2.zip"
	pluto_cmd_zip := exec.Command("curl", "-L", pluto_zip_url, "-o", "package.zip")
	err = pluto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pluto_bin_url := "https://github.com/FairwindsOps/pluto/archive/refs/tags/v5.20.2.bin"
	pluto_cmd_bin := exec.Command("curl", "-L", pluto_bin_url, "-o", "binary.bin")
	err = pluto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pluto_src_url := "https://github.com/FairwindsOps/pluto/archive/refs/tags/v5.20.2.src.tar.gz"
	pluto_cmd_src := exec.Command("curl", "-L", pluto_src_url, "-o", "source.tar.gz")
	err = pluto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pluto_cmd_direct := exec.Command("./binary")
	err = pluto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
