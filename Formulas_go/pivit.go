package main

// Pivit - Sign and verify data using hardware (Yubikey) backed x509 certificates (PIV)
// Homepage: https://github.com/cashapp/pivit

import (
	"fmt"
	
	"os/exec"
)

func installPivit() {
	// Método 1: Descargar y extraer .tar.gz
	pivit_tar_url := "https://github.com/cashapp/pivit/archive/refs/tags/v0.9.1.tar.gz"
	pivit_cmd_tar := exec.Command("curl", "-L", pivit_tar_url, "-o", "package.tar.gz")
	err := pivit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pivit_zip_url := "https://github.com/cashapp/pivit/archive/refs/tags/v0.9.1.zip"
	pivit_cmd_zip := exec.Command("curl", "-L", pivit_zip_url, "-o", "package.zip")
	err = pivit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pivit_bin_url := "https://github.com/cashapp/pivit/archive/refs/tags/v0.9.1.bin"
	pivit_cmd_bin := exec.Command("curl", "-L", pivit_bin_url, "-o", "binary.bin")
	err = pivit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pivit_src_url := "https://github.com/cashapp/pivit/archive/refs/tags/v0.9.1.src.tar.gz"
	pivit_cmd_src := exec.Command("curl", "-L", pivit_src_url, "-o", "source.tar.gz")
	err = pivit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pivit_cmd_direct := exec.Command("./binary")
	err = pivit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: pcsc-lite")
exec.Command("latte", "install", "pcsc-lite")
}
