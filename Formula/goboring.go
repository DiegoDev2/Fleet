package main

// GoBoring - Go programming language with BoringCrypto
// Homepage: https://go.googlesource.com/go/+/dev.boringcrypto/README.boringcrypto.md

import (
	"fmt"
	
	"os/exec"
)

func installGoBoring() {
	// Método 1: Descargar y extraer .tar.gz
	goboring_tar_url := "https://go-boringcrypto.storage.googleapis.com/go1.18.10b7.src.tar.gz"
	goboring_cmd_tar := exec.Command("curl", "-L", goboring_tar_url, "-o", "package.tar.gz")
	err := goboring_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goboring_zip_url := "https://go-boringcrypto.storage.googleapis.com/go1.18.10b7.src.zip"
	goboring_cmd_zip := exec.Command("curl", "-L", goboring_zip_url, "-o", "package.zip")
	err = goboring_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goboring_bin_url := "https://go-boringcrypto.storage.googleapis.com/go1.18.10b7.src.bin"
	goboring_cmd_bin := exec.Command("curl", "-L", goboring_bin_url, "-o", "binary.bin")
	err = goboring_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goboring_src_url := "https://go-boringcrypto.storage.googleapis.com/go1.18.10b7.src.src.tar.gz"
	goboring_cmd_src := exec.Command("curl", "-L", goboring_src_url, "-o", "source.tar.gz")
	err = goboring_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goboring_cmd_direct := exec.Command("./binary")
	err = goboring_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
