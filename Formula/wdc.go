package main

// Wdc - WebDAV Client provides easy and convenient to work with WebDAV-servers
// Homepage: https://cloudpolis.github.io/webdav-client-cpp

import (
	"fmt"
	
	"os/exec"
)

func installWdc() {
	// Método 1: Descargar y extraer .tar.gz
	wdc_tar_url := "https://github.com/CloudPolis/webdav-client-cpp/archive/refs/tags/v1.1.5.tar.gz"
	wdc_cmd_tar := exec.Command("curl", "-L", wdc_tar_url, "-o", "package.tar.gz")
	err := wdc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wdc_zip_url := "https://github.com/CloudPolis/webdav-client-cpp/archive/refs/tags/v1.1.5.zip"
	wdc_cmd_zip := exec.Command("curl", "-L", wdc_zip_url, "-o", "package.zip")
	err = wdc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wdc_bin_url := "https://github.com/CloudPolis/webdav-client-cpp/archive/refs/tags/v1.1.5.bin"
	wdc_cmd_bin := exec.Command("curl", "-L", wdc_bin_url, "-o", "binary.bin")
	err = wdc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wdc_src_url := "https://github.com/CloudPolis/webdav-client-cpp/archive/refs/tags/v1.1.5.src.tar.gz"
	wdc_cmd_src := exec.Command("curl", "-L", wdc_src_url, "-o", "source.tar.gz")
	err = wdc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wdc_cmd_direct := exec.Command("./binary")
	err = wdc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: pugixml")
	exec.Command("latte", "install", "pugixml").Run()
}
