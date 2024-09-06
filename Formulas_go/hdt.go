package main

// Hdt - Header Dictionary Triples (HDT) is a compression format for RDF data
// Homepage: https://github.com/rdfhdt/hdt-cpp

import (
	"fmt"
	
	"os/exec"
)

func installHdt() {
	// Método 1: Descargar y extraer .tar.gz
	hdt_tar_url := "https://github.com/rdfhdt/hdt-cpp/archive/refs/tags/v1.3.3.tar.gz"
	hdt_cmd_tar := exec.Command("curl", "-L", hdt_tar_url, "-o", "package.tar.gz")
	err := hdt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hdt_zip_url := "https://github.com/rdfhdt/hdt-cpp/archive/refs/tags/v1.3.3.zip"
	hdt_cmd_zip := exec.Command("curl", "-L", hdt_zip_url, "-o", "package.zip")
	err = hdt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hdt_bin_url := "https://github.com/rdfhdt/hdt-cpp/archive/refs/tags/v1.3.3.bin"
	hdt_cmd_bin := exec.Command("curl", "-L", hdt_bin_url, "-o", "binary.bin")
	err = hdt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hdt_src_url := "https://github.com/rdfhdt/hdt-cpp/archive/refs/tags/v1.3.3.src.tar.gz"
	hdt_cmd_src := exec.Command("curl", "-L", hdt_src_url, "-o", "source.tar.gz")
	err = hdt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hdt_cmd_direct := exec.Command("./binary")
	err = hdt_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: serd")
exec.Command("latte", "install", "serd")
}
