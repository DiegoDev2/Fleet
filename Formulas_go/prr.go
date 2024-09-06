package main

// Prr - Mailing list style code reviews for github
// Homepage: https://github.com/danobi/prr

import (
	"fmt"
	
	"os/exec"
)

func installPrr() {
	// Método 1: Descargar y extraer .tar.gz
	prr_tar_url := "https://github.com/danobi/prr/archive/refs/tags/v0.18.0.tar.gz"
	prr_cmd_tar := exec.Command("curl", "-L", prr_tar_url, "-o", "package.tar.gz")
	err := prr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prr_zip_url := "https://github.com/danobi/prr/archive/refs/tags/v0.18.0.zip"
	prr_cmd_zip := exec.Command("curl", "-L", prr_zip_url, "-o", "package.zip")
	err = prr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prr_bin_url := "https://github.com/danobi/prr/archive/refs/tags/v0.18.0.bin"
	prr_cmd_bin := exec.Command("curl", "-L", prr_bin_url, "-o", "binary.bin")
	err = prr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prr_src_url := "https://github.com/danobi/prr/archive/refs/tags/v0.18.0.src.tar.gz"
	prr_cmd_src := exec.Command("curl", "-L", prr_src_url, "-o", "source.tar.gz")
	err = prr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prr_cmd_direct := exec.Command("./binary")
	err = prr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2@1.7")
exec.Command("latte", "install", "libgit2@1.7")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
