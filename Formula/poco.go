package main

// Poco - C++ class libraries for building network and internet-based applications
// Homepage: https://pocoproject.org/

import (
	"fmt"
	
	"os/exec"
)

func installPoco() {
	// Método 1: Descargar y extraer .tar.gz
	poco_tar_url := "https://pocoproject.org/releases/poco-1.13.3/poco-1.13.3-all.tar.gz"
	poco_cmd_tar := exec.Command("curl", "-L", poco_tar_url, "-o", "package.tar.gz")
	err := poco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	poco_zip_url := "https://pocoproject.org/releases/poco-1.13.3/poco-1.13.3-all.zip"
	poco_cmd_zip := exec.Command("curl", "-L", poco_zip_url, "-o", "package.zip")
	err = poco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	poco_bin_url := "https://pocoproject.org/releases/poco-1.13.3/poco-1.13.3-all.bin"
	poco_cmd_bin := exec.Command("curl", "-L", poco_bin_url, "-o", "binary.bin")
	err = poco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	poco_src_url := "https://pocoproject.org/releases/poco-1.13.3/poco-1.13.3-all.src.tar.gz"
	poco_cmd_src := exec.Command("curl", "-L", poco_src_url, "-o", "source.tar.gz")
	err = poco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	poco_cmd_direct := exec.Command("./binary")
	err = poco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
