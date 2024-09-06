package main

// VulsioGost - Local CVE tracker & notification system
// Homepage: https://github.com/vulsio/gost

import (
	"fmt"
	
	"os/exec"
)

func installVulsioGost() {
	// Método 1: Descargar y extraer .tar.gz
	vulsiogost_tar_url := "https://github.com/vulsio/gost/archive/refs/tags/v0.4.5.tar.gz"
	vulsiogost_cmd_tar := exec.Command("curl", "-L", vulsiogost_tar_url, "-o", "package.tar.gz")
	err := vulsiogost_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulsiogost_zip_url := "https://github.com/vulsio/gost/archive/refs/tags/v0.4.5.zip"
	vulsiogost_cmd_zip := exec.Command("curl", "-L", vulsiogost_zip_url, "-o", "package.zip")
	err = vulsiogost_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulsiogost_bin_url := "https://github.com/vulsio/gost/archive/refs/tags/v0.4.5.bin"
	vulsiogost_cmd_bin := exec.Command("curl", "-L", vulsiogost_bin_url, "-o", "binary.bin")
	err = vulsiogost_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulsiogost_src_url := "https://github.com/vulsio/gost/archive/refs/tags/v0.4.5.src.tar.gz"
	vulsiogost_cmd_src := exec.Command("curl", "-L", vulsiogost_src_url, "-o", "source.tar.gz")
	err = vulsiogost_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulsiogost_cmd_direct := exec.Command("./binary")
	err = vulsiogost_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
