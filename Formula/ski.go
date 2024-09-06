package main

// Ski - Evade the deadly Yeti on your jet-powered skis
// Homepage: http://catb.org/~esr/ski/

import (
	"fmt"
	
	"os/exec"
)

func installSki() {
	// Método 1: Descargar y extraer .tar.gz
	ski_tar_url := "http://www.catb.org/~esr/ski/ski-6.15.tar.gz"
	ski_cmd_tar := exec.Command("curl", "-L", ski_tar_url, "-o", "package.tar.gz")
	err := ski_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ski_zip_url := "http://www.catb.org/~esr/ski/ski-6.15.zip"
	ski_cmd_zip := exec.Command("curl", "-L", ski_zip_url, "-o", "package.zip")
	err = ski_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ski_bin_url := "http://www.catb.org/~esr/ski/ski-6.15.bin"
	ski_cmd_bin := exec.Command("curl", "-L", ski_bin_url, "-o", "binary.bin")
	err = ski_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ski_src_url := "http://www.catb.org/~esr/ski/ski-6.15.src.tar.gz"
	ski_cmd_src := exec.Command("curl", "-L", ski_src_url, "-o", "source.tar.gz")
	err = ski_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ski_cmd_direct := exec.Command("./binary")
	err = ski_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
}
