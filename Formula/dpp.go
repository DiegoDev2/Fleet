package main

// Dpp - Directly include C headers in D source code
// Homepage: https://github.com/atilaneves/dpp

import (
	"fmt"
	
	"os/exec"
)

func installDpp() {
	// Método 1: Descargar y extraer .tar.gz
	dpp_tar_url := "https://github.com/atilaneves/dpp.git"
	dpp_cmd_tar := exec.Command("curl", "-L", dpp_tar_url, "-o", "package.tar.gz")
	err := dpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dpp_zip_url := "https://github.com/atilaneves/dpp.git"
	dpp_cmd_zip := exec.Command("curl", "-L", dpp_zip_url, "-o", "package.zip")
	err = dpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dpp_bin_url := "https://github.com/atilaneves/dpp.git"
	dpp_cmd_bin := exec.Command("curl", "-L", dpp_bin_url, "-o", "binary.bin")
	err = dpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dpp_src_url := "https://github.com/atilaneves/dpp.git"
	dpp_cmd_src := exec.Command("curl", "-L", dpp_src_url, "-o", "source.tar.gz")
	err = dpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dpp_cmd_direct := exec.Command("./binary")
	err = dpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dtools")
	exec.Command("latte", "install", "dtools").Run()
	fmt.Println("Instalando dependencia: dub")
	exec.Command("latte", "install", "dub").Run()
	fmt.Println("Instalando dependencia: ldc")
	exec.Command("latte", "install", "ldc").Run()
}
