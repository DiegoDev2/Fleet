package main

// BeancountLanguageServer - Language server for beancount files
// Homepage: https://github.com/polarmutex/beancount-language-server

import (
	"fmt"
	
	"os/exec"
)

func installBeancountLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	beancountlanguageserver_tar_url := "https://github.com/polarmutex/beancount-language-server/archive/refs/tags/v1.3.5.tar.gz"
	beancountlanguageserver_cmd_tar := exec.Command("curl", "-L", beancountlanguageserver_tar_url, "-o", "package.tar.gz")
	err := beancountlanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	beancountlanguageserver_zip_url := "https://github.com/polarmutex/beancount-language-server/archive/refs/tags/v1.3.5.zip"
	beancountlanguageserver_cmd_zip := exec.Command("curl", "-L", beancountlanguageserver_zip_url, "-o", "package.zip")
	err = beancountlanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	beancountlanguageserver_bin_url := "https://github.com/polarmutex/beancount-language-server/archive/refs/tags/v1.3.5.bin"
	beancountlanguageserver_cmd_bin := exec.Command("curl", "-L", beancountlanguageserver_bin_url, "-o", "binary.bin")
	err = beancountlanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	beancountlanguageserver_src_url := "https://github.com/polarmutex/beancount-language-server/archive/refs/tags/v1.3.5.src.tar.gz"
	beancountlanguageserver_cmd_src := exec.Command("curl", "-L", beancountlanguageserver_src_url, "-o", "source.tar.gz")
	err = beancountlanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	beancountlanguageserver_cmd_direct := exec.Command("./binary")
	err = beancountlanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
