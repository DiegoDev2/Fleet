package main

// Libwbxml - Library and tools to parse and encode WBXML documents
// Homepage: https://github.com/libwbxml/libwbxml

import (
	"fmt"
	
	"os/exec"
)

func installLibwbxml() {
	// Método 1: Descargar y extraer .tar.gz
	libwbxml_tar_url := "https://github.com/libwbxml/libwbxml/archive/refs/tags/libwbxml-0.11.10.tar.gz"
	libwbxml_cmd_tar := exec.Command("curl", "-L", libwbxml_tar_url, "-o", "package.tar.gz")
	err := libwbxml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwbxml_zip_url := "https://github.com/libwbxml/libwbxml/archive/refs/tags/libwbxml-0.11.10.zip"
	libwbxml_cmd_zip := exec.Command("curl", "-L", libwbxml_zip_url, "-o", "package.zip")
	err = libwbxml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwbxml_bin_url := "https://github.com/libwbxml/libwbxml/archive/refs/tags/libwbxml-0.11.10.bin"
	libwbxml_cmd_bin := exec.Command("curl", "-L", libwbxml_bin_url, "-o", "binary.bin")
	err = libwbxml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwbxml_src_url := "https://github.com/libwbxml/libwbxml/archive/refs/tags/libwbxml-0.11.10.src.tar.gz"
	libwbxml_cmd_src := exec.Command("curl", "-L", libwbxml_src_url, "-o", "source.tar.gz")
	err = libwbxml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwbxml_cmd_direct := exec.Command("./binary")
	err = libwbxml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: wget")
exec.Command("latte", "install", "wget")
}
