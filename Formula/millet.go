package main

// Millet - Language server for Standard ML (SML)
// Homepage: https://github.com/azdavis/millet

import (
	"fmt"
	
	"os/exec"
)

func installMillet() {
	// Método 1: Descargar y extraer .tar.gz
	millet_tar_url := "https://github.com/azdavis/millet/archive/refs/tags/v0.14.6.tar.gz"
	millet_cmd_tar := exec.Command("curl", "-L", millet_tar_url, "-o", "package.tar.gz")
	err := millet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	millet_zip_url := "https://github.com/azdavis/millet/archive/refs/tags/v0.14.6.zip"
	millet_cmd_zip := exec.Command("curl", "-L", millet_zip_url, "-o", "package.zip")
	err = millet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	millet_bin_url := "https://github.com/azdavis/millet/archive/refs/tags/v0.14.6.bin"
	millet_cmd_bin := exec.Command("curl", "-L", millet_bin_url, "-o", "binary.bin")
	err = millet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	millet_src_url := "https://github.com/azdavis/millet/archive/refs/tags/v0.14.6.src.tar.gz"
	millet_cmd_src := exec.Command("curl", "-L", millet_src_url, "-o", "source.tar.gz")
	err = millet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	millet_cmd_direct := exec.Command("./binary")
	err = millet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
