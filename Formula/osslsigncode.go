package main

// Osslsigncode - OpenSSL based Authenticode signing for PE/MSI/Java CAB files
// Homepage: https://github.com/mtrojnar/osslsigncode

import (
	"fmt"
	
	"os/exec"
)

func installOsslsigncode() {
	// Método 1: Descargar y extraer .tar.gz
	osslsigncode_tar_url := "https://github.com/mtrojnar/osslsigncode/archive/refs/tags/2.9.tar.gz"
	osslsigncode_cmd_tar := exec.Command("curl", "-L", osslsigncode_tar_url, "-o", "package.tar.gz")
	err := osslsigncode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osslsigncode_zip_url := "https://github.com/mtrojnar/osslsigncode/archive/refs/tags/2.9.zip"
	osslsigncode_cmd_zip := exec.Command("curl", "-L", osslsigncode_zip_url, "-o", "package.zip")
	err = osslsigncode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osslsigncode_bin_url := "https://github.com/mtrojnar/osslsigncode/archive/refs/tags/2.9.bin"
	osslsigncode_cmd_bin := exec.Command("curl", "-L", osslsigncode_bin_url, "-o", "binary.bin")
	err = osslsigncode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osslsigncode_src_url := "https://github.com/mtrojnar/osslsigncode/archive/refs/tags/2.9.src.tar.gz"
	osslsigncode_cmd_src := exec.Command("curl", "-L", osslsigncode_src_url, "-o", "source.tar.gz")
	err = osslsigncode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osslsigncode_cmd_direct := exec.Command("./binary")
	err = osslsigncode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
