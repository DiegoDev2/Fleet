package main

// Qxmpp - Cross-platform C++ XMPP client and server library
// Homepage: https://github.com/qxmpp-project/qxmpp/

import (
	"fmt"
	
	"os/exec"
)

func installQxmpp() {
	// Método 1: Descargar y extraer .tar.gz
	qxmpp_tar_url := "https://github.com/qxmpp-project/qxmpp/archive/refs/tags/v1.8.1.tar.gz"
	qxmpp_cmd_tar := exec.Command("curl", "-L", qxmpp_tar_url, "-o", "package.tar.gz")
	err := qxmpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qxmpp_zip_url := "https://github.com/qxmpp-project/qxmpp/archive/refs/tags/v1.8.1.zip"
	qxmpp_cmd_zip := exec.Command("curl", "-L", qxmpp_zip_url, "-o", "package.zip")
	err = qxmpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qxmpp_bin_url := "https://github.com/qxmpp-project/qxmpp/archive/refs/tags/v1.8.1.bin"
	qxmpp_cmd_bin := exec.Command("curl", "-L", qxmpp_bin_url, "-o", "binary.bin")
	err = qxmpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qxmpp_src_url := "https://github.com/qxmpp-project/qxmpp/archive/refs/tags/v1.8.1.src.tar.gz"
	qxmpp_cmd_src := exec.Command("curl", "-L", qxmpp_src_url, "-o", "source.tar.gz")
	err = qxmpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qxmpp_cmd_direct := exec.Command("./binary")
	err = qxmpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
