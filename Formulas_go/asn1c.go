package main

// Asn1c - Compile ASN.1 specifications into C source code
// Homepage: http://www.lionet.info/asn1c/

import (
	"fmt"
	
	"os/exec"
)

func installAsn1c() {
	// Método 1: Descargar y extraer .tar.gz
	asn1c_tar_url := "https://github.com/vlm/asn1c/releases/download/v0.9.28/asn1c-0.9.28.tar.gz"
	asn1c_cmd_tar := exec.Command("curl", "-L", asn1c_tar_url, "-o", "package.tar.gz")
	err := asn1c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asn1c_zip_url := "https://github.com/vlm/asn1c/releases/download/v0.9.28/asn1c-0.9.28.zip"
	asn1c_cmd_zip := exec.Command("curl", "-L", asn1c_zip_url, "-o", "package.zip")
	err = asn1c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asn1c_bin_url := "https://github.com/vlm/asn1c/releases/download/v0.9.28/asn1c-0.9.28.bin"
	asn1c_cmd_bin := exec.Command("curl", "-L", asn1c_bin_url, "-o", "binary.bin")
	err = asn1c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asn1c_src_url := "https://github.com/vlm/asn1c/releases/download/v0.9.28/asn1c-0.9.28.src.tar.gz"
	asn1c_cmd_src := exec.Command("curl", "-L", asn1c_src_url, "-o", "source.tar.gz")
	err = asn1c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asn1c_cmd_direct := exec.Command("./binary")
	err = asn1c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
