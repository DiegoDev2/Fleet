package main

// Dynomite - Generic dynamo implementation for different k-v storage engines
// Homepage: https://github.com/Netflix/dynomite

import (
	"fmt"
	
	"os/exec"
)

func installDynomite() {
	// Método 1: Descargar y extraer .tar.gz
	dynomite_tar_url := "https://github.com/Netflix/dynomite/archive/refs/tags/v0.6.22.tar.gz"
	dynomite_cmd_tar := exec.Command("curl", "-L", dynomite_tar_url, "-o", "package.tar.gz")
	err := dynomite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dynomite_zip_url := "https://github.com/Netflix/dynomite/archive/refs/tags/v0.6.22.zip"
	dynomite_cmd_zip := exec.Command("curl", "-L", dynomite_zip_url, "-o", "package.zip")
	err = dynomite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dynomite_bin_url := "https://github.com/Netflix/dynomite/archive/refs/tags/v0.6.22.bin"
	dynomite_cmd_bin := exec.Command("curl", "-L", dynomite_bin_url, "-o", "binary.bin")
	err = dynomite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dynomite_src_url := "https://github.com/Netflix/dynomite/archive/refs/tags/v0.6.22.src.tar.gz"
	dynomite_cmd_src := exec.Command("curl", "-L", dynomite_src_url, "-o", "source.tar.gz")
	err = dynomite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dynomite_cmd_direct := exec.Command("./binary")
	err = dynomite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
