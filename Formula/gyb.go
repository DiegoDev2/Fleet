package main

// Gyb - CLI for backing up and restoring Gmail messages
// Homepage: https://github.com/GAM-team/got-your-back/

import (
	"fmt"
	
	"os/exec"
)

func installGyb() {
	// Método 1: Descargar y extraer .tar.gz
	gyb_tar_url := "https://github.com/GAM-team/got-your-back/archive/refs/tags/v1.82.tar.gz"
	gyb_cmd_tar := exec.Command("curl", "-L", gyb_tar_url, "-o", "package.tar.gz")
	err := gyb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gyb_zip_url := "https://github.com/GAM-team/got-your-back/archive/refs/tags/v1.82.zip"
	gyb_cmd_zip := exec.Command("curl", "-L", gyb_zip_url, "-o", "package.zip")
	err = gyb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gyb_bin_url := "https://github.com/GAM-team/got-your-back/archive/refs/tags/v1.82.bin"
	gyb_cmd_bin := exec.Command("curl", "-L", gyb_bin_url, "-o", "binary.bin")
	err = gyb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gyb_src_url := "https://github.com/GAM-team/got-your-back/archive/refs/tags/v1.82.src.tar.gz"
	gyb_cmd_src := exec.Command("curl", "-L", gyb_src_url, "-o", "source.tar.gz")
	err = gyb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gyb_cmd_direct := exec.Command("./binary")
	err = gyb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
